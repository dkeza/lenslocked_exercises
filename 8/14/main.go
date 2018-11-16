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
	password = "3000"
	dbname   = "lenslocked_dev"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
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
	db.AutoMigrate(&User{})

	var u User
	var users []User

	db.First(&u)

	fmt.Println(u)

	db.Last(&u)

	fmt.Println(u)

	db.Find(&users)

	fmt.Println(users)

	db.Where("id > ?", 1).Find(&users)

	fmt.Println(users)

}
