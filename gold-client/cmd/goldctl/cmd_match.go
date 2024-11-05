package main

import (
	"bytes"
	"context"
	"image"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"go.skia.org/infra/go/skerr"
	"go.skia.org/infra/go/util"
	"go.skia.org/infra/gold-client/go/imgmatching"
	"go.skia.org/infra/gold-client/go/imgmatching/exact"
	"go.skia.org/infra/gold-client/go/imgmatching/fuzzy"
	"go.skia.org/infra/gold-client/go/imgmatching/positive_if_only_image"
	"go.skia.org/infra/gold-client/go/imgmatching/sample_area"
	"go.skia.org/infra/gold-client/go/imgmatching/sobel"
)

// matchEnv provides the environment for the match command.
type matchEnv struct {
	algorithmName string
	parameters    []string
}

// getMatchCmd returns the definition of the match command.
func getMatchCmd() *cobra.Command {
	env := &matchEnv{}
	cmd := &cobra.Command{
		Use:   "match",
		Short: "Runs an image matching algorithm against two images on disk",
		Long: `
Takes a (potentially non-exact) image matching algorithm (and any algorithm-specific parameters)
and executes it against the two given images on disk.

Reports whether or not the specified algorithm considers the two images to be equivalent.

This command is intended for experimenting with various combinations of non-exact image matching
algorithms and parameters.
`,
		Args: cobra.ExactArgs(2), // Takes exactly two images as positional arguments.
		Run:  env.runMatchCmd,
	}

	cmd.Flags().StringVar(&env.algorithmName, "algorithm", "", "Image matching algorithm (e.g. exact, fuzzy, sobel).")
	cmd.Flags().StringArrayVar(&env.parameters, "parameter", []string{}, "Any number of algorithm-specific parameters represented as name:value pairs (e.g. sobel_edge_threshold:10).")
	must(cmd.MarkFlagRequired("algorithm"))

	return cmd
}

func (m *matchEnv) runMatchCmd(cmd *cobra.Command, args []string) {
	ctx := cmd.Context()
	m.Match(ctx, args[0], args[1])
}

// Match instantiates the specified image matching algorithm and runs it against two images.
func (m *matchEnv) Match(ctx context.Context, leftFile, rightFile string) {
	// Load input images.
	leftImage, err := loadPng(leftFile)
	ifErrLogExit(ctx, err)
	rightImage, err := loadPng(rightFile)
	ifErrLogExit(ctx, err)

	// Instantiate the specified algorithm.
	optionalKeys, err := makeOptionalKeys(m.algorithmName, m.parameters)
	ifErrLogExit(ctx, err)
	algorithmName, matcher, err := imgmatching.MakeMatcher(optionalKeys)
	ifErrLogExit(ctx, err)

	// Run the algorithm against the two input images.
	imagesMatch := matcher.Match(leftImage, rightImage)

	// Report the algorithm's boolean output.
	if imagesMatch {
		logInfo(ctx, "Images match.\n")
	} else {
		logInfo(ctx, "Images do not match.\n")
	}

	// Print out algorithm-specific debug information.
	switch algorithmName {
	case imgmatching.ExactMatching:
		printOutExactDebugInfo(ctx, matcher.(*exact.Matcher))
	case imgmatching.FuzzyMatching:
		printOutFuzzyDebugInfo(ctx, matcher.(*fuzzy.Matcher))
	case imgmatching.PositiveIfOnlyImageMatching:
		printOutPositiveIfOnlyImageDebugInfo(ctx, matcher.(*positive_if_only_image.Matcher))
	case imgmatching.SampleAreaMatching:
		printOutSampleAreaDebugInfo(ctx, matcher.(*sample_area.Matcher))
	case imgmatching.SobelFuzzyMatching:
		err := printOutSobelDebugInfo(ctx, matcher.(*sobel.Matcher))
		ifErrLogExit(ctx, err)
	}

	exitProcess(ctx, 0)
}

// loadPng loads a PNG image from disk.
func loadPng(fileName string) (image.Image, error) {
	// Load the image and save the bytes because we need to return them.
	imgBytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, skerr.Wrapf(err, "loading file %s", fileName)
	}
	img, err := png.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return nil, skerr.Wrapf(err, "decoding PNG file %s", fileName)
	}
	return img, nil
}

// makeOptionalKeys turns the specified algorithm name and parameters into the equivalent map of
// optional keys that function imgmatching.MakeMatcher() expects.
func makeOptionalKeys(algorithmName string, parameters []string) (map[string]string, error) {
	keys := map[string]string{
		imgmatching.AlgorithmNameOptKey: algorithmName,
	}

	for _, pair := range parameters {
		split := strings.SplitN(pair, ":", 2)
		if len(split) != 2 {
			return nil, skerr.Fmt("parameter %q must be a key:value pair", pair)
		}
		keys[split[0]] = split[1]
	}

	return keys, nil
}

// printOutExactDebugInfo prints out stats debug info reported by the given exact.Matcher.
func printOutExactDebugInfo(ctx context.Context, matcher *exact.Matcher) {
	printDebugInfoItem(ctx, "Last different pixel found", matcher.LastDifferentPixelFound())
}

// printOutFuzzyDebugInfo prints out stats reported by the given fuzzy.Matcher.
func printOutFuzzyDebugInfo(ctx context.Context, matcher *fuzzy.Matcher) {
	printDebugInfoItem(ctx, "Number of different pixels", matcher.NumDifferentPixels())
	printDebugInfoItem(ctx, "Maximum delta", matcher.MaxPixelDelta())
	printDebugInfoItem(ctx, "Pixel comparison method", matcher.PixelComparisonMethod())
}

// printOutPositiveIfOnlyImageDebugInfo prints out stats reported by the given
// positive_if_only_image.Matcher.
func printOutPositiveIfOnlyImageDebugInfo(ctx context.Context, matcher *positive_if_only_image.Matcher) {
	printDebugInfoItem(ctx, "Last different pixel found", matcher.LastDifferentPixelFound())
}

// printOutSampleAreaDebugInfo prints out stats reported by the given
// sample_area.Matcher.
func printOutSampleAreaDebugInfo(ctx context.Context, matcher *sample_area.Matcher) {
	printDebugInfoItem(ctx, "Failed sample area", matcher.FailedSampleArea())
	printDebugInfoItem(ctx, "Number of different pixels", matcher.NumDifferentPixels())
	printDebugInfoItem(ctx, "Sample area width too small", matcher.SampleAreaWidthTooSmall())
	printDebugInfoItem(ctx, "Sample area width too large", matcher.SampleAreaWidthTooLarge())
	printDebugInfoItem(ctx, "Max different pixels per area out of range", matcher.MaxDifferentPixelsPerAreaOutOfRange())
	printDebugInfoItem(ctx, "Sample area pixel delta threshold out of range", matcher.SampleAreaChannelDeltaThresholdOutOfRange())
}

// printOutSobelDebugInfo writes intermediate images generated by the given sobel.Matcher to a
// temporary directory and prints out the resulting paths. It also prints out the stats reported by
// the embedded fuzzy.Matcher.
func printOutSobelDebugInfo(ctx context.Context, matcher *sobel.Matcher) error {
	tempDir, err := os.MkdirTemp("", "goldctl-*")
	if err != nil {
		return skerr.Wrap(err)
	}

	writeToDiskAndPrintOut := func(ctx context.Context, msg, filename string, img image.Image) error {
		p := filepath.Join(tempDir, filename)
		err := writePngToTmp(img, p)
		if err != nil {
			return skerr.Wrap(err)
		}
		printDebugInfoItem(ctx, msg, p)
		return nil
	}

	err = writeToDiskAndPrintOut(ctx, "Sobel operator output (left image)", "sobel-output.png", matcher.SobelOutput())
	if err != nil {
		return skerr.Wrap(err)
	}

	err = writeToDiskAndPrintOut(ctx, "Left image with edges removed", "image1-no-edges.png", matcher.ExpectedImageWithEdgesRemoved())
	if err != nil {
		return skerr.Wrap(err)
	}

	err = writeToDiskAndPrintOut(ctx, "Right image with edges removed", "image2-no-edges.png", matcher.ActualImageWithEdgesRemoved())
	if err != nil {
		return skerr.Wrap(err)
	}

	printOutFuzzyDebugInfo(ctx, &matcher.Matcher)
	return nil
}

// writePngToTmp takes an image, saves it to disk as a PNG image with the given filename.
func writePngToTmp(img image.Image, filename string) error {
	err := util.WithWriteFile(filename, func(writer io.Writer) error {
		err := png.Encode(writer, img)
		if err != nil {
			return skerr.Wrap(err)
		}
		return nil
	})
	if err != nil {
		return skerr.Wrap(err)
	}
	return nil
}

// printDebugInfoItem prints out the given debug name/value pair. The name is left-padded so that
// multiple calls to this function produce vertically aligned output that is easier to read.
func printDebugInfoItem(ctx context.Context, name string, value interface{}) {
	logInfof(ctx, "%36s: %v\n", name, value)
}
