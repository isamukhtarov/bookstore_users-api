package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	mysql_users_username = "root"
	mysql_users_password = "root"
	mysql_users_host     = "127.0.0.1:3306"
	mysql_users_schema   = "users_db"
)

var (
	Client *sql.DB
	username = mysql_users_username
	password = mysql_users_password
	host     = mysql_users_host
	schema   = mysql_users_schema
)

func init()  {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database configuration successfully")
}
