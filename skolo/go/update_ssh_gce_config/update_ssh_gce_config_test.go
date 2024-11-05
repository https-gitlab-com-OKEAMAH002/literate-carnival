package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.skia.org/infra/go/executil"
	"go.skia.org/infra/go/util"
)

func TestUpdateSshCfg_Success(t *testing.T) {
	ctx := executil.FakeTestsContext(
		"Test_FakeExe_Gcloud_SkiaSwarmingBots",
		"Test_FakeExe_Gcloud_SkiaSwarmingBotsInternal")

	// Write a ssh.cfg file with stale contents.
	sshCfgFile := filepath.Join(t.TempDir(), "ssh.cfg")
	originalSshCfgContents := `Stuff before the autogenerated block.
# BEGIN GCE MACHINES. DO NOT EDIT! This block is automatically generated.
Host should-go-away
	Hostname 1.2.3.4
# END GCE MACHINES.
Stuff after the autogenerated block.
`
	require.NoError(t, util.WithWriteFile(sshCfgFile, func(w io.Writer) error {
		_, err := w.Write([]byte(originalSshCfgContents))
		return err
	}))

	// Update the ssh.cfg file.
	require.NoError(t, updateSshCfg(ctx, sshCfgFile))

	// Read ssh.cfg file and assert that it was correctly updated.
	updatedSshCfgFileBytes, err := os.ReadFile(sshCfgFile)
	require.NoError(t, err)
	expectedSshCfgContents := `Stuff before the autogenerated block.
# BEGIN GCE MACHINES. DO NOT EDIT! This block is automatically generated.
Host skia-d-gce-100
  Hostname 1.1.1.1
Host skia-d-gce-101
  Hostname 2.2.2.2
Host skia-d-gce-102
  Hostname 3.3.3.3
Host skia-d-gce-200
  Hostname 4.4.4.4
Host skia-d-gce-201
  Hostname 5.5.5.5
Host skia-d-gce-202
  Hostname 6.6.6.6
Host skia-d-gce-300
  Hostname 7.7.7.7
Host skia-d-gce-301
  Hostname 8.8.8.8
Host skia-d-gce-302
  Hostname 9.9.9.9
Host skia-e-gce-100
  Hostname 11.11.11.11
Host skia-e-gce-101
  Hostname 12.12.12.12
Host skia-e-gce-102
  Hostname 13.13.13.13
Host skia-e-gce-200
  Hostname 14.14.14.14
Host skia-e-gce-201
  Hostname 15.15.15.15
Host skia-e-gce-202
  Hostname 16.16.16.16
Host skia-e-gce-300
  Hostname 17.17.17.17
Host skia-e-gce-301
  Hostname 18.18.18.18
Host skia-e-gce-302
  Hostname 19.19.19.19
Host skia-i-gce-100
  Hostname 21.21.21.21
Host skia-i-gce-101
  Hostname 22.22.22.22
Host skia-i-gce-102
  Hostname 23.23.23.23
Host skia-i-gce-200
  Hostname 24.24.24.24
Host skia-i-gce-201
  Hostname 25.25.25.25
Host skia-i-gce-202
  Hostname 26.26.26.26
Host skia-i-gce-300
  Hostname 27.27.27.27
Host skia-i-gce-301
  Hostname 28.28.28.28
Host skia-i-gce-302
  Hostname 29.29.29.29
# END GCE MACHINES.
Stuff after the autogenerated block.
`
	assert.Equal(t, expectedSshCfgContents, string(updatedSshCfgFileBytes))
}

func Test_FakeExe_Gcloud_SkiaSwarmingBots(t *testing.T) {
	if !executil.IsCallingFakeCommand() {
		return
	}
	require.Equal(t, []string{
		"gcloud",
		"compute",
		"instances",
		"list",
		"--project=skia-swarming-bots",
		"--format=csv(name, networkInterfaces[0].accessConfigs[0].natIP)",
		"--filter=name~skia-[dei]-*",
		"--sort-by=name",
	}, executil.OriginalArgs())
	fmt.Printf(`name,nat_ip
skia-d-gce-100,1.1.1.1
skia-d-gce-101,2.2.2.2
skia-d-gce-102,3.3.3.3
skia-d-gce-200,4.4.4.4
skia-d-gce-201,5.5.5.5
skia-d-gce-202,6.6.6.6
skia-d-gce-300,7.7.7.7
skia-d-gce-301,8.8.8.8
skia-d-gce-302,9.9.9.9
skia-e-gce-100,11.11.11.11
skia-e-gce-101,12.12.12.12
skia-e-gce-102,13.13.13.13
skia-e-gce-200,14.14.14.14
skia-e-gce-201,15.15.15.15
skia-e-gce-202,16.16.16.16
skia-e-gce-300,17.17.17.17
skia-e-gce-301,18.18.18.18
skia-e-gce-302,19.19.19.19
`)
	os.Exit(0)
}

func Test_FakeExe_Gcloud_SkiaSwarmingBotsInternal(t *testing.T) {
	if !executil.IsCallingFakeCommand() {
		return
	}
	require.Equal(t, []string{
		"gcloud",
		"compute",
		"instances",
		"list",
		"--project=skia-swarming-bots-internal",
		"--format=csv(name, networkInterfaces[0].accessConfigs[0].natIP)",
		"--filter=name~skia-[dei]-*",
		"--sort-by=name",
	}, executil.OriginalArgs())
	fmt.Printf(`name,nat_ip
skia-i-gce-100,21.21.21.21
skia-i-gce-101,22.22.22.22
skia-i-gce-102,23.23.23.23
skia-i-gce-200,24.24.24.24
skia-i-gce-201,25.25.25.25
skia-i-gce-202,26.26.26.26
skia-i-gce-300,27.27.27.27
skia-i-gce-301,28.28.28.28
skia-i-gce-302,29.29.29.29
`)
	os.Exit(0)
}
