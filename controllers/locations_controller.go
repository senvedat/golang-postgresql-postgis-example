package controllers

import (
	"net/http"

	"go-example-vdt-postgis/infra/database"
	models "go-example-vdt-postgis/models"

	"github.com/gin-gonic/gin"
)

// getLocations
func GetLocations(ctx *gin.Context) {

	var locations []models.Locations
	var locationReq models.GetLocationReq

	if err := ctx.ShouldBindJSON(&locationReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if locationReq.Type == 0 {
		if err := database.DB.Raw("SELECT * FROM my_table WHERE ST_DWithin(coordinates::geography, ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography, ?) ORDER BY CASE WHEN ST_Distance(coordinates::geography,ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography) < 50 THEN rating END", locationReq.Longitude, locationReq.Latitude, locationReq.Radius, locationReq.Longitude, locationReq.Latitude).Scan(&locations).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		} else {
			ctx.JSON(http.StatusOK, &locations)
		}
	} else if locationReq.Type == 1 {
		var radiusSquare float64 = float64(locationReq.Radius) / 100000
		if err := database.DB.Raw("SELECT * FROM my_table WHERE ST_Intersects(coordinates::geography, ST_Envelope(ST_Buffer(ST_SetSRID(ST_MakePoint(?, ?), 4326), ?))::geography) ORDER BY CASE WHEN ST_Distance(coordinates::geography,ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography) < 50 THEN rating END", locationReq.Longitude, locationReq.Latitude, radiusSquare, locationReq.Longitude, locationReq.Latitude).Scan(&locations).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		} else {
			ctx.JSON(http.StatusOK, &locations)
		}
	} else {
		ctx.JSON(http.StatusBadRequest, "Type error!")
	}

}
