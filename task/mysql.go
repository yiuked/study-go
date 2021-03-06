package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)

var db *sql.DB

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

func GetOne(query string, args ...interface{}) []string {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if rows == nil {
		return nil
	}

	cols, err := rows.Columns()

	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))
	dest := make([]interface{}, len(cols))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	if rows.Next() {
		err = rows.Scan(dest...)
		for i, raw := range rawResult {
			if raw == nil {
				result[i] = ""
			} else {
				result[i] = string(raw)
			}
		}
	} else {
		return nil
	}

	_ = err
	return result
}

type Row map[string]string
type Rows [] Row

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
