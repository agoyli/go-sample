package main

import (
	"log"
	"main/config"
	"main/internal/handlers"
	"main/internal/store"
	"main/internal/store/pgx"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	defer store.Connect().(*pgx.Access).Close()

	runHttp(config.App.ProdEnabled, config.App.HttpPort)
}

func runHttp(isProd bool, port int) error {
	if isProd {
		gin.SetMode(gin.ReleaseMode)
	}
	routes := gin.Default()

	// middleware and routes
	handlers.Routes(routes)

	// run http
	if err := routes.Run(":" + strconv.Itoa(port)); err != nil {
		log.Fatal(err)
	}
	return nil
}
