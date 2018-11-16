package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

	type User struct {
		gorm.Model
		Name  string
		Email string `gorm:"not null;unique_index"`
	}

	db.LogMode(true)
	//db.DropTableIfExists(&User{})
	db.AutoMigrate(&User{})

	name, email := getInfo()
	u := User{
		Name:  name,
		Email: email,
	}

	if err = db.Create(&u).Error; err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", u)
}

func getInfo() (name, email string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Name?")
	name, _ = reader.ReadString('\n')
	fmt.Println("E-Mail?")
	email, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	return name, email
}
