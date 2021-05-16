package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func NewConfig() (Config, error) {
	var config Config
	config.DSN = os.Getenv("WT_DSN")
	if config.DSN == "" {
		return config, fmt.Errorf("missing env \"WT_DSN\"")
	}
	config.Node = os.Getenv("WT_NODE")
	if config.Node == "" {
		return config, fmt.Errorf("missing env \"WT_NODE\"")
	}
	bytes, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return config, err
	}
	if err := yaml.Unmarshal(bytes, &config.Graph); err != nil {
		return config, err
	}
	if _, ok := config.Graph[config.Node]; !ok {
		return config, fmt.Errorf("missing path for node \"%s\"", config.Node)
	}
	for source, targets := range config.Graph {
		for target := range targets {
			if source == target {
				return config, fmt.Errorf("node \"%s\" has path to itself", source)
			}
			if _, ok := config.Graph[target]; !ok {
				return config, fmt.Errorf(
					"node \"%s\" has path to missing node \"%s\"",
					source,
					target,
				)
			}
		}
	}
	return config, nil
}

type Config struct {
	DSN   string
	Node  string
	Graph map[string]Set
}
