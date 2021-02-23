package db

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestConf(t *testing.T) {
	var e map[string]Environment
	if err := yaml.Unmarshal([]byte(data), &e); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%#v\n", e)
	fmt.Println(e["development"].Versions["V1"].MimeTypes)
}

var data = `
---
development:
  skip-header-validation: true
  V1:
    current: "1.0.0"
    mime_types:
      - application/vnd.company.jk.identity+json;
      - application/vnd.company.jk.user+json;
      - application/vnd.company.jk.role+json;
      - application/vnd.company.jk.scope+json;
      - application/vnd.company.jk.test+json;
    skip-mime-type-validation: true
    skip-version-validation: true
  V2:
    current: "2.0.0"
    mime_types:
      - application/vnd.company.jk.identity+json;
      - application/vnd.company.jk.user+json;
      - application/vnd.company.jk.role+json;
      - application/vnd.company.jk.scope+json;
      - application/vnd.company.jk.test+json;

`

type MajorVersion struct {
	Current                string   `yaml:"current"`
	MimeTypes              []string `yaml:"mime_types"`
	SkipVersionValidation  bool     `yaml:"skip-version-validation"`
	SkipMimeTypeValidation bool     `yaml:"skip-mime-type-validation"`
}

type Environment struct {
	SkipHeaderValidation bool
	Versions             map[string]MajorVersion
}

func (e *Environment) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var params struct {
		SkipHeaderValidation bool `yaml:"skip-header-validation"`
	}
	if err := unmarshal(&params); err != nil {
		return err
	}
	var versions map[string]MajorVersion
	if err := unmarshal(&versions); err != nil {
		// Here we expect an error because a boolean cannot be converted to a
		// a MajorVersion
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	}
	e.SkipHeaderValidation = params.SkipHeaderValidation
	e.Versions = versions
	return nil
}
