package mysql

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func tableName(s string) string {
	return fmt.Sprintf("app-%s-v1", strings.ToLower(s))
}

func convert(kind reflect.Kind, field reflect.Value) (value interface{}) {
	switch kind {
	case reflect.Int:
		value = field.Int()
	case reflect.Int8:
		value = field.Int()
	case reflect.Int16:
		value = field.Int()
	case reflect.Int32:
		value = field.Int()
	case reflect.Int64:
		value = field.Int()
	case reflect.Bool:
		value = 0
		if field.Bool() {
			value = 1
		}
	case reflect.String:
		value = field.String()
	case reflect.Struct:
		value = field.String()
	case reflect.Float64:
		value = field.Float()
	case reflect.Float32:
		value = field.Float()
	}
	return
}

func (m *Model) init(v interface{}) {
	reft := reflect.TypeOf(v)
	refv := reflect.ValueOf(v)
	m.name = reft.Elem().Name()
	count := 0
	num := reft.Elem().NumField()
	columns := make([]string, 0)
	columns_ := make([]string, 0)
	fieldKinds := make([]reflect.Kind, 0)
	indexs := make([]int, 0)
	values := make([]interface{}, 0)
	m.ranges = make(map[string]int)
	for i := 0; i < num; i++ {
		field := reft.Elem().Field(i)
		if tag, ok := field.Tag.Lookup("sql"); ok {
			columns = append(columns, fmt.Sprintf("`%s`", tag))
			columns_ = append(columns_, tag)
			kind := field.Type.Kind()
			fieldKinds = append(fieldKinds, kind)
			indexs = append(indexs, i)
			values = append(values, convert(kind, refv.Elem().Field(i)))
			count++
			if strings.ToLower(tag) == "id" || strings.ToLower(field.Name) == "id" {
				m.id = refv.Elem().Field(i).String()
			}
			if l, ok := field.Tag.Lookup("len"); ok {
				if i, e := strconv.Atoi(l); e == nil {
					m.ranges[tag] = i
					display("lenght-->", i, tag)
				}
			}

		} else {
			continue
		}
	}
	m.fieldKinds = fieldKinds
	m.columns = columns
	m.count = count
	m.indexs = indexs
	m.table = tableName(m.name)
	m.values = values
	m.fields = columns_
}

//CloneRow 生成一行空的指针列表
func (m *Model) CloneRow() (row []interface{}) {
	row = make([]interface{}, m.count)
	for i, k := range m.fieldKinds {
		switch k {
		case reflect.Int:
			row[i] = new(int)
		case reflect.Int8:
			row[i] = new(int8)
		case reflect.Int16:
			row[i] = new(int16)
		case reflect.Int32:
			row[i] = new(int32)
		case reflect.Int64:
			row[i] = new(int64)
		case reflect.Bool:
			row[i] = new(int)
		case reflect.String:
			row[i] = new(string)
		case reflect.Struct:
			row[i] = new(string)
		case reflect.Float64:
			row[i] = new(float64)
		case reflect.Float32:
			row[i] = new(float32)
		}
	}
	return
}

func (m *Model) SetColumn(v interface{}, values []interface{}) {
	refv := reflect.ValueOf(v)
	display("SetColumn values")
	for i, v := range m.indexs {
		kind := m.fieldKinds[i]
		value := values[i]
		display("SetColumn: ", v, kind, value, refv.Elem().Field(i).CanSet())
		var ok bool
		switch kind {
		case reflect.Int:
			_, ok = value.(*int)
		case reflect.Int8:
			_, ok = value.(*int8)
		case reflect.Int16:
			_, ok = value.(*int16)
		case reflect.Int32:
			_, ok = value.(*int32)
		case reflect.Int64:
			_, ok = value.(*int64)
		case reflect.Bool:
			_, ok = value.(*int)
		case reflect.String:
			_, ok = value.(*string)
		//case reflect.Struct:
		//	_, ok = value.(struct)
		case reflect.Float64:
			_, ok = value.(*float64)
		case reflect.Float32:
			_, ok = value.(*float32)
		}
		if ok {
			if reflect.Bool == kind {
				v := *value.(*int)
				if v > 0 {
					refv.Elem().Field(i).SetBool(true)
				} else {
					refv.Elem().Field(i).SetBool(false)
				}
			} else {
				refv.Elem().Field(i).Set(reflect.ValueOf(value).Elem())
			}
		}
	}
}
