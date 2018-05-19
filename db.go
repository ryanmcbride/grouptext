package main

import (
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var globalDB *gorm.DB

func initDB() *gorm.DB {
	dbSpec := os.Getenv("DATABASE_URL")

	if dbSpec == "" {
		dbSpec = "host=127.0.0.1 port=5432 user=ryanmcbride dbname=postgres sslmode=disable"
	}

	db, err := gorm.Open("postgres", dbSpec)
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	// Migrate the schema
	db.AutoMigrate(&Contact{})

	globalDB = db
	return db
}

func getAll(m interface{}, c *gin.Context) {
	globalDB.Model(m).Find(m)
	c.JSON(http.StatusOK, m)
}

func handleCRUDForModel(i interface{}, c *gin.Context) {
	switch c.Request.Method {
	case "POST":
		createModel(i, c)
	case "GET":
		readModel(i, c)
	case "PUT":
		updateModel(i, c)
	case "DELETE":
		deleteModel(i, c)
	}
}
func createModel(m interface{}, c *gin.Context) {
	c.BindJSON(&m)
	globalDB.Create(m)
	c.JSON(http.StatusOK, m)
}
func readModel(m interface{}, c *gin.Context) {
	id := c.Params.ByName("id")
	globalDB.First(m, id)
	c.JSON(http.StatusOK, m)
}
func updateModel(m interface{}, c *gin.Context) {
	id := c.Params.ByName("id")
	globalDB.First(m, id)
	updated := reflect.New(reflect.TypeOf(m))
	c.BindJSON(updated)
	globalDB.Model(m).Update(updated)
	c.JSON(http.StatusOK, m)
}
func deleteModel(m interface{}, c *gin.Context) {
	id := c.Params.ByName("id")
	globalDB.First(m, id)
	globalDB.Delete(m)
	c.JSON(http.StatusOK, m)
}
