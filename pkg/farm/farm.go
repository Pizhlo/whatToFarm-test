package farm

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Server(port int) {
	fmt.Println("launching the server...")

	fmt.Printf("Listening at port %d...\n", port)

	router := gin.Default()

	router.GET("/api/v1/rates/")

	router.Run(fmt.Sprintf(":%d", port))
}
