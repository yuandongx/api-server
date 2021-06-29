package mysql

import "runtime"

func init() {
	if runtime.GOOS != "windows" {
		mysqlDns = "root:admin@12345@tcp(127.0.0.1:3306)/mysql"
	}
}
