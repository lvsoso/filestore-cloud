package db

import (
	"log"

	mydb "github.com/moxiaomomo/filestore-cloud/db/mysql"
)

func UserRegister(phone string, pwd string) bool {
	stmt, _ := mydb.DBConn().Prepare("insert into tbl_user (`phone`,`user_pwd`) values (?, ?)")
	defer stmt.Close()

	ret, err := stmt.Exec(phone, pwd)
	if err != nil {
		log.Printf("insert data error: %v\n", err)
		return false
	}
	if rowsAffected, err := ret.RowsAffected(); nil == err && rowsAffected > 0 {
		return true
	}
	return false
}

func UserLogin(phone string, pwd string) bool {
	stmt, _ := mydb.DBConn().Prepare("select * from tbl_user where phone=? limit 1")
	defer stmt.Close()

	rows, err := stmt.Query(phone)
	if err != nil {
		log.Printf("query data error: %v\n", err)
		return false
	} else if rows == nil {
		log.Printf("phone not found: %v\n", phone)
		return false
	}

	parsedRows := mydb.ParseRows(rows)
	if len(parsedRows) > 0 && string(parsedRows[0]["user_pwd"].([]byte)) == pwd {
		return true
	}
	return false
}

func UserUpdateToken(phone string, token string) bool {
	stmt, _ := mydb.DBConn().Prepare("replace into tbl_user_token (`user_id`,`token`) values (?, ?)")
	defer stmt.Close()

	ret, err := stmt.Exec(phone, token)
	if err != nil {
		log.Printf("insert token error: %v\n", err)
		return false
	}
	if rowsAffected, err := ret.RowsAffected(); nil == err && rowsAffected > 0 {
		return true
	}
	return false
}

func TokenValid(phone string, token string) bool {
	stmt, _ := mydb.DBConn().Prepare("select 1 from tbl_user_token where user_id=? and token=? limit 1")
	defer stmt.Close()

	rows, err := stmt.Query(phone, token)
	if err != nil {
		log.Printf("query token error: %v\n", err)
		return false
	} else if rows == nil || !rows.Next() {
		return false
	}

	return true
}
