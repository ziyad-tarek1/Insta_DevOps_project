package main

import (
	"database/sql"
	"fmt"
	"os"

	mysqldriver "github.com/go-sql-driver/mysql"
)

var _connection *sql.DB

func init() {
	mysqlConfig := &mysqldriver.Config{
		User:                 os.Getenv("MYSQL_USER"),
			Passwd:               os.Getenv("MYSQL_PASS"),
			Net:                  "tcp",
			Addr:                 fmt.Sprintf("%s:%s", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT")),
			AllowNativePasswords: true,
			ParseTime:            true,
	}

	var err error
	_connection, err = sql.Open("mysql", mysqlConfig.FormatDSN())
	if err!= nil {
		panic(err)
	}

	_, err = _connection.Exec(`CREATE DATABASE IF NOT EXISTS internship`)
	if err!= nil {
		panic(err)
	}

	mysqlConfig.DBName = `internship`

	_, err = _connection.Exec(`USE internship`)
	if err!= nil {
		panic(err)
	}

	schema()
}

func schema() {
	_, err := _connection.Exec(`CREATE TABLE IF NOT EXISTS stuff (
		id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		created_at DATETIME NOT NULL
	)`)

	if err!= nil {
		panic(err)
	}
}