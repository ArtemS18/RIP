package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) GetSystemCalc(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error(err)
	}
	components, err := h.Repository.GetComponentsInSystemCalc(uint(id))
	if err != nil {
		log.Error(err)
	}

	ctx.HTML(http.StatusOK, "application.html", gin.H{
		"components": components,
	})
}

func (h *Handler) AddComponentInSystemCalc(ctx *gin.Context) {
	var err error
	strId := ctx.PostForm("component_id")
	componentId, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	search := ctx.PostForm("search")

	err = h.Repository.AddComponentInSystemCalc(uint(componentId), 1)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.Set("search", search)
	h.GetComponents(ctx)
}
