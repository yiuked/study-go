package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

type Row map[string]string
type Rows [] Row

func init() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(localhost:3308)/pmp?charset=utf8")
	if err != nil {
		fmt.Print(err)
	}

	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}


func GetOne(query string, args ...interface{}) Row {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if rows == nil {
		return nil
	}

	// 获取查询到的行名，返回了一个[]string
	cols, err := rows.Columns()

	result := Row{}

	rawResult := make([][]byte, len(cols))
	dest := make([]interface{}, len(cols))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	if rows.Next() {
		// 传一个内存地址数组进行切割后，分散传入,每个内存地址保存一个字段值。
		err = rows.Scan(dest...)
		// 将存在内存地址中的值，传给result
		for i, raw := range rawResult {
			if raw == nil {
				result[cols[i]] = ""
			} else {
				result[cols[i]] = string(raw)
			}
		}
	} else {
		return nil
	}

	_ = err
	return result
}

func GetAll(query string, args ...interface{}) Rows {
	rows, err := db.Query(query, args...)

	if rows == nil {
		return nil
	}

	cols, err := rows.Columns()

	rawResult := make([][]byte, len(cols))
	result := Rows{}
	dest := make([]interface{}, len(cols))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		err = rows.Scan(dest...)

		tmpResult := make(Row, len(cols))

		for i, raw := range rawResult {
			if raw == nil {
				tmpResult[cols[i]] = ""
			} else {
				tmpResult[cols[i]] = string(raw)
			}
		}
		result = append(result, tmpResult)
	}

	_ = err
	return result
}
