package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main () {
	fmt.Println("a"+"b"+"v")
	fmt.Println("a","b","c")

	db,err :=sql.Open("mysql","root:123321@tcp(127.0.0.1:3306)/leetcode?charset=utf8")
	if err!=nil {
		fmt.Println("cuowu",err)
		return
	}
	fmt.Println(db)

}
