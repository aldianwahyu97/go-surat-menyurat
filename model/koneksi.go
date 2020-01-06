package model

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func connect() (*sql.DB, error){
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/apsgolang_surat")
	if err != nil{
		return nil, err
	}
	fmt.Println("Berhasil terkoneksi ke database")
	return db,nil
}