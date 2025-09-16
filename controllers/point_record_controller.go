package controllers

import (
	"spe/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

type PointRecordController struct{}

func (PointRecordController) Create(c *gin.Context) {
	scholarshipID := c.GetInt64("user_id")
	newPointRecord, err := services.PointRecordService{}.Create(scholarshipID)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":  "Erro ao registrar ponto.",
			"status": http.StatusForbidden,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"novo_ponto": newPointRecord,
		"error":      "",
		"status":     http.StatusOK,
	})
}

func (PointRecordController) FindAllByScholarshipID(c *gin.Context) {
	scholarshipID := c.GetInt64("user_id")
	pointRecords, err := services.PointRecordService{}.FindAllByScholarshipID(scholarshipID)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":  "Erro ao recuperar pontos.",
			"status": http.StatusForbidden,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"pontos": pointRecords,
		"error":  "",
		"status": http.StatusOK,
	})
}

func (PointRecordController) FindLastByScholarshipID(c *gin.Context) {
	scholarshipID := c.GetInt64("user_id")
	lastRecord, err := services.PointRecordService{}.FindLastByScholarshipID(scholarshipID)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":  "Erro ao recuperar pontos.",
			"status": http.StatusForbidden,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"ultimo": lastRecord,
		"error":  "",
		"status": http.StatusOK,
	})
}
