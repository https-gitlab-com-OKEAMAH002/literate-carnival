package urlprovider

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"go.skia.org/infra/go/sklog"
	perfgit "go.skia.org/infra/perf/go/git"
	"go.skia.org/infra/perf/go/types"
)

type URLProvider struct {
	perfGit perfgit.Git
}

// Explore generates a url to the explore page for the given parameters
func (prov *URLProvider) Explore(ctx context.Context, startCommitNumber int, endCommitNumber int, parameters map[string][]string, disableFilterParentTraces bool) string {
	queryUrl := url.Values{}
	prov.fillCommonParams(ctx, queryUrl, startCommitNumber, endCommitNumber)
	// Now let's look at the parameters for the query

	queryUrl["queries"] = []string{prov.GetQueryStringFromParameters(parameters)}
	if disableFilterParentTraces {
		queryUrl["disable_filter_parent_traces"] = []string{"true"}
	}

	return "/e/?" + queryUrl.Encode()
}

func (prov *URLProvider) MultiGraph(ctx context.Context, startCommitNumber int, endCommitNumber int, shortcutId string) string {
	queryUrl := url.Values{}
	prov.fillCommonParams(ctx, queryUrl, startCommitNumber, endCommitNumber)
	queryUrl["shortcut"] = []string{shortcutId}

	return "/m/?" + queryUrl.Encode()
}

// New creates a new instance of the UrlProvider struct
func New(perfgit perfgit.Git) *URLProvider {
	return &URLProvider{
		perfGit: perfgit,
	}
}

func (prov *URLProvider) GetQueryStringFromParameters(parameters map[string][]string) string {
	query_portion := url.Values{}
	for paramName, paramValues := range parameters {
		query_portion[paramName] = paramValues
	}

	return query_portion.Encode()
}

func (prov *URLProvider) fillCommonParams(ctx context.Context, queryUrl url.Values, startCommitNumber int, endCommitNumber int) {
	startCommit, err := prov.perfGit.CommitFromCommitNumber(ctx, types.CommitNumber(startCommitNumber))
	if err != nil {
		sklog.Error("Error getting commit info")
	}
	endCommit, _ := prov.perfGit.CommitFromCommitNumber(ctx, types.CommitNumber(endCommitNumber))
	queryUrl["begin"] = []string{strconv.Itoa(int(startCommit.Timestamp))}

	// We will shift the end time by a day so the graph doesn't render the anomalies right at the end
	endTime := time.Unix(endCommit.Timestamp, 0).AddDate(0, 0, 1)
	queryUrl["end"] = []string{strconv.Itoa(int(endTime.Unix()))}

	queryUrl["summary"] = []string{"true"}
}
