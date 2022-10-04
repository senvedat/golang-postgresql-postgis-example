package main

import (
	"go-example-vdt-postgis/config"
	"go-example-vdt-postgis/infra/database"
	"go-example-vdt-postgis/infra/logger"
	"go-example-vdt-postgis/routers"
	"time"

	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("SERVER_TIMEZONE", "Europe/London")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}
	masterDSN := config.DbConfiguration()

	if err := database.DBConnection(masterDSN); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}

	router := routers.Routes()

	logger.Fatalf("%v", router.Run(config.ServerConfig()))

}
