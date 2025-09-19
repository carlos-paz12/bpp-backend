package controllers

import (
	"spe/models"
	"spe/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

type PointRecordController struct{}

func (PointRecordController) GetMyPoints(c *gin.Context) {
	uid := c.GetInt64("user_id")
	records, err := services.PointRecordService{}.GetMyPoints(uid)

	if err != nil {
		c.JSON(http.StatusForbidden, models.ApiResponse{
			Message:  "",
			Error:    err.Error(),
			HttpCode: http.StatusForbidden,
		})
		return
	}

	c.JSON(http.StatusOK, models.ApiResponse{
		Data:     records,
		Message:  "it's ok.",
		Error:    "",
		HttpCode: http.StatusOK,
	})
}

func (PointRecordController) RegisterMyPoint(c *gin.Context) {
	uid := c.GetInt64("user_id")
	record, err := services.PointRecordService{}.RegisterMyPoint(uid)

	if err != nil {
		c.JSON(http.StatusForbidden, models.ApiResponse{
			Message:  "",
			Error:    err.Error(),
			HttpCode: http.StatusForbidden,
		})
		return
	}

	c.JSON(http.StatusOK, models.ApiResponse{
		Data:     record,
		Message:  "it's ok.",
		Error:    "",
		HttpCode: http.StatusOK,
	})
}

func (PointRecordController) GetMyLastPoint(c *gin.Context) {
	uid := c.GetInt64("user_id")
	last, err := services.PointRecordService{}.GetMyLastPoint(uid)

	if err != nil {
		c.JSON(http.StatusForbidden, models.ApiResponse{
			Message:  "",
			Error:    err.Error(),
			HttpCode: http.StatusForbidden,
		})
		return
	}

	c.JSON(http.StatusOK, models.ApiResponse{
		Data:     last,
		Message:  "it's ok.",
		Error:    "",
		HttpCode: http.StatusOK,
	})
}

func (PointRecordController) GetPointsByScholarshipID(c *gin.Context) {
	// Todo
}

func (PointRecordController) GetLastPointByScholarshipID(c *gin.Context) {
	// Todo
}
