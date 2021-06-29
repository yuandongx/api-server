package mysql

import "log"

func display(v ...interface{}) {
	if mysql_log_on {
		log.Println(v...)
	}
}

func Display(v ...interface{}) {
	display(v...)
}