package main

import (
	"github.com/sfreiberg/gotwilio"
)

func SendSMS() {
	accountSid := "AC875040d5a71b14bbe993c6ec87f37fac"
	authToken := "add66f06a8da7d9603786c7b449613b1"
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := "13854742956"
	to := "18016733529"
	message := "Welcome to gotwilio!"
	twilio.SendSMS(from, to, message, "", "")
}
