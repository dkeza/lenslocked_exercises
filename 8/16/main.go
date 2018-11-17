package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "lenslocked_dev"
)

type User struct {
	gorm.Model
	Name   string
	Email  string `gorm:"not null;unique_index"`
	Color  string
	Orders []Order
}

type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
}

func main() {
	var (
		psqlInfo string
	)
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.LogMode(true)
	//db.DropTableIfExists(&User{})
	db.AutoMigrate(&User{}, &Order{})

	var users []User
	if err := db.Preload("Orders").Find(&users).Error; err != nil {
		panic(err)
	}

	fmt.Println(users)
	// createOrder(db, u, 1001, "Desc 1")
	// createOrder(db, u, 5000, "Desc 2")
	// createOrder(db, u, 998, "Desc 3")

}

func createOrder(db *gorm.DB, user User, amount int, desc string) {
	err := db.Create(&Order{
		UserID:      user.ID,
		Amount:      amount,
		Description: desc,
	}).Error
	if err != nil {
		panic(err)
	}
}
