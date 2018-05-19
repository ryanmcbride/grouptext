package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Filters string `json:"filters" binding:"required"`
}

func getContacts(c *gin.Context) {
	getAll(&[]Contact{}, c)
}
func handleContact(c *gin.Context) {
	//handleCRUDForModel(&Contact{}, c)
	switch c.Request.Method {
	case "POST":
	case "GET":
	case "PUT":
		id := c.Params.ByName("id")
		contact := Contact{}
		globalDB.First(&contact, id)
		updated := Contact{}
		err := c.BindJSON(&updated)
		if err != nil {
			log.Print(err)
		}
		globalDB.Model(&contact).Update(updated)
		c.JSON(http.StatusOK, &contact)
	case "DELETE":
		id := c.Params.ByName("id")
		contact := Contact{}
		globalDB.First(&contact, id)
		globalDB.Delete(&contact)
	}
}
func newContact(c *gin.Context) {
	contact := Contact{}
	c.BindJSON(&contact)
	globalDB.Create(&contact)
	c.JSON(http.StatusOK, contact)
}
