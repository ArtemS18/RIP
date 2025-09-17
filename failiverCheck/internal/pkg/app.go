package pkg

import (
	"failiverCheck/internal/app/config"
	"failiverCheck/internal/app/handler"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Application struct {
	Config  *config.Config
	Router  *gin.Engine
	Handler *handler.Handler
}

func NewApplication(c *config.Config, r *gin.Engine, h *handler.Handler) *Application {
	return &Application{
		Config:  c,
		Router:  r,
		Handler: h,
	}
}

func (app *Application) RunApplication() {
	logrus.Println("Server start up")
	app.Handler.RegisterHandlers(app.Router)
	app.Handler.RegisterStatic(app.Router, "templates")
	address := fmt.Sprintf("%s:%d", app.Config.ServiceHost, app.Config.ServicePort)
	if err := app.Router.Run(address); err != nil {
		logrus.Fatal(err)
	}
	logrus.Println("Server down")

}
