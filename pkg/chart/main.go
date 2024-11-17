package chart

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/pmoscode/helm-chart-update-check/pkg/cli"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
)

func NewChart(chartPath string) *Chart {
	b, err := os.ReadFile(chartPath + "/Chart.yaml")
	if err != nil {
		log.Printf("Cannot open file: %v\n", err)
		log.Fatalf("Does path '%s' contain a 'Chart.yaml?", chartPath)
	}

	chart := &Chart{
		chartPath: chartPath,
	}

	var document yaml.Node
	err = yaml.Unmarshal(b, &document)
	if err != nil {
		log.Fatal("Cannot unmarshall document", err)
	}

	chart.root = document.Content[0]

	return chart
}

func FetchChartVersion(cliOptions *cli.Options) *semver.Version {
	chart := NewChart(*cliOptions.HelmChartPath)

	appVersion := strings.Trim(chart.AppVersion(), "\"")
	fmt.Println("Helm chart AppVersion:")
	fmt.Println(appVersion)

	v, err := semver.NewVersion(appVersion)
	if err != nil {
		log.Fatal("Problem creating semver: ", err)
	}
	return v
}
