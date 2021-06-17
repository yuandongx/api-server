package mysql

import (
	"database/sql"
	"reflect"
)

var (
	mysql_log_on   bool    = true
	connection     *sql.DB = nil
	creat_table            = "CREATE TABLE `%s` (%s) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;"
	insert_sql             = "INSERT INTO `%s` (%s) VALUES (%s);"
	drop_if_exists         = "DROP TABLE IF EXISTS `%s`;"
	select_sql             = "SELECT %s FROM `%s`;"
)

type Model struct {
	name       string         // 对象名称
	_type      string         // 对象类型信息
	indexs     []int          // 索引信息
	id         string         // 单个查找的ID信息
	fieldKinds []reflect.Kind //字段类型
	columns    []string       //列名
	table      string         // 表名
	count      int            //字段数量
}

type Object interface {
	Clone() Object
}
