package main

import (
	"fmt"
	"reflect"
)

type Model interface {
	Test()
}

type A struct {
	S int
}

func (a A)Test() {

}

func test() {
	var m Model = A{43}
	rv := reflect.ValueOf(m)

	fmt.Println(rv.Kind().String())
	fmt.Println(rv.Field(0).CanAddr())
	fmt.Println(rv.Field(0).CanSet())
	mm := m.(A)
	fmt.Println(mm)
}


func main() {
	//app := SetUp()
	//err := app.Run()
	//if err != nil {
	//	return
	//}
	test()
}
