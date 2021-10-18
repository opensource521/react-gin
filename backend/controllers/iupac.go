package controllers

import (
	"fmt"
	"net/http"
	"regexp"

	"connamara/hw-oteron/engine"

	"github.com/gin-gonic/gin"
)

type IUPACController struct{}

var smilesRegex = regexp.MustCompile("^[C()]+")

func (ctrl IUPACController) GetIUPACFromSMILES(c *gin.Context) {
	smiles := c.Query("smiles")
	
	if !isSmilesValid(smiles) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": "SMILES is invalid",
		})

		return
	}

	result, err := engine.GetIUPAC(smiles)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"result" : result})
	} else {
		fmt.Println(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": err.Error(),
		})
	}
}

func isSmilesValid(smiles string) bool {
	if smiles == "" {
		return false
	}

	if smiles[0] != 'C' {
		return false
	}

	return smilesRegex.MatchString(smiles)
}
