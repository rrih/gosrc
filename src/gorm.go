// https://qiita.com/daitasu/items/9428220810816972b5d5#gorm
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	- "github.com/jinzhu/dialects/mysql"
)

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := <user>
	PASS := <db pass>
	PROPTOCOL := tcp(<db ip address>:<port>)
	DBNAME := <db name>
	CONNECT := USER + ":" + PASS + "@" + PROPTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected:". &db)
	return db
}

func main() {
	db := gormConnect()
	defer db.Close()
	db.LogMode(true)
}