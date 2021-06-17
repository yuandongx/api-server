package host

import "ping/com/mysql"

func addHost(h interface{}) (id int64, err error) {
	return mysql.Insert(h)
}

func getAll(obj mysql.Object) []interface{} {
	return mysql.Fetch(obj)
}

func read(object mysql.Object) interface{} {
	mysql.Read(object)
	return object
}
