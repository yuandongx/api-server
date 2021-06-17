package mysql

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)


func (m *Model) init(v interface{}) {
	reft := reflect.TypeOf(v)
	refv := reflect.ValueOf(v)
	m.name = reft.Elem().Name()
	count := 0
	num := reft.Elem().NumField()
	columns := make([]string, 0)
	fieldKinds := make([]reflect.Kind, 0)
	indexs := make([]int, 0)
	for i := 0; i < num; i++ {
		field := reft.Elem().Field(i)
		if tag, ok := field.Tag.Lookup("sql"); ok {
			columns = append(columns, fmt.Sprintf("`%s`", tag))
			kind := field.Type.Kind()
			fieldKinds = append(fieldKinds, kind)
			indexs = append(indexs, i)
			count++
			if strings.ToLower(tag) == "id" || strings.ToLower(field.Name) == "id" {
				m.id = refv.Elem().Field(i).String()
			}
		} else {
			continue
		}
	}
	m.fieldKinds = fieldKinds
	m.columns = columns
	m.count = count
	m.indexs = indexs
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

func Insert(v interface{})(id int64, err error) {
	vt := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)
	table := vt.Name()
	vaules := make([]interface{}, 0)
	colums := make([]string, 0)
	_colums := make([]string, 0)

	for i:=0; i<vv.NumField(); i++ {
		c := vt.Field(i).Tag.Get("sql")
		kind := vt.Field(i).Type.Kind()
		field := vv.Field(i)
		if c != "" {
			colums = append(colums, c)
			_colums = append(_colums, "?")
			switch kind {
			case reflect.Int:
				vaules = append(vaules, field.Int())
			case reflect.Int8:
				vaules = append(vaules, field.Int())
			case reflect.Int16:
				vaules = append(vaules, field.Int())
			case reflect.Int32:
				vaules = append(vaules,field.Int())
			case reflect.Int64:
				vaules = append(vaules, field.Int())
			case reflect.Bool:
				v := 0
				if field.Bool() {
					v = 1
				}
				vaules = append(vaules, v)
			case reflect.String:
				vaules = append(vaules, field.String())
			case reflect.Struct:
				vaules = append(vaules, field.String())
			case reflect.Float64:
				vaules = append(vaules, field.Float())
			case reflect.Float32:
				vaules = append(vaules, field.Float())
			}
		}

	}
	sqlString := fmt.Sprintf(insert_sql, table, strings.Join(colums, ", "), strings.Join(_colums, ", "))
	db := getConnection()
	display(vaules)
	result, err := db.Exec(sqlString, vaules...)
	if err == nil {
		id, err = result.LastInsertId()
	}
	return
}

func Create(v interface{}) bool {
	vt := reflect.TypeOf(v)
	table := vt.Name()
	n := vt.NumField()
	columns := make([]string, 0)
	for i:=0; i<n ;i++ {
		sql := vt.Field(i).Tag.Get("sql")
		if sql == "" {
			continue
		}
		t := vt.Field(i).Type.String()
		if t == "string" {
			length := vt.Field(i).Tag.Get("len")
			if length == "" {
				columns = append(columns, fmt.Sprintf("`%s` varchar(256)", sql))
			} else if n, err := strconv.Atoi(length); n < 256 && err == nil {
				columns = append(columns, fmt.Sprintf("`%s` varchar(%d)", sql, n))
			} else {
				columns = append(columns, fmt.Sprintf("`%s` text", sql))
			}
		} else if t == "bool" {
			columns = append(columns, fmt.Sprintf("`%s` int", sql))
		} else  if t == "int64" || t == "int32" || t == "int16" || t == "int8" || t == "int" {
			columns = append(columns, fmt.Sprintf("`%s` int", sql))
		} else {
			columns = append(columns, fmt.Sprintf("`%s` json", sql))
		}

	}
	sqlstring := fmt.Sprintf(creat_table, table, strings.Join(columns, ", "))
	db := getConnection()
	db.Exec(sqlstring, )
	display(sqlstring)
	return true
}

//指针数组为每一行数据
func Fetch(object Object) (values []interface{}) {
	m := Model{}
	m.init(object)
	values = make([]interface{}, 0)
	sqlLine := fmt.Sprintf(select_sql, strings.Join(m.columns, ", "), m.name)
	db := getConnection()
	rows, err := db.Query(sqlLine)
	if err == nil {
		for rows.Next() {
			row := m.CloneRow()
			rows.Scan(row...)
			fmt.Println(row)
			obj := object.Clone()
			m.SetColumn(obj, row)
			values = append(values, obj)
		}
	} else {
		display("Query all data failed:", err)
	}
	return values
}


func Read(v interface{}) {
	m := Model{}
	m.init(v)
	sqlLine := fmt.Sprintf(select_sql, strings.Join(m.columns, ", "), m.name)
	fmt.Println(sqlLine)
	db := getConnection()
	row := db.QueryRow(sqlLine)
	values := m.CloneRow()
	err := row.Scan(values...)
	if err != nil {
		display(err)
	}
	m.SetColumn(v, values)
	return
}

func Update(v interface{})(id int64, err error) {
	id = 0
	err = nil
	return
}