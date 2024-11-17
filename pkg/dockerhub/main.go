package dockerhub

import (
	"github.com/Masterminds/semver/v3"
	"github.com/pmoscode/helm-chart-update-check/pkg/cli"
	"strings"
)

const (
	apiUri   = "https://hub.docker.com/v2/repositories/"
	tagsPath = "/tags"
)

func FetchDockerVersions(cliOptions *cli.Options) []*semver.Version {
	dockerHub := CreateDockerHub(*cliOptions.DockerHubRepository, *cliOptions.Debug)
	versions := dockerHub.GetVersions()

	return versions
}

func CreateDockerHub(repository string, debug bool) *DockerHub {
	dockerHub := &DockerHub{
		uri:   apiUri + repository + tagsPath,
		debug: debug,
	}

	return dockerHub
}

func CreateDockerHubWithUri(uri string, debug bool) *DockerHub {
	dockerHub := &DockerHub{
		uri:   uri,
		debug: debug,
	}

	return dockerHub
}

func isVersionApplicable(version string) bool {
	if version == "latest" || strings.Count(version, ".") < 2 {
		return false
	}

	return true
}
