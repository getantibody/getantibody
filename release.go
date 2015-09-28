package getantibody

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/go-github/github"
)

const repo = "https://github.com/caarlos0/antibody"
const downloadURL = repo + "/releases/download/%s/antibody_%s_%s.tar.gz"

// OS type defines an Operating System
type OS struct {
	ID   string
	Name string
}

// Arch type defines an architecture
type Arch struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Distribution defines a distribution, combining OS and Arch
type Distribution struct {
	OS     string `json:"os"`
	Arches []Arch `json:"arches"`
	Name   string `json:"name"`
}

var arches = []Arch{
	{
		ID:   "amd64",
		Name: "amd64",
	}, {
		ID:   "386",
		Name: "i386",
	},
}

var oses = []OS{
	{
		ID:   "darwin",
		Name: "Mac OS",
	}, {
		ID:   "linux",
		Name: "Linux",
	}, {
		ID:   "freebsd",
		Name: "FreeBSD",
	}, {
		ID:   "openbsd",
		Name: "OpenBSD",
	}, {
		ID:   "netsd",
		Name: "NetBSD",
	},
}

// LatestRelease return the last release tag name.
func LatestRelease() (string, error) {
	client := github.NewClient(nil)
	releases, _, err := client.Repositories.ListReleases(
		"caarlos0", "antibody", nil,
	)
	if err != nil {
		return "", err
	}
	return *releases[0].TagName, nil
}

// DownloadURL gets the download url for the given version, os and arch.
// os and arch should be in the format of uname` commands.
func DownloadURL(version, os, arch string) (string, error) {
	parsedArch := strings.ToLower(arch)
	parsedOs := strings.ToLower(os)
	if parsedArch == "x86_64" {
		parsedArch = "amd64"
	}
	if !isValidArch(parsedArch) {
		return "", errors.New("Arch " + parsedArch + " is not supported!")
	}
	if !isValidOS(parsedOs) {
		return "", errors.New("OS " + parsedOs + " is not supported!")
	}
	return fmt.Sprintf(
		downloadURL,
		version,
		parsedOs,
		parsedArch,
	), nil
}

func isValidOS(s string) bool {
	for _, os := range oses {
		if s == os.ID {
			return true
		}
	}
	return false
}

func isValidArch(s string) bool {
	for _, arch := range arches {
		if s == arch.ID {
			return true
		}
	}
	return false
}

// Distributions lists the available antibody flavors
func Distributions() []Distribution {
	var distributions []Distribution
	for _, os := range oses {
		distributions = append(
			distributions,
			Distribution{
				OS:     os.ID,
				Arches: arches,
				Name:   os.Name,
			},
		)
	}
	return distributions
}
