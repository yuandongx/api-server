package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func getConnection() (driver *sql.DB) {
	if connection != nil {
		if err := connection.Ping(); err == nil {
			return connection
		}
	}
	driver, err := sql.Open("mysql", mysqlDns)
	if err != nil {
		display("数据库连接失败！")
		display(err)
		return nil
	}
	display("db connected!")
	return driver
}

func exec(driver *sql.DB, prepare string, values []interface{}) bool {
	stmt, err := driver.Prepare(prepare)
	if err != nil {
		display(err)
		return false
	}
	defer stmt.Close()
	stmt.Exec(values...)
	return true
}
