package main

/*import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/russross/blackfriday"

	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
		//log.Fatal("$PORT must be set")
	}

	db := initDB()
	defer db.Close()

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

	router.GET("/twilio", SendSMS)

	router.GET("/contacts", getContacts)
	router.POST("/contact/new", newContact)
	router.PUT("/contact/:id", handleContact)
	router.DELETE("/contact/:id", handleContact)
	router.GET("/contact/:id", handleContact)

	router.Run(":" + port)
}*/
import (
	"os"

	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
		//log.Fatal("$PORT must be set")
	}
	alexa.Run(Applications, port)
}
func EchoIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	if echoReq.Request.Intent.Name == "balance" {
		echoResp.OutputSpeech("Balance request")
		return
	}
	if echoReq.Request.Intent.Name == "number" {
		echoResp.OutputSpeech("number request")
		return
	}
	if echoReq.Request.Intent.Name == "networth" {
		echoResp.OutputSpeech("net worth request")
		return
	}
	if echoReq.Request.Intent.Name == "spending" {
		for k, v := range echoReq.Request.Intent.Slots {
			echoResp.OutputSpeech("spending request key "+k+" name "+v.Name+" and value "+v.Value).Card("Hello World", "spending")
		}
		if len(echoReq.Request.Intent.Slots) > 0 {
			return
		}
		echoResp.OutputSpeech("spending request").Card("Hello World", "spending")
		return
	}
	echoResp.OutputSpeech("Why is Connor super cool!").Card("Hello World", "This is a test card.")
}

var Applications = map[string]interface{}{
	"/echo/": alexa.EchoApplication{ // Route
		AppID:    "amzn1.ask.skill.36681e9c-17d6-4051-8f12-30ba2328e619", // Echo App ID from Amazon Dashboard
		OnIntent: EchoIntentHandler,
		OnLaunch: EchoIntentHandler,
	},
}
