package farm

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/Pizhlo/whatToFarm-test/server"
)

// Server makes a server with routes and specified port
func Server(port int) {
	fmt.Println("Launching the server...")

	fmt.Printf("Listening at port %d...\n", port)

	router := gin.Default()

	router.GET("/api/v1/rates", server.GetRates)
	//router.POST("/api/v1/rates", server.GetRates)

	router.Run(fmt.Sprintf(":%d", port))
}
