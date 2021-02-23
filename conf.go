package db

import (
	"io/ioutil"
	"sync"

	"github.com/xf005/logger"
	"gopkg.in/yaml.v3"
)

type Database struct {
	Alias map[string]Datasource
}

type Datasource struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Db      string `yaml:"db"`
	User    string `yaml:"user"`
	Pass    string `yaml:"pass"`
	LogMode bool   `yaml:"logMode"`
}

func (e *Database) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var alias map[string]Datasource
	if err := unmarshal(&alias); err != nil {
		// Here we expect an error because a boolean cannot be converted to a
		// a MajorVersion
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	}
	e.Alias = alias
	return nil
}

var (
	once sync.Once
	conf map[string]Database
)

func DataSourceConf() map[string]Database {
	once.Do(func() {
		logger.Info("conf init...")
		file, err := ioutil.ReadFile("./conf/app.yml")
		if err != nil {
			logger.Error(err.Error())
		}
		if err := yaml.Unmarshal(file, &conf); err != nil {
			logger.Error(err.Error())
		}
	})
	return conf
}
