package db

import (
	"fmt"
	"testing"
)

func TestConfiguration(t *testing.T) {
	Configuration()
	fmt.Println(conf)
	cfg := defaultDbConfig(conf.Database["db"])
	fmt.Println(cfg)
	fmt.Println(conf.Database["db"].Debug == true)
}
