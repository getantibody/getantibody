package getantibody

import (
	"fmt"
	"github.com/google/go-github/github"
	"strings"
)

const repo = "https://github.com/caarlos0/antibody"
const downloadURL = repo+"/releases/download/%s/antibody_%s_%s.tar.gz"

type OS struct {
	ID string
	Name string
}

type Arch struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

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

func DownloadURL(version, os, arch string) string {
	parsedArch := strings.ToLower(arch)
	if parsedArch == "x86_64" {
		parsedArch = "amd64"
	}
	return fmt.Sprintf(
		downloadURL,
		version,
		strings.ToLower(os),
		parsedArch,
	)
}

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
