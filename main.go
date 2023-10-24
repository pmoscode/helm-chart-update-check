package main

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	chart2 "github.com/pmoscode/helm-chart-update-check/pkg/chart"
	"gopkg.in/yaml.v3"
	"log"
	"strings"
	"time"
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

func main() {
	//response, err := http.Get("https://hub.docker.com/v2/repositories/pmoscode/axelor-open-suite/tags")
	//
	//if err != nil {
	//	fmt.Print(err.Error())
	//	os.Exit(1)
	//}
	//
	//responseData, err := io.ReadAll(response.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//var responseBody = ResponseBody{}
	//err = json.Unmarshal([]byte(responseData), &responseBody)
	//if err != nil {
	//	return
	//}
	//
	//for _, item := range responseBody.Results {
	//	if item.Name == "latest" {
	//		continue
	//	}
	//	fmt.Println(item.Name)
	//	v, _ := semver.NewVersion(item.Name)
	//	fmt.Println(v.Minor())
	//	fmt.Println(v.Value())
	//}

	chart := chart2.NewChart("/home/peter/Arbeit/GIT/GitHub/Helm-Charts/airsonic-advanced")

	//version := chart.Version()
	//fmt.Println("version:")
	//fmt.Println(version)

	appVersion := strings.Trim(chart.AppVersion(), "\"")

	v, err := semver.NewVersion(appVersion)
	if err != nil {
		log.Fatal("Problem creating semver: ", err)
	}

	testVersion := "<= 11.0.1"

	constraint, _ := semver.NewConstraint(testVersion)

	compare := constraint.Check(v)
	fmt.Printf("Version %s is %s: %t\n", appVersion, testVersion, compare)

	if compare {
		fmt.Println(v.IncMinor())
	} else {
		fmt.Println("No change needed")
	}

	// Create a modified yaml file
	//f, err := os.Create("/home/peter/Arbeit/GIT/GitHub/Helm-Charts/airsonic-advanced/Chart.yaml")
	//if err != nil {
	//	log.Fatalf("Problem creating file: %v", err)
	//}
	//defer f.Close()
	//yamlEncoder := yaml.NewEncoder(f)
	//yamlEncoder.SetIndent(2)
	//yamlEncoder.Encode(dockerCompose.Content[0])
}

// Recusive function to find the child node by value that we care about.
// Probably needs tweaking so use with caution.
func findChildNode(value string, node *yaml.Node) *yaml.Node {
	for _, v := range node.Content {
		// If we found the value we are looking for, return it.
		fmt.Printf("%+v", v)
		fmt.Println()
		if v.Value == value {
			return v
		}
		// Otherwise recursively look more
		if child := findChildNode(value, v); child != nil {
			return child
		}
	}
	return nil
}
