package db

import (
	"fmt"
	"testing"
)

func TestDb(t *testing.T) {
	result := make(map[string]interface{})
	if err := DB().Raw("select now() as a").Find(&result).Error; err != nil {
		fmt.Println("查询出错.")
	} else {
		fmt.Println("查询结果.", result["a"])
	}
	//
	var confs []News
	if err := DB().Model(&News{}).Find(&confs).Error; err == nil {
		for _, conf := range confs {
			fmt.Println(conf.Title)
		}
	}
}

type News struct {
	Id    int64
	Title string
}
