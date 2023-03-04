package main

import (
	"bpjs/config"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/thoas/go-funk"
	"gorm.io/driver/mysql"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func genRandomString(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func stubUsers(b *testing.B) (orders []*Orders) {
	for i := 0; i < b.N; i++ {
		order := &Orders{
			Customer:  genRandomString(1000),
			Quantity:  rand.Intn(1000),
			Price:     rand.Intn(1000),
			Timestamp: time.Now().Format(config.AppTLayout),
		}
		orders = append(orders, order)
	}

	return orders
}

func BenchmarkOrmCreate(b *testing.B) {
	dsn := "root:root@tcp(localhost:3309)/bpjs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
	// defer db.Close()

	orders := stubUsers(b)
	for _, order := range orders {
		db.Create(order)
	}
}

func BenchmarkCreate(b *testing.B) {
	dsn := "root:root@tcp(localhost:3309)/bpjs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
	// defer db.Close()

	orders := stubUsers(b)
	tx := db.Begin()
	valueStrings := []string{}
	valueArgs := []interface{}{}
	for _, order := range orders {
		valueStrings = append(valueStrings, "(?, ?, ?, ?)")
		valueArgs = append(valueArgs, order.Customer)
		valueArgs = append(valueArgs, order.Quantity)
		valueArgs = append(valueArgs, order.Price)
		valueArgs = append(valueArgs, order.Timestamp)

	}

	stmt := fmt.Sprintf("INSERT INTO orders (customer, quantity, price, timestamp) VALUES %s", strings.Join(valueStrings, ","))
	err = tx.Exec(stmt, valueArgs...).Error
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
	}
	err = tx.Commit().Error
	if err != nil {
		fmt.Println(err)
	}
}

func BenchmarkBulkCreate(b *testing.B) {
	dsn := "root:root@tcp(localhost:3309)/bpjs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
	// defer db.Close()

	orders := stubUsers(b)
	size := 500
	tx := db.Begin()
	chunkList := funk.Chunk(orders, size)
	for _, chunk := range chunkList.([][]*Orders) {
		valueStrings := []string{}
		valueArgs := []interface{}{}
		for _, order := range chunk {
			valueStrings = append(valueStrings, "(?, ?, ?, ?)")
			valueArgs = append(valueArgs, order.Customer)
			valueArgs = append(valueArgs, order.Quantity)
			valueArgs = append(valueArgs, order.Price)
			valueArgs = append(valueArgs, order.Timestamp)
		}

		stmt := fmt.Sprintf("INSERT INTO orders (customer, quantity, price, timestamp) VALUES %s", strings.Join(valueStrings, ","))

		err = tx.Exec(stmt, valueArgs...).Error
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
		}
	}
	err = tx.Commit().Error
	if err != nil {
		fmt.Println(err)
	}
}

func benchmarkBulkCreate(size int, b *testing.B) {
	dsn := "root:root@tcp(localhost:3309)/bpjs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	// defer db.Close()

	orders := stubUsers(b)
	tx := db.Begin()
	chunkList := funk.Chunk(orders, size)
	for _, chunk := range chunkList.([][]*Orders) {
		valueStrings := []string{}
		valueArgs := []interface{}{}
		for _, order := range chunk {
			now := time.Now()
			valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?)")
			valueArgs = append(valueArgs, now)
			valueArgs = append(valueArgs, now)
			valueArgs = append(valueArgs, order.Customer)
			valueArgs = append(valueArgs, order.Quantity)
			valueArgs = append(valueArgs, order.Price)
			valueArgs = append(valueArgs, order.Timestamp)
		}

		stmt := fmt.Sprintf("INSERT INTO orders (customer, quantity, price, timestamp) VALUES %s", strings.Join(valueStrings, ","))
		err = tx.Exec(stmt, valueArgs...).Error
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
		}
	}
	err = tx.Commit().Error
	if err != nil {
		fmt.Println(err)
	}
}

func BenchmarkBulkCreateSize1(b *testing.B) {
	benchmarkBulkCreate(1, b)
}

func BenchmarkBulkCreateSize100(b *testing.B) {
	benchmarkBulkCreate(100, b)
}

func BenchmarkBulkCreateSize500(b *testing.B) {
	benchmarkBulkCreate(500, b)
}

func BenchmarkBulkCreateSize1000(b *testing.B) {
	benchmarkBulkCreate(1000, b)
}
