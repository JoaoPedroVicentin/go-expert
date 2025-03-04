package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// db.Create(&Product{
	// 	Name:  "PC",
	// 	Price: 3550.99,
	// })

	products := []Product{
		{Name: "Notebook", Price: 1990.5},
		{Name: "Smartphone", Price: 999.9},
		{Name: "Tablet", Price: 599.9},
	}

	db.Create(&products)
}
