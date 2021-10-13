package main

import (
	"github.com/gin-gonic/gin"
	"github.com/proficientwizard/hw-oteron/controllers"
)

func main() {
	router := gin.Default()

	apiGroup := router.Group("/api")
	{
		iupac := new(controllers.IUPACController)
			
		apiGroup.GET("/iupac", iupac.GetIUPACFromSMILES)
	}

	router.Run("localhost:8080")
}