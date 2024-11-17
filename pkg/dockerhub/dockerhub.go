package dockerhub

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Masterminds/semver/v3"
	"io"
	"net/http"
	"os"
)

type DockerHub struct {
	uri   string
	debug bool
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
