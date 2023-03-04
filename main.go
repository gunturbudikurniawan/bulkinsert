package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	ID        int
	Customer  string
	Quantity  int
	Price     int
	Timestamp string
}

func main() {
	dsn := "root:root@tcp(localhost:3309)/bpjs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connection to database is good")
	db.AutoMigrate(&Orders{})
}
