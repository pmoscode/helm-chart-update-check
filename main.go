package main

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	chart2 "github.com/pmoscode/helm-chart-update-check/pkg/chart"
	"github.com/pmoscode/helm-chart-update-check/pkg/dockerhub"
	"log"
	"os"
	"strings"
)

func main() {
	debugEnv := os.Getenv("HCUC_DEBUG_ENABLED")
	debugEnabled := debugEnv == "TRUE" || debugEnv == "true"

	dockerHubRepositoryEnv, err := getRequiredEnv("HCUC_DOCKERHUB_REPO")
	if err != nil {
		log.Fatal(err)
	}

	helmChartPathEnv, err := getRequiredEnv("HCUC_HELM_CHART_PATH")
	if err != nil {
		log.Fatal(err)
	}

	dockerHub := dockerhub.CreateDockerHub(dockerHubRepositoryEnv, debugEnabled)
	versions := dockerHub.GetVersions()

	chart := chart2.NewChart(helmChartPathEnv)

	appVersion := strings.Trim(chart.AppVersion(), "\"")
	fmt.Println("Helm chart AppVersion:")
	fmt.Println(appVersion)

	v, err := semver.NewVersion(appVersion)
	if err != nil {
		log.Fatal("Problem creating semver: ", err)
	}

	constraint, _ := semver.NewConstraint("<=" + v.String())

	newerVersions := make([]*semver.Version, 0)

	for _, item := range versions {
		if !constraint.Check(item) {
			newerVersions = append(newerVersions, item)
		}
	}

	if len(newerVersions) > 0 {
		fmt.Println("Newer versions exists: ")
		for _, item := range newerVersions {
			fmt.Println(item)
		}

		os.Exit(1)
	}
}

func getRequiredEnv(env string) (string, error) {
	value, exists := os.LookupEnv(env)
	if !exists {
		return "", fmt.Errorf("required environment variable '%s' is missing", env)
	}

	if value == "" {
		return "", fmt.Errorf("required environment variable '%s' cannot be empty", env)
	}

	return value, nil
}
