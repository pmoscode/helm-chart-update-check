package main

import (
	"encoding/json"
	"github.com/Masterminds/semver/v3"
	"github.com/pmoscode/helm-chart-update-check/pkg/cli"
	"github.com/pmoscode/helm-chart-update-check/pkg/dockerhub"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Tests struct {
	name           string
	server         *httptest.Server
	expectedResult []string
}

func TestCheckVersionNormal(t *testing.T) {
	innerTest(t, "")
}

func TestCheckVersionWithExcludesSimple(t *testing.T) {
	innerTest(t, "1.2.3")
}

func TestCheckVersionWithExcludesRangeMajor(t *testing.T) {
	innerTest(t, "^1.0.0-0")
}

func TestCheckVersionWithExcludesRangeMinor(t *testing.T) {
	innerTest(t, "~3.3.0-0")
}

func innerTest(t *testing.T, excludeVersions string) {
	debug := true
	fail := false
	cliOptions := &cli.Options{
		Debug:                &debug,
		FailOnExistingUpdate: &fail,
		ExcludeVersions:      &excludeVersions,
	}

	test := Tests{
		name: "complete",
		server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(getTestData())
		})),
	}

	defer test.server.Close()

	hub := dockerhub.CreateDockerHubWithUri(test.server.URL, *cliOptions.Debug)
	dockerVersions := hub.GetVersions()

	chartVersion, _ := semver.NewVersion("v1.5.0")

	_, err := checkVersion(chartVersion, dockerVersions, cliOptions)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func getTestData() []byte {
	var serverResponseBody = &dockerhub.ResponseBody{
		Results: []dockerhub.Results{
			{
				Name: "1.2.3",
			},
			{
				Name: "v1.2.3",
			},
			{
				Name: "1.2.3-pre.1",
			},
			{
				Name: "1.2.3-dev",
			},
			{
				Name: "v1.5.0-rc",
			},
			{
				Name: "v1.5.0-rc1",
			},
			{
				Name: "v1.5.0",
			},
			{
				Name: "v1.5.0-nightly.1",
			},
			{
				Name: "v2.2.0",
			},
			{
				Name: "v3.2.1",
			},
			{
				Name: "v3.3.4",
			},
		},
	}

	var serverResponseBodyString, _ = json.Marshal(serverResponseBody)

	return serverResponseBodyString
}
