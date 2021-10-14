package main

import (
	"fmt"

	"connamara/hw-oteron/controllers"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}
	}
}

func main() {
	router := gin.Default()

	router.Use(CORSMiddleware())
	
	apiGroup := router.Group("/api")
	{
		iupac := new(controllers.IUPACController)
			
		apiGroup.GET("/iupac", iupac.GetIUPACFromSMILES)
	}

	router.Run("localhost:8080")
}