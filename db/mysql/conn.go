package mysql

import (
	"database/sql"
	"log"
	"os"

	// using mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	err := db.Ping()
	if err != nil {
		log.Fatal("initialize mysql connection failed, err:" + err.Error())
		os.Exit(1)
	}
}

// 封装结果集
func RowResult(querySQL string) map[string]string {
	rows, err := db.Query(querySQL)
	defer rows.Close()
	checkErr(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]string)
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		checkErr(err)

		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
	}
	return record
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
