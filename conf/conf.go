package conf

import (
	"io/ioutil"
	"sync"

	"github.com/xf005/logger"
	"gopkg.in/yaml.v3"
)

type Conf struct {
	Db Datasource `yaml:"datasource"`
}

type Datasource struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Name    string
	User    string
	Pass    string
	LogMode bool `yaml:"logMode"`
}

var (
	once sync.Once
	conf *Conf
)

func NewConf() *Conf {
	once.Do(func() {
		logger.Info("conf init...")
		file, err := ioutil.ReadFile("./conf/app.yml")
		if err != nil {
			logger.Error(err.Error())
		}
		var c Conf
		err = yaml.Unmarshal(file, &c)
		if err != nil {
			logger.Error(err.Error())
		}
		conf = &c
	})
	return conf
}
