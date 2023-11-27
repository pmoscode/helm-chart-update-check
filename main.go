package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Masterminds/semver/v3"
	chart2 "github.com/pmoscode/helm-chart-update-check/pkg/chart"
	"github.com/pmoscode/helm-chart-update-check/pkg/dockerhub"
	"log"
	"os"
	"strings"
)

type CliOptions struct {
	dockerHubRepository  *string
	helmChartPath        *string
	failOnExistingUpdate *bool
	debug                *bool
}

func getCliOptions() CliOptions {
	dockerHubRepository := flag.String("docker-hub-repo", "", "DockHub repo to check tag versions")
	helmChartPath := flag.String("helm-chart-path", ".", "Helm chart to check for updates")
	failOnExistingUpdate := flag.Bool("fail-on-update", false, "Return exit code 1, if update is available")
	debug := flag.Bool("debug", false, "Enable debug outputs")

	flag.Parse()

	return CliOptions{
		dockerHubRepository:  dockerHubRepository,
		helmChartPath:        helmChartPath,
		failOnExistingUpdate: failOnExistingUpdate,
		debug:                debug,
	}
}

func main() {
	cliOptions := getCliOptions()

	if *cliOptions.dockerHubRepository == "" {
		log.Fatal(errors.New("parameter 'docker-hub-repo' is required"))
	}
	if *cliOptions.helmChartPath == "" {
		log.Fatal(errors.New("parameter 'helm-chart-path' is required"))
	}

	dockerHub := dockerhub.CreateDockerHub(*cliOptions.dockerHubRepository, *cliOptions.debug)
	versions := dockerHub.GetVersions()

	chart := chart2.NewChart(*cliOptions.helmChartPath)

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

		if *cliOptions.failOnExistingUpdate {
			os.Exit(1)
		}
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
//func findChildNode(value string, node *yaml.Node) *yaml.Node {
//	for _, v := range node.Content {
//		// If we found the value we are looking for, return it.
//		fmt.Printf("%+v", v)
//		fmt.Println()
//		if v.Value == value {
//			return v
//		}
//		// Otherwise recursively look more
//		if child := findChildNode(value, v); child != nil {
//			return child
//		}
//	}
//	return nil
//}
//
//func getRequiredEnv(env string) (string, error) {
//	value, exists := os.LookupEnv(env)
//	if !exists {
//		return "", fmt.Errorf("required environment variable '%s' is missing", env)
//	}
//
//	if value == "" {
//		return "", fmt.Errorf("required environment variable '%s' cannot be empty", env)
//	}
//
//	return value, nil
//}
