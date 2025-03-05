package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID       string `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         string `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	// SerialNumber SerialNumber
	gorm.Model
}

// type SerialNumber struct {
// 	ID        string `gorm:"primaryKey"`
// 	Number    string
// 	ProductID string
// }

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.Name = "Eletrônicos"
	tx.Debug().Save(&c)
	tx.Commit()

	// var product Product
	// var product2 Product
	// var products []Product

	// category := Category{
	// 	Name: "Cozinha",
	// }
	// db.Create(&category)

	// category2 := Category{
	// 	Name: "Eletrônicos",
	// }
	// db.Create(&category2)

	// db.Create(&Product{
	// 	Name:       "PC",
	// 	Price:      3550.99,
	// 	CategoryID: category.ID,
	// })

	// db.Create(&Product{
	// 	Name:       "Mouse",
	// 	Price:      150.99,
	// 	Categories: []Category{category, category2},
	// })

	// db.Create(&SerialNumber{
	// 	Number:    "123456",
	// 	ProductID: "1",
	// })

	// var categories []Category
	// err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	// if err != nil {
	// 	panic(err)
	// }
	// for _, category := range categories {
	// 	fmt.Println(category.Name, ":")
	// 	for _, product := range category.Products {
	// 		fmt.Println("- ", product.Name)
	// 	}
	// }

	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	// }

	// products := []Product{
	// 	{Name: "Notebook", Price: 1990.5},
	// 	{Name: "Smartphone", Price: 999.9},
	// 	{Name: "Tablet", Price: 599.9},
	// }

	// db.Create(&products)

	// db.First(&product, 1)
	// fmt.Println(product)
	// db.First(&product, "name = ?", "Notebook")
	// fmt.Println(product)

	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// db.Where("price > ?", 1000).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// db.Where("name LIKE ?", "%book%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// db.First(&product, 1)
	// product.Name = "PC"
	// db.Save(&product)

	// db.First(&product2, 1)
	// fmt.Println(product2.Name)
	// db.Delete(&product2)
}
