package dockerhub

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Masterminds/semver/v3"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	apiUri   = "https://hub.docker.com/v2/repositories/"
	tagsPath = "/tags"
)

type ResponseBody struct {
	Count    int       `json:"count,omitempty"`
	Next     any       `json:"next,omitempty"`
	Previous any       `json:"previous,omitempty"`
	Results  []Results `json:"results,omitempty"`
}

type Images struct {
	Architecture string    `json:"architecture,omitempty"`
	Features     string    `json:"features,omitempty"`
	Variant      any       `json:"variant,omitempty"`
	Digest       string    `json:"digest,omitempty"`
	Os           string    `json:"os,omitempty"`
	OsFeatures   string    `json:"os_features,omitempty"`
	OsVersion    any       `json:"os_version,omitempty"`
	Size         int       `json:"size,omitempty"`
	Status       string    `json:"status,omitempty"`
	LastPulled   time.Time `json:"last_pulled,omitempty"`
	LastPushed   time.Time `json:"last_pushed,omitempty"`
}

type Results struct {
	Creator             int       `json:"creator,omitempty"`
	ID                  int       `json:"id,omitempty"`
	Images              []Images  `json:"images,omitempty"`
	LastUpdated         time.Time `json:"last_updated,omitempty"`
	LastUpdater         int       `json:"last_updater,omitempty"`
	LastUpdaterUsername string    `json:"last_updater_username,omitempty"`
	Name                string    `json:"name,omitempty"`
	Repository          int       `json:"repository,omitempty"`
	FullSize            int       `json:"full_size,omitempty"`
	V2                  bool      `json:"v2,omitempty"`
	TagStatus           string    `json:"tag_status,omitempty"`
	TagLastPulled       time.Time `json:"tag_last_pulled,omitempty"`
	TagLastPushed       time.Time `json:"tag_last_pushed,omitempty"`
	MediaType           string    `json:"media_type,omitempty"`
	ContentType         string    `json:"content_type,omitempty"`
	Digest              string    `json:"digest,omitempty"`
}

type DockerHub struct {
	uri        string
	repository string
	debug      bool
}

func (hub *DockerHub) GetVersions() []*semver.Version {
	responseBody := hub.receiveData()

	return hub.mapToSemver(responseBody)
}

func (hub *DockerHub) receiveData() *ResponseBody {
	response, err := http.Get(hub.uri)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	var responseBody = &ResponseBody{}
	err = json.Unmarshal(responseData, responseBody)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	return responseBody
}

func (hub *DockerHub) mapToSemver(responseBody *ResponseBody) []*semver.Version {
	semverData := make([]*semver.Version, 0)

	for _, item := range responseBody.Results {
		if !isVersionApplicable(item.Name) {
			continue
		}

		v, err := semver.NewVersion(item.Name)
		if err != nil {
			if !errors.Is(err, semver.ErrInvalidSemVer) {
				fmt.Printf("Version '%s' is not a valid Semver version\n", item.Name)
				fmt.Println(err)
			} else {
				if hub.debug {
					fmt.Printf("Skipping partial version: %s\n", item.Name)
				}
			}
		} else {
			if hub.debug {
				fmt.Printf("Converting %s to %s\n", item.Name, v.String())
			}
			semverData = append(semverData, v)
		}
	}

	return semverData
}

func isVersionApplicable(version string) bool {
	if version == "latest" || strings.Count(version, ".") < 2 {
		return false
	}

	return true
}

func CreateDockerHub(repository string, debug bool) *DockerHub {
	dockerHub := &DockerHub{
		uri:        apiUri + repository + tagsPath,
		repository: repository,
		debug:      debug,
	}

	return dockerHub
}
