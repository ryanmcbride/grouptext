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
	echoResp.OutputSpeech("Hello world from my new Echo test app!").Card("Hello World", "This is a test card.")
}

var Applications = map[string]interface{}{
	"/echo/hello": alexa.EchoApplication{ // Route
		AppID:    "amzn1.ask.skill.36681e9c-17d6-4051-8f12-30ba2328e619", // Echo App ID from Amazon Dashboard
		OnIntent: EchoIntentHandler,
		OnLaunch: EchoIntentHandler,
	},
	"/echo": alexa.EchoApplication{ // Route
		AppID:    "amzn1.ask.skill.36681e9c-17d6-4051-8f12-30ba2328e619", // Echo App ID from Amazon Dashboard
		OnIntent: EchoIntentHandler,
		OnLaunch: EchoIntentHandler,
	},
}
