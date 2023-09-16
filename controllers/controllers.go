package controllers

import (
	"log"
	"net/http"

	"github.com/GetterSethya/golang-gorm/db"
	"github.com/GetterSethya/golang-gorm/models"
	"github.com/gin-gonic/gin"
)



type NewGrocery struct {
	Name string `json:"name" binding:"required"`
	Quantity int `json:"quantity" binding:"required"`
}

type GroceryUpdate struct {
	Name string `json:"name"`
	Quantity int `json:"quantity"`

}


func GetGroceries(c *gin.Context){
	var groceries []models.Grocery

	db,err :=db.Database()

	if err != nil{
		log.Println(err)
	}

	if err := db.Find(&groceries).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error":err.Error()})
		return
	}


	c.IndentedJSON(http.StatusOK, groceries)

}

func GetGrocery(c *gin.Context) {
	var grocery models.Grocery

	db,err := db.Database()

	if err != nil{
		log.Fatal(err)
	}

	if err := db.Where("id= ?", c.Param("id")).First(&grocery).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Grocery not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, grocery)

}


func CreateGrocery(c *gin.Context) {
	var grocery NewGrocery

	if err := c.ShouldBindJSON(&grocery); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	newGrocery := models.Grocery{Name: grocery.Name, Quantity: grocery.Quantity}

	db,err := db.Database()
	
	if err!=nil {
		log.Fatal(err)
	}

	if err:=db.Create(&newGrocery).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newGrocery)



}


func UpdateGrocery(c *gin.Context) {

	var grocery models.Grocery

	db, err := db.Database()

	if err!=nil {
		log.Fatal(err)
	}

	if err:=db.Where("id= ?", c.Param("id")).First(&grocery).Error; err!=nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "grocery not found"})
		return
	}

	var updateGrocery GroceryUpdate

	if err := c.ShouldBindJSON(&updateGrocery);err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err:= db.Model(&grocery).Updates(models.Grocery{Name: updateGrocery.Name, Quantity: updateGrocery.Quantity}).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":"internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, grocery)
}


func DeleteGrocery(c *gin.Context){
	var grocery models.Grocery

	db,err := db.Database()

	if err != nil {
		log.Fatal(err)
	}

	if err:=db.Where("id= ?", c.Param("id")).First(&grocery).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Grocery not found"})
		return
	}

	if err := db.Delete(&grocery).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message":"Grocery deleted successfully"})
}