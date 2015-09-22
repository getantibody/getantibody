package getantibody

import "github.com/google/go-github/github"

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
