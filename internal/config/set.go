package config

import (
	"gopkg.in/yaml.v3"
)

type Set map[string]struct{}

func NewSet(items []string) Set {
	set := make(Set, len(items))
	for _, item := range items {
		set[item] = struct{}{}
	}
	return set
}

func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s *Set) UnmarshalYAML(node *yaml.Node) error {
	var items []string
	if err := node.Decode(&items); err != nil {
		return err
	}
	*s = NewSet(items)
	return nil
}
