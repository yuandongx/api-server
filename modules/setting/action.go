package setting

import (
	"ping/com/mysql"
	"strconv"
)

func (a *AccessCredentials) add() (int64, error) {
	return mysql.Insert(a)
}

func fetchAll(a AccessCredentials) []interface{} {
	return mysql.Fetch(&a)
}

func read(id string) interface{} {
	a := AccessCredentials{}
	if i64, err := strconv.ParseInt(id, 10, 64); err == nil {
		a.Id = i64
		mysql.Read(&a)
	}
	return a
}

func CreateAccessCredentialTable() error {
	_, err := mysql.Create(&AccessCredentials{})
	return err
}
