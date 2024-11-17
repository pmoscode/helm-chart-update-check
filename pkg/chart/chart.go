package chart

import (
	"gopkg.in/yaml.v3"
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
