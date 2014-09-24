package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

type Email struct {
	Id     int64
	Email  string
	UserId int64
}

type User struct {
	Id     int64
	Name   string `sql:"type:varchar(100);"`
	Emails []Email
}

func main() {
	fmt.Println("Hello world!")

	db, err := gorm.Open("postgres", "postgres://Niklas:nxdvqm@127.0.0.1:5432/Niklas?sslmode=disable")

	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	db.DB()

	//db.DropTable(&Test{})

	db.CreateTable(&User{})
	db.CreateTable(&Email{})

	test := User{
		Name:   "hej2",
		Emails: []Email{{Email: "hej"}, {Email: "twa"}},
	}

	db.Create(&test)

	fmt.Println("Hello world 2!")
}
