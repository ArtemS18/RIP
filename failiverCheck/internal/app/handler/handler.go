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
	OrderId      int
}

func NewHandler(r *repository.Repository) *Handler {
	return &Handler{r}
}

func (h *Handler) GetApplication(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		logrus.Error(err)
	}
	calculation, err := h.Repository.GetComponentsInApplication(id)
	componentsInCalc := calculation.Components
	if err != nil {
		logrus.Error(err)
	}

	ctx.HTML(http.StatusOK, "application.html", gin.H{
		"components": componentsInCalc,
		"name":       calculation.Name,
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

	searchQuery := ctx.Query("searchComponents")
	orderId := 1
	calculation, _ := h.Repository.GetComponentsInApplication(1)
	components := calculation.Components
	count := len(components)

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
			CurrentCount: count,
			OrderId:      orderId,
		},
	})
}
