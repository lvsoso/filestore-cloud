package db

import (
	"log"

	mydb "github.com/moxiaomomo/distributed-fileserver/db/mysql"
)

func UserRegister(user string, pwd string) bool {
	stmt, _ := mydb.DBConn().Prepare("insert into tbl_user (`user_name`,`user_pwd`) values (?, ?)")
	defer stmt.Close()

	ret, err := stmt.Exec(user, pwd)
	if err != nil {
		log.Printf("insert data error: %v\n", err)
		return false
	}
	if rowsAffected, err := ret.RowsAffected(); nil == err && rowsAffected > 0 {
		return true
	}
	return false
}
