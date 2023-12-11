package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Masterminds/semver/v3"
	chart2 "github.com/pmoscode/helm-chart-update-check/pkg/chart"
	"github.com/pmoscode/helm-chart-update-check/pkg/dockerhub"
	"log"
	"strings"
)

type CliOptions struct {
	dockerHubRepository  *string
	helmChartPath        *string
	failOnExistingUpdate *bool
	debug                *bool
}

func getCliOptionsParameters() *CliOptions {
	dockerHubRepository := flag.String("docker-hub-repo", "", "DockHub repo to check tag versions")
	helmChartPath := flag.String("helm-chart-path", ".", "Helm chart to check for updates")
	failOnExistingUpdate := flag.Bool("fail-on-update", false, "Return exit code 1, if update is available")
	debug := flag.Bool("debug", false, "Enable debug outputs")

	flag.Parse()

	return &CliOptions{
		dockerHubRepository:  dockerHubRepository,
		helmChartPath:        helmChartPath,
		failOnExistingUpdate: failOnExistingUpdate,
		debug:                debug,
	}
}

func main() {
	cliOptions := getCliOptions()

	dockerVersions := getDockerVersions(cliOptions)

	chartVersion := getChartVersion(cliOptions)

	_, err := checkVersion(chartVersion, dockerVersions, cliOptions)
	if err != nil {
		log.Fatalln(err)
	}
}

func getCliOptions() *CliOptions {
	cliOptions := getCliOptionsParameters()

	if *cliOptions.dockerHubRepository == "" {
		log.Fatal(errors.New("parameter 'docker-hub-repo' is required"))
	}
	if *cliOptions.helmChartPath == "" {
		log.Fatal(errors.New("parameter 'helm-chart-path' is required"))
	}

	return cliOptions
}

func getDockerVersions(cliOptions *CliOptions) []*semver.Version {
	dockerHub := dockerhub.CreateDockerHub(*cliOptions.dockerHubRepository, *cliOptions.debug)
	versions := dockerHub.GetVersions()

	return versions
}

func getChartVersion(cliOptions *CliOptions) *semver.Version {
	chart := chart2.NewChart(*cliOptions.helmChartPath)

	appVersion := strings.Trim(chart.AppVersion(), "\"")
	fmt.Println("Helm chart AppVersion:")
	fmt.Println(appVersion)

	v, err := semver.NewVersion(appVersion)
	if err != nil {
		log.Fatal("Problem creating semver: ", err)
	}
	return v
}

func checkVersion(chartVersion *semver.Version, dockerVersions []*semver.Version, cliOptions *CliOptions) (int, error) {
	constraintStr := fmt.Sprintf("<= %s-0", chartVersion.String())
	// See: https://github.com/Masterminds/semver?tab=readme-ov-file#working-with-prerelease-versions
	constraint, _ := semver.NewConstraint(constraintStr)

	newerVersions := make([]*semver.Version, 0)

	fmt.Printf("Checking, if some version is > %s\n", chartVersion.String())
	for _, item := range dockerVersions {
		if *cliOptions.debug {
			fmt.Printf("Checking if Helm chart version %v is > DockerHub version %v: ", chartVersion.String(), item.String())
		}
		if !constraint.Check(item) {
			newerVersions = append(newerVersions, item)
			if *cliOptions.debug {
				fmt.Println(false)
			}
		} else {
			if *cliOptions.debug {
				fmt.Println(true)
			}
		}
	}

	newerVersionsCnt := len(newerVersions)

	if newerVersionsCnt > 0 {
		fmt.Println("Newer dockerVersions exists: ")
		for _, item := range newerVersions {
			fmt.Println(item.Original())
		}

		if *cliOptions.failOnExistingUpdate {
			return newerVersionsCnt, fmt.Errorf("FAIL: Found %d new versions", newerVersionsCnt)
		}
	} else {
		fmt.Println("No newer versions found.")
	}

	return newerVersionsCnt, nil
}
