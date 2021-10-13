package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUPACController struct{}

func (ctrl IUPACController) GetIUPACFromSMILES(c *gin.Context) {
	smiles := c.Query("smiles")

	c.JSON(http.StatusOK, smiles)
}