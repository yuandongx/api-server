package mysql

import (
	"fmt"
	"reflect"
	"strings"
)

func Insert(v interface{}) (id int64, err error) {
	m := Model{}
	m.init(v)
	values := m.values
	_columns := make([]string, m.count)
	columns := make([]string, m.count)
	for i := 0; i < m.count; i++ {
		_columns[i] = "?"
		columns[i] = fmt.Sprintf("`%s`", m.fields[i])
	}
	sqlString := fmt.Sprintf(insert_sql, m.table, strings.Join(columns, ", "), strings.Join(_columns, ", "))
	db := getConnection()
	result, err := db.Exec(sqlString, values...)
	if err == nil {
		id, err = result.LastInsertId()
	}
	return
}

func Create(v interface{}) (int64, error) {
	m := Model{}
	m.init(v)
	columns := make([]string, 0)
	id := ""
	for i, kind := range m.fieldKinds {
		c := m.fields[i]
		switch kind {
		case reflect.String:
			if r, ok := m.ranges[c]; ok {
				if r > 256 {
					columns = append(columns, fmt.Sprintf("`%s` text", c))
				} else {
					columns = append(columns, fmt.Sprintf("`%s` varchar(%d)", c, r))
				}
			} else {
				columns = append(columns, fmt.Sprintf("`%s` varchar(256)", c))
			}
		case reflect.Int, reflect.Int64, reflect.Int16, reflect.Int8, reflect.Int32, reflect.Bool:
			if strings.ToLower(c) == "id" {
				id = c
				columns = append(columns, fmt.Sprintf("`%s` int NOT NULL AUTO_INCREMENT", c))
			} else {

				columns = append(columns, fmt.Sprintf("`%s` int", c))
			}
		case reflect.Float64, reflect.Float32:
			columns = append(columns, fmt.Sprintf("`%s` float", c))
		case reflect.Struct:
			columns = append(columns, fmt.Sprintf("`%s` text", c))

		}
	}
	if id != "" {
		columns = append(columns, fmt.Sprintf("PRIMARY KEY ( `%s` )", id))
	}
	sqlstring := fmt.Sprintf(creat_table, m.table, strings.Join(columns, ", "))
	sqlDrop := fmt.Sprintf("DROP TABLE IF EXISTS `%s`;", m.table)
	db := getConnection()
	db.Exec(sqlDrop)
	result, err := db.Exec(sqlstring)
	display(sqlstring)
	if err == nil {
		return result.LastInsertId()
	}
	return 0, err
}

//指针数组为每一行数据
func Fetch(object Object) (values []interface{}) {
	m := Model{}
	m.init(object)
	values = make([]interface{}, 0)
	sqlLine := fmt.Sprintf(select_sql, strings.Join(m.columns, ", "), m.table)
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
	sqlLine := fmt.Sprintf(select_sql, strings.Join(m.columns, ", "), m.table)
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

func Update(v interface{}) (id int64, err error) {
	id = 0
	err = nil
	return
}
