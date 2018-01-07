package goTestDeps

import yaml "gopkg.in/yaml.v2"

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func Unmarshall(data []byte, t *T) error {
	err := yaml.Unmarshal(data, t)
	return err
}
