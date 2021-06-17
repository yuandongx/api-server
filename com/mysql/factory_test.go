package mysql

import (
	"fmt"
	"testing"
)

type SubTes struct {

}

type TestDeamo struct {
	T1 string `sql:"t1" orm:"a32"`
	T2 int `sql:"t2"`
	T3 int `sql:"t3"`
	T4 bool `sql:"t4"`
}

func (t *TestDeamo)Clone() (v Object) {
	v = &TestDeamo{}
	return
}

func TestInsert(t *testing.T) {
	test := TestDeamo{"123", 23, 43, true}
	Insert(test)
}

func TestCreate(t *testing.T) {
	test := TestDeamo{"123", 23, 43, true}
	Create(test)
}

func TestFetch(t *testing.T) {
	test := &TestDeamo{}
	result := Fetch(test)
	display("Test Fetch, result:", result)
	for i, v := range result {
		display(fmt.Sprintf("%d, %v", i, v))
	}
}
func TestRead(t *testing.T) {
	test := TestDeamo{}
	Read(&test)
}

func TestModel_SetColumn(t *testing.T) {
	test := TestDeamo{"123", 23, 43, true}
	m := Model{}
	m.init(test)
	m.SetColumn(&test, []interface{}{"434", 44, 5, false})
	fmt.Println("32323", test)
}