package controllers

import (
	"net/http"

	"go-example-vdt-postgis/infra/database"
	"go-example-vdt-postgis/models"

	"github.com/gin-gonic/gin"
)

func DomainNameChanges(ctx *gin.Context) {
	var locations []models.Locations

	regexString := `(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]`

	if err := database.DB.Raw("UPDATE my_table SET website = substring(website, ?)", regexString).Scan(&locations).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	} else {
		ctx.JSON(http.StatusOK, "Update successful")
	}

}
