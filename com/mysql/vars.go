package mysql

import (
	"database/sql"
	"reflect"
)

var (
	mysqlDns   string  = "root:admin@12345@tcp(81.70.9.203:3306)/mysql"
	mysqlLogOn bool    = true
	connection *sql.DB = nil
	creatTable         = "CREATE TABLE `%s` (%s) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;"
	insertSql          = "INSERT INTO `%s` (%s) VALUES (%s);"
	selectSql          = "SELECT %s FROM `%s`;"
)

type Model struct {
	name       string         // 对象名称
	_type      string         // 对象类型信息
	indexs     []int          // 索引信息
	id         string         // 单个查找的ID信息
	fieldKinds []reflect.Kind //字段类型
	columns    []string       //列名
	fields     []string       // 字段名
	values     []interface{}  // 列值
	table      string         // 表名
	count      int            //字段数量
	ranges     map[string]int // 字符串长度 如（0, 254）
}

type Object interface {
	Clone() Object
}
