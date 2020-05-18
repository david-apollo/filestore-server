package mysql

import (
	"os"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


var db *sql.DB


func init() {
	db, _ = sql.Open("mysql", "root:chainext123456@tcp(47.74.235.176:5557)/fileserver?charset=utf8")
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql, err:" + err.Error())
		os.Exit(1)
	}
}


// DBConn : 返回数据库连接对象
func DBConn() *sql.DB {
	return db
}