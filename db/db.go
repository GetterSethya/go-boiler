package db

import (
	"log"

	"github.com/GetterSethya/golang-gorm/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Database()(*gorm.DB, error){
	db,err := gorm.Open(sqlite.Open("./db/dev.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&models.Grocery{}); err != nil {
		log.Println(err)
	}


	return db, err
}