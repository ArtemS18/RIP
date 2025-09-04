package handler

import (
	"failiverCheck/internal/app/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repository *repository.Repository
}

func NewHandler(r *repository.Repository) *Handler {
	return &Handler{r}
}

func (h *Handler) GetComponent(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.Error(err)
	}

	order, err := h.Repository.GetComponent(id)
	if err != nil {
		logrus.Error(err)
	}

	ctx.HTML(http.StatusOK, "component.html", gin.H{
		"component": order,
	})
}

func (h *Handler) GetComponents(ctx *gin.Context) {
	var orders []repository.Component
	var err error

	searchQuery := ctx.Query("query")
	logrus.Info(searchQuery)
	if searchQuery == "" {
		orders, err = h.Repository.GetComponents()
		if err != nil {
			logrus.Error(err)
		}
	} else {
		orders, err = h.Repository.GetComponentsByTitle(searchQuery)
		if err != nil {
			logrus.Error(err)
		}
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"time":       time.Now().Format("15:04:05"),
		"components": orders,
		"query":      searchQuery,
	})
}
