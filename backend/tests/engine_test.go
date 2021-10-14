package tests

import (
	"testing"

	"connamara/hw-oteron/engine"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	smiles string
	iupac string
}

var mockData = []TestCase{
	{
		"CCC(CC)CC",
		"3-ethylpentane",
	},
  {
		"CCC(CC(C)C)CCC",
		"4-ethyl-2-methyloctane",
	},
  {
		"CC(C)CCCCC",
		"2-methylheptane",
	},
  {
		"CCCC(C)(CC)CC(C)",
		"4-ethyl-4-methylheptane",
	},
  {
		"CCC(C)(CC)CC(C)C",
		"2-ethyl-2,4-methylhexane",
	},
}

func TestGetIUPACNomenclature(t *testing.T) {
	t.Run("should convert to IUPAC nomenclature correctly", func (t *testing.T) {
		for i := 0; i < 5; i ++ {
			assert.Equal(t, mockData[i].iupac, engine.GetIUPACNomenclature(mockData[i].smiles))
		}
	})
}