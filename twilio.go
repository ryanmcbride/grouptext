package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"github.com/sfreiberg/gotwilio"
)

func SendSMS(c *gin.Context) {
	accountSid := "AC875040d5a71b14bbe993c6ec87f37fac"
	authToken := "add66f06a8da7d9603786c7b449613b1"
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := "13852557576"
	to := "18016733529"
	message := "Welcome to gotwilio!"
	response, exception, err := twilio.SendSMS(from, to, message, "", "")
	if err != nil {
		log.Print(err)
		log.Print(response)
		log.Print(exception)
	}
	c.String(http.StatusOK, string(blackfriday.MarkdownBasic([]byte("**Message Sent**"))))
}
