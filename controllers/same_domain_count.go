package controllers

import (
	"net/http"

	"go-example-vdt-postgis/infra/database"
	models "go-example-vdt-postgis/models"

	"github.com/gin-gonic/gin"
)

func SameDomainCount(ctx *gin.Context) {
	var locationsCount []models.LocationsCount

	if err := database.DB.Raw("SELECT name, website, COUNT (website) AS count FROM my_table GROUP BY name, website HAVING COUNT(website)> 1").Scan(&locationsCount).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	} else {
		ctx.JSON(http.StatusOK, &locationsCount)
	}

}
