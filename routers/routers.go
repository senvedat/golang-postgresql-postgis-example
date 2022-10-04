package routers

import (
	"go-example-vdt-postgis/controllers"

	"github.com/gin-gonic/gin"
)

func MainRoutes(route *gin.Engine) {
	v1 := route.Group("/locations")
	v1.POST("get-nearby-locations", controllers.GetLocations)
	v1.GET("domain-name-change", controllers.DomainNameChanges)
	v1.GET("same-domain-count", controllers.SameDomainCount)

}
