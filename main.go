package main

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/pmoscode/go-common/shutdown"
	chart2 "github.com/pmoscode/helm-chart-update-check/pkg/chart"
	"github.com/pmoscode/helm-chart-update-check/pkg/cli"
	"github.com/pmoscode/helm-chart-update-check/pkg/dockerhub"
	"github.com/pmoscode/helm-chart-update-check/pkg/utils"
	"log"
)

func main() {
	defer shutdown.ExitOnPanic()

	cliOptions := cli.GetCliOptions()

	dockerVersions := dockerhub.FetchDockerVersions(cliOptions)
	chartVersion := chart2.FetchChartVersion(cliOptions)

	_, err := checkVersion(chartVersion, dockerVersions, cliOptions)
	if err != nil {
		log.Fatalln(err)
	}
}

func checkVersion(chartVersion *semver.Version, dockerVersions []*semver.Version, cliOptions *cli.Options) (int, error) {
	// See: https://github.com/Masterminds/semver?tab=readme-ov-file#working-with-prerelease-versions
	constraint, err := semver.NewConstraint(fmt.Sprintf("<= %s-0", chartVersion.IncPatch().String()))
	if err != nil {
		log.Fatalln(err)
	}

	newerVersions := make([]*semver.Version, 0)
	excludeVersions := utils.GetExcludedVersionsSimple(*cliOptions.ExcludeVersions)

	fmt.Printf("Checking, if some version is > %s\n", chartVersion.String())
	for _, item := range dockerVersions {
		if *cliOptions.Debug {
			fmt.Printf("Checking if Helm chart version %v is >= DockerHub version %v: ", chartVersion.Original(), item.Original())
		}

		skipVersion := false
		if excludeVersions != nil {
			for _, version := range excludeVersions {
				constraintExclude, err := semver.NewConstraint(fmt.Sprintf("%s", version))
				if err != nil {
					log.Fatalln(err)
				}

				if constraintExclude.Check(item) {
					skipVersion = true
					break
				}
			}
		}

		if !skipVersion {
			if !constraint.Check(item) {
				newerVersions = append(newerVersions, item)
				if *cliOptions.Debug {
					fmt.Println(false)
				}
			} else {
				if *cliOptions.Debug {
					fmt.Println(true)
				}
			}
		} else {
			fmt.Println("skipped")
		}
	}

	newerVersionsCnt := len(newerVersions)

	if newerVersionsCnt > 0 {
		fmt.Println("Newer dockerVersions exists: ")
		for _, item := range newerVersions {
			fmt.Println(item.Original())
		}

		if *cliOptions.FailOnExistingUpdate {
			return newerVersionsCnt, fmt.Errorf("FAIL: Found %d new versions", newerVersionsCnt)
		}
	} else {
		fmt.Println("No newer versions found.")
	}

	return newerVersionsCnt, nil
}
