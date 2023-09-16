package main

import (
	"log"

	"github.com/GetterSethya/golang-gorm/controllers"
	"github.com/GetterSethya/golang-gorm/db"
	"github.com/gin-gonic/gin"
)

func main() {

	db, err := db.Database()
	if err != nil {
		log.Println(err)
	}

	db.DB()

	router := gin.Default()

	router.GET("/groceries",controllers.GetGroceries)
	router.GET("/grocery/:id",controllers.GetGrocery)
	router.POST("/grocery",controllers.CreateGrocery)
	router.PATCH("/grocery/:id", controllers.UpdateGrocery)
	router.DELETE("/grocery/:id", controllers.DeleteGrocery)


	log.Fatal(router.Run(":1437"))
}