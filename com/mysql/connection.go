package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"ping/com"
)

func getConnection() (driver *sql.DB) {
	if connection != nil {
		if err := connection.Ping(); err == nil {
			return connection
		}
	}
	driver, err := sql.Open("mysql", "root:admin@12345@tcp(81.70.9.203:3306)/xuyuandong")
	if err != nil {
		com.display("数据库连接失败！")
		com.display(err)
		return nil
	}
	com.display("db connected!")
	return driver
}

func exec(driver *sql.DB, prepare string, values []interface{}) bool {
	stmt, err := driver.Prepare(prepare)
	if err != nil {
		com.display(err)
		return false
	}
	defer stmt.Close()
	stmt.Exec(values...)
	return true
}
