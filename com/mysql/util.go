package mysql

import "log"

func display(v ...interface{}) {
	if mysqlLogOn {
		log.Println(v...)
	}
}

func Display(v ...interface{}) {
	display(v...)
}
