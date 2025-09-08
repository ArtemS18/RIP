package api

import (
	"failiverCheck/internal/app/handler"
	"failiverCheck/internal/app/repository"
	"log"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	log.Println("Server start up")
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.Info(">>>> Сервер запускается <<<<")
	repo, err := repository.NewRepository()

	if err != nil {
		logrus.Error("Repo init error")
	}
	h := handler.NewHandler(repo)

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.Static("/static", "./resources")

	r.GET("/hello", h.GetComponents)
	r.GET("/component/:id", h.GetComponent)
	r.GET("/order", h.GetApplication)

	r.Run()
	log.Println("Server down")
}
