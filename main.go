package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/motestt65677/housingapp/controller"
	"github.com/orandin/lumberjackrus"
	"github.com/sirupsen/logrus"
)

func init() {

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)

	hook, err := lumberjackrus.NewHook(
		&lumberjackrus.LogFile{
			Filename: "./log/warning.log",
			MaxSize:  100,
			// MaxBackups: 1,
			// MaxAge:     1,
			Compress:  false,
			LocalTime: false,
		},
		logrus.InfoLevel,
		&logrus.JSONFormatter{},
		&lumberjackrus.LogFileOpts{
			logrus.InfoLevel: &lumberjackrus.LogFile{
				Filename: "./log/info.log",
			},
			logrus.ErrorLevel: &lumberjackrus.LogFile{
				Filename: "./log/error.log",
				MaxSize:  100, // optional
				// MaxBackups: 1,     // optional
				// MaxAge:     1,     // optional
				Compress:  false, // optional
				LocalTime: false, // optional
			},
		},
	)

	if err != nil {
		panic(err)
	}

	logrus.AddHook(hook)
}

func main() {

	// models.ConnectDatabase()

	router := gin.New()
	// router.LoadHTMLFiles("index.html")
	router.LoadHTMLGlob("views/*")

	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", controller.Home)

	router.StaticFS("/public", http.Dir("public"))
	router.Run("0.0.0.0:50080")

}
