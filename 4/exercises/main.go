package main

import (
	"html/template"
	"os"
)

type User struct {
	Name   string
	Dog    string
	Months []string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name:   "Pero",
		Dog:    "Miki",
		Months: []string{"Jan", "Feb", "Mar", "Apr"},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
