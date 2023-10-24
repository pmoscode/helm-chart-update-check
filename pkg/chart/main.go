package chart

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
)

type Chart struct {
	chartPath string
	root      *yaml.Node
}

func (c *Chart) Version() string {
	return strings.TrimSpace(c.searchNode("version"))
}

func (c *Chart) AppVersion() string {
	return strings.TrimSpace(c.searchNode("appVersion"))
}

func (c *Chart) UpdateAppVersion() {

}

func (c *Chart) searchNode(nodeName string) string {
	//var keyNode *yaml.Node
	var valNode *yaml.Node

	for i := 0; i < len(c.root.Content); i += 2 {
		node := c.root.Content[i]
		if node.Kind == yaml.ScalarNode && node.Value == nodeName {
			//keyNode = c.root.Content[i]
			valNode = c.root.Content[i+1]
		}
	}

	//locationSectionKey, _ := yaml.Marshal(keyNode)
	//fmt.Println(string(locationSectionKey))
	locationSectionVal, _ := yaml.Marshal(valNode)

	return string(locationSectionVal)
}

func NewChart(chartPath string) *Chart {
	b, err := os.ReadFile(chartPath + "/Chart.yaml")
	if err != nil {
		log.Fatalf("Cannot open file: %v", err)
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
