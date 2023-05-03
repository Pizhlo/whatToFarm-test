package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Start makes a server with routes and specified port
func Start(port int) {
	fmt.Println("Launching the server...")

	fmt.Printf("Listening at port %d...\n", port)

	router := gin.Default()

	router.GET("/api/v1/rates", GetRates)
	//router.POST("/api/v1/rates", server.GetRates)

	router.Run(fmt.Sprintf(":%d", port))
}
