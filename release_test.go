package getantibody_test

import (
	"testing"

	"github.com/caarlos0/getantibody"
	"github.com/stretchr/testify/assert"
)

func TestLatest(t *testing.T) {
	release, err := getantibody.LatestRelease()
	assert.NotEmpty(t, release)
	assert.NoError(t, err)
}

func TestDownloadURLValidOSAndArch(t *testing.T) {
	release, _ := getantibody.LatestRelease()
	url, err := getantibody.DownloadURL(
		release,
		"darwin",
		"x86_64",
	)
	assert.Contains(t, url, "github.com")
	assert.NoError(t, err)
}

func TestDownloadURLInvalidOS(t *testing.T) {
	release, _ := getantibody.LatestRelease()
	url, err := getantibody.DownloadURL(
		release,
		"windows",
		"x86_64",
	)
	assert.Empty(t, url)
	assert.Error(t, err)
}

func TestDownloadURLInvalidArch(t *testing.T) {
	release, _ := getantibody.LatestRelease()
	url, err := getantibody.DownloadURL(
		release,
		"darwin",
		"ppc",
	)
	assert.Empty(t, url)
	assert.Error(t, err)
}

func TestDistributions(t *testing.T) {
	assert.Len(t, getantibody.Distributions(), 5)
}
