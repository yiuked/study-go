package main

import "database/sql"

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3309)/go?charset=utf8")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

func GetOne(query string, args ...interface{}) []string {
	rows, err := db.Query(query, args...)

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
