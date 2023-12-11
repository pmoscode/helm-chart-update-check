package dockerhub

import (
	"testing"
)

func TestMapToSemver(t *testing.T) {
	var skipResults = []Results{
		{
			Name: "1.2",
		},
		{
			Name: "1",
		},
	}
	var lengthSkippedResults = len(skipResults)

	var okResults = []Results{
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
			Name: "v1.4.3",
		},
	}

	allResults := append(skipResults, okResults...)

	var responseBody = &ResponseBody{
		Results: allResults,
	}

	dockerHub := CreateDockerHub("", false)
	semver := dockerHub.mapToSemver(responseBody)

	lengthResponseBody := len(allResults)
	lengthResult := len(semver)

	lengthEquals := (lengthResponseBody - lengthSkippedResults) == lengthResult

	if !lengthEquals {
		t.Fatalf("The testing array must equals the resulting array: %d != %d", lengthResponseBody, lengthResult)
	}
}
