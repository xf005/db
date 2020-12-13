package conf

import (
	"fmt"
	"testing"
)

type code struct {
	Num int
}

func TestConf(t *testing.T) {
	fmt.Println(NewConf().Server.Port)
	//// int test
	//fmt.Println("----------int test")
	//c := NewCache(50, "test", "int")
	//for i := 0; i < 10; i++ {
	//	c.Set(fmt.Sprint(i), i)
	//}
	//var a int
	//err := c.Get("5", &a)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(a)
	//c.Set("5", 10000000)
	//var ttt []int
	//err = c.List(&ttt)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//for _, kk := range ttt {
	//	fmt.Println(kk)
	//}
	//fmt.Println("keys len:", c.Len())
	//
	//fmt.Println("----------string test")
	//// int test
	//c = NewCache(5, "test", "string")
	//for i := 0; i < 10; i++ {
	//	c.Set(fmt.Sprint(i), fmt.Sprint(i))
	//}
	//var as string
	//err = c.Get("5", &as)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(as)
	//c.Set("5", fmt.Sprint(10000000))
	//var ttts []string
	//err = c.List(&ttts)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//for _, kk := range ttts {
	//	fmt.Println(kk)
	//}
	//fmt.Println("keys len:", c.Len())
	//
	////
	//fmt.Println("----------struct test")
	//c = NewCache(5, "test", "struct")
	//for i := 0; i < 10; i++ {
	//	c.Set(fmt.Sprint(i), &code{Num: i})
	//}
	//var aa code
	//err = c.Get("5", &aa)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(aa.Num)
	//fmt.Println(c.Exist("5"))
	//c.Delete("6")
	//c.Set("5", &code{Num: 10000000})
	//var tttaa []code
	//err = c.List(&tttaa)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//for _, kk := range tttaa {
	//	fmt.Println(kk.Num)
	//}
	//fmt.Println("keys len:", c.Len())
	//fmt.Println(c.Exist("5"))
	//fmt.Println(c.Exist("6"))
}
