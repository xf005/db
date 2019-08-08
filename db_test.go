package db

import (
	"fmt"
	"testing"
)

func TestDb(t *testing.T) {
	sum := 0
	if err := DB().Raw("select 1+1").Count(&sum).Error; err != nil {
		fmt.Println("查询出错.")
	} else {
		fmt.Println("查询结果.", sum)
	}
}
