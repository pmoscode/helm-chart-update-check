package dockerhub

import (
	"encoding/json"
	"github.com/Masterminds/semver/v3"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type Tests struct {
	name           string
	server         *httptest.Server
	expectedResult []string
}

func TestGetDockerVersions(t *testing.T) {
	tests := []Tests{
		{
			name: "one",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write(getTestData())
			})),
			expectedResult: []string{"1.2.3", "1.2.3", "1.2.3-pre.1", "1.2.3-dev", "1.5.0-rc", "1.5.0-rc1", "1.5.0", "1.5.0-nightly.1"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer test.server.Close()

			hub := CreateDockerHubWithUri(test.server.URL, false)
			versions := hub.GetVersions()

			if !reflect.DeepEqual(convertSemverToStringArray(versions), test.expectedResult) {
				t.Errorf("FAILED: expected %v, got %v\n", test.expectedResult, versions)
			}
		})
	}
}

func convertSemverToStringArray(semverVersions []*semver.Version) []string {
	versions := make([]string, len(semverVersions))

	for idx, item := range semverVersions {
		versions[idx] = item.String()
	}

	return versions
}

func getTestData() []byte {
	var serverResponseBody = &ResponseBody{
		Results: []Results{
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
		},
	}

	var serverResponseBodyString, _ = json.Marshal(serverResponseBody)

	return serverResponseBodyString
}
