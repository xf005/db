package db

import (
	"fmt"
	"testing"
)

func TestDb(t *testing.T) {
	for i := 0; i < 4; i++ {
		result := make(map[string]interface{})
		if err := DB().Raw("select now() as a").Find(&result).Error; err != nil {
			fmt.Println("查询出错.")
		} else {
			fmt.Println("查询结果.", result["a"])
		}
	}
	var confs []Conf
	if err := DB().Model(&Conf{}).Find(&confs).Error; err == nil {
		for _, conf := range confs {
			fmt.Println(conf.Keywords)
		}
	}
}

type Conf struct {
	Id       int64
	Title    string
	Keywords string
}
