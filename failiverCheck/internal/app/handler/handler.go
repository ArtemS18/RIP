package handler

import (
	"failiverCheck/internal/app/ds"
	"failiverCheck/internal/app/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	Repository *repository.Repository
}

type TemplateData struct {
	Components   []ds.Component
	CountQuery   int
	CurrentCount int
	SearchQuery  string
}

func NewHandler(r *repository.Repository) *Handler {
	return &Handler{r}
}

func (h *Handler) RegisterHandlers(r *gin.Engine) {
	r.GET("/components", h.GetComponents)
	r.GET("/components/:id", h.GetComponent)
	r.GET("/availability_calc", h.GetApplication)
}

func (h *Handler) RegisterStatic(r *gin.Engine, path string) {
	r.LoadHTMLGlob(fmt.Sprintf("%s/*", path))
	r.Static("/static", "./resources")
}

func (h *Handler) errorHandler(ctx *gin.Context, errorCode int, err error) {
	log.Error(err.Error())
	ctx.JSON(errorCode, gin.H{
		"status":      "error",
		"description": err.Error(),
	})

}
func (h *Handler) GetApplication(ctx *gin.Context) {

	components, err := h.Repository.GetComponents()
	if err != nil {
		log.Error(err)
	}

	ctx.HTML(http.StatusOK, "application.html", gin.H{
		"components": components,
	})
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

	searchQuery := ctx.Query("search")
	sum := ctx.Query("addComponent")
	count := 0
	if sum != "" {
		count, err = strconv.Atoi(sum)
		if err != nil {
			count = 0
		}
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
			CountQuery:   count + 1,
			CurrentCount: count},
	})
}
