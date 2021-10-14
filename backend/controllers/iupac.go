package controllers

import (
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
			"smiles": "Smiles is invalid",
		})

		return
	}

	iupac := engine.GetIUPACNomenclature(smiles)

	c.JSON(http.StatusOK, gin.H{"result" : iupac})
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
