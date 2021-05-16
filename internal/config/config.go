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
	for from, nodes := range config.Graph {
		for to := range nodes {
			if _, ok := config.Graph[to]; !ok {
				return config, fmt.Errorf(
					"node \"%s\" has path to missing node \"%s\"",
					from,
					to,
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
