package main

import (
	"flag"
	"log"

	"github.com/mballantyne3/Squad_up.git/api"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	firstName   string
	lastName    string
	phoneNumber string
	Age         uint
}

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "server address")
	flag.Parse()
	server := api.NewServer(*listenAddr)
	log.Fatal(server.Start())

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create User
	db.Create(&User{firstName: "", lastName: "", phoneNumber: "", Age: })

	// Read
	var user User
	db.First(&user, 1)           // find user with integer primary key
	db.First(&user.lastName, "") // find user with last name ""

	// Update - update users first name
	db.Model(&user).Update("age", 25)
	// Update - update multiple fields
	db.Model(&user).Updates(&User{firstName: "Mary"}) // non-zero fields
	db.Model(&user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete user
	db.Delete(&user, 1)
}
