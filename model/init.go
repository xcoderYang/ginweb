package model

import "database/sql"

func init() {
	db, err := sql.Open("mysql", "root:yangxuechao123@/medicalsystem?charset=utf8")
}
