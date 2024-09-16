package handler

import (
	"avito-tender/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "ok")
		})

		tenders := api.Group("/tenders")
		{
			tenders.GET("/", h.GetTenders)
			tenders.GET("/my", h.GetUserTenders)
			tenders.POST("/new", h.CreateTender)
			tenders.PATCH("/:id/edit", h.EditTender)
			tenders.PUT("/:id/rollback/:version", h.RollbackTender)
		}

		bids := api.Group("/bids")
		{
			bids.GET("/:tenderId/list", h.GetTenderBids)
			bids.GET("/my", h.GetUserBids)
			bids.POST("/new", h.CreateBid)
			bids.PATCH("/:id/edit", h.EditBids)
			bids.PUT("/:id/rollback/:version", h.RollbackBids)
		}
	}

	return router
}
