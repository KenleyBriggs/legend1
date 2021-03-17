package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func DataBase() {
	db, err := sql.Open("mysql", "root:***@/pricing")
	if err != nil {
		panic(err)
	}
	type SQLService struct {
		THREE_TWELVE int
		SIX_TWELVE   int
		TEN_TWELVE   int
	}
	rows, err := db.Query("SELECT THREE_TWELVE, SIX_TWELVE, TEN_TWELVE FROM metal")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		metalp := SQLService{}
		err = rows.Scan(&metalp.THREE_TWELVE, &metalp.SIX_TWELVE, &metalp.TEN_TWELVE)
		if err != nil {
			panic(err)
		}
		fmt.Println(metalp)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
