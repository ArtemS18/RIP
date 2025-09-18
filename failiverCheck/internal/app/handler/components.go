package handler

import (
	"failiverCheck/internal/app/ds"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type TemplateData struct {
	Components   []ds.Component
	CountQuery   int
	CurrentCount int
	SearchQuery  string
	SystemCalcId int
}

func (h *Handler) GetComponent(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error(err)
	}

	order, err := h.Repository.GetComponentById(id)
	if err != nil {
		log.Error(err)
	}

	ctx.HTML(http.StatusOK, "component.html", gin.H{
		"component": order,
	})
}

func (h *Handler) GetComponents(ctx *gin.Context) {
	var orders []ds.Component
	var err error
	var searchQuery string
	param, ok := ctx.Get("search")
	if !ok {
		searchQuery = ctx.Query("search")
	} else {
		searchQuery, ok = param.(string)
		if !ok {
			log.Error(err)
		}
	}
	count, err := h.Repository.GetCountComnponents(1)
	if err != nil {
		log.Error(err)
	}

	log.Info(searchQuery)
	if searchQuery == "" {
		orders, err = h.Repository.GetComponents()
		if err != nil {
			log.Error(err)
		}
	} else {
		orders, err = h.Repository.GetComponentsByTitle(searchQuery)
		if err != nil {
			log.Error(err)
		}
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"data": TemplateData{
			Components:   orders,
			SearchQuery:  searchQuery,
			SystemCalcId: 1,
			CurrentCount: count},
	})
}
