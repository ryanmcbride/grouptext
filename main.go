package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/russross/blackfriday"
)

func main() {
	dbSpec := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	if dbSpec == "" {
		dbSpec = "host=127.0.0.1 port=5432 user=ryanmcbride dbname=postgres sslmode=disable"
	}

	db, err := gorm.Open("postgres", dbSpec)
	defer db.Close()
	if err != nil {
		log.Fatal("Can't connect to the DB")
	}

	if port == "" {
		port = "5000"
		//log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/mark", func(c *gin.Context) {
		c.String(http.StatusOK, string(blackfriday.MarkdownBasic([]byte("**hi!**"))))
	})

	router.Run(":" + port)
}
