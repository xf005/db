package db

import (
	"io/ioutil"
	"sync"

	"github.com/xf005/logger"
	"gopkg.in/yaml.v3"
)

type Datasource struct {
	Database map[string]Database
}

type Database struct {
	Dsn          string
	MaxIdleConns int
	MaxOpenConns int
	Debug        bool
}

// 默认设置
func defaultDbConfig(cfg Database) Database {
	newCfg := cfg
	if newCfg.MaxIdleConns == 0 {
		newCfg.MaxIdleConns = 10
	}
	if newCfg.MaxOpenConns == 0 {
		newCfg.MaxOpenConns = 20
	}
	return newCfg
}

var (
	syncOnce sync.Once
	conf     *Datasource
)

func Configuration() {
	syncOnce.Do(func() {
		logger.Info("db init...")
		file, err := ioutil.ReadFile("./conf.yml")
		if err != nil {
			logger.Error(err.Error())
		}
		var config Datasource
		if err := yaml.Unmarshal(file, &config); err != nil {
			logger.Error(err.Error())
		}
		conf = &config
	})
}
