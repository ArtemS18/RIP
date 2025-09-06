package handler

import (
	"failiverCheck/internal/app/repository"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repository *repository.Repository
}

type TemplateData struct {
	Components   []repository.Component
	CountQuery   int
	CurrentCount int
	SearchQuery  string
}

func NewHandler(r *repository.Repository) *Handler {
	return &Handler{r}
}

func (h *Handler) GetApplication(ctx *gin.Context) {

	components, err := h.Repository.GetComponentsInApplication()
	if err != nil {
		logrus.Error(err)
	}

	ctx.HTML(http.StatusOK, "application.html", gin.H{
		"components": components,
	})
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

	searchQuery := ctx.Query("search")
	sum := ctx.Query("addComponent")
	count := 0
	if sum != "" {
		count, err = strconv.Atoi(sum)
		if err != nil {
			count = 0
		}
	}

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
		"data": TemplateData{
			Components:   orders,
			SearchQuery:  searchQuery,
			CountQuery:   count + 1,
			CurrentCount: count},
	})
}
