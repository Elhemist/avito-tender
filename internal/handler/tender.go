package handler

import (
	"avito-tender/internal/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetTenders(c *gin.Context) {
	logrus.Info("Try to get all tenders list")
	tenderList, err := h.services.Tender.GetAllTenders()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tenderList)
}

func (h *Handler) GetUserTenders(c *gin.Context) {
	username := c.Query("username")
	logrus.Info("Try to get tenders from user:", username)
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
		return
	}

	tenderList, err := h.services.Tender.GetUserTenders(username)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tenderList)
}

func (h *Handler) CreateTender(c *gin.Context) {
	logrus.Info("Try to create tender")
	var newTender models.Tender
	if err := c.ShouldBindJSON(&newTender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logrus.Info(newTender, "debug")

	tender, err := h.services.Tender.CreateTender(newTender)
	logrus.Info("создали ", "debug")
	if err != nil {
		if err == fmt.Errorf("access denied") {
			newErrorResponse(c, http.StatusForbidden, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tender)
}

func (h *Handler) EditTender(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	logrus.Info("id ", id, "debug")
	var update models.UpdateTenderRequest
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Info("up ", update, "debug")
	updatedTender, err := h.services.Tender.EditTender(id, update)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, updatedTender)
}

func (h *Handler) RollbackTender(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	versionStr := c.Param("version")
	version, err := strconv.Atoi(versionStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid version format"})
		return
	}

	updatedTender, err := h.services.Tender.RollbackTender(id, version)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, updatedTender)
}
