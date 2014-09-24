package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

type Test struct {
	Id   int64
	Test string
}

func main() {
	fmt.Println("Hello world!")

	db, err := gorm.Open("postgres", "postgres://Niklas:nxdvqm@127.0.0.1:5432/Niklas?sslmode=disable")

	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	db.DB()

	//db.CreateTable(&Test{})

	db.Save(&Test{Test: "hej2"})

	fmt.Println("Hello world 2!")
}
