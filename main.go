package main

import (
	"net/http"
	"log"
	"time"
	routes "farmme-api/route"
	config "farmme-api/config"
	"farmme-api/config/db"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)     

func main() {


	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	database, err := db.GetDBCollection()
	if err != nil {
		log.Fatalf("todo: database configuration failed: %v", err)
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	routes.Route(router, database)
	// router.GET("/swg/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// use ginSwagger middleware to serve the API docs
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Run(config.PORT_WEB_SERVICE)
}