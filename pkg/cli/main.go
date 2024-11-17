package cli

import (
	"errors"
	"github.com/pmoscode/go-common/cli"
	"log"
)

type Options struct {
	DockerHubRepository  *string
	HelmChartPath        *string
	ExcludeVersions      *string
	FailOnExistingUpdate *bool
	Debug                *bool
}

func getCliOptionsParameters() *Options {
	dockerHubRepository := cli.NewParameter[string]("docker-hub-repo", "", "DockHub repo to check tag versions", "HCUC_DOCKER_HUB_REPOSITORY")
	helmChartPath := cli.NewParameter[string]("helm-chart-path", ".", "Helm chart to check for updates", "HCUC_HELM_CHART_PATH")
	excludeVersions := cli.NewParameter[string]("exclude-versions", "", "Versions to exclude from check (on multiple versions: separated by comma)", "HCUC_EXCLUDE_VERSIONS")
	failOnExistingUpdate := cli.NewParameter[bool]("fail-on-update", false, "Return exit code 1, if update is available", "HCUC_FAIL_ON_EXISTING_UPDATE")
	debug := cli.NewParameter[bool]("debug", false, "Enable debug outputs", "HCUC_DEBUG")

	cliManager := cli.New()
	cliManager.AddStringParameter(dockerHubRepository)
	cliManager.AddStringParameter(helmChartPath)
	cliManager.AddStringParameter(excludeVersions)
	cliManager.AddBoolParameter(failOnExistingUpdate)
	cliManager.AddBoolParameter(debug)

	cliManager.Parse()

	return &Options{
		DockerHubRepository:  dockerHubRepository.GetValue(),
		HelmChartPath:        helmChartPath.GetValue(),
		ExcludeVersions:      excludeVersions.GetValue(),
		FailOnExistingUpdate: failOnExistingUpdate.GetValue(),
		Debug:                debug.GetValue(),
	}
}

func GetCliOptions() *Options {
	cliOptions := getCliOptionsParameters()

	if *cliOptions.DockerHubRepository == "" {
		log.Fatal(errors.New("parameter 'docker-hub-repo' is required"))
	}
	if *cliOptions.HelmChartPath == "" {
		log.Fatal(errors.New("parameter 'helm-chart-path' is required"))
	}

	return cliOptions
}
