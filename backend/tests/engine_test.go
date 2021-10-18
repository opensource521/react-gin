package tests

import (
	"testing"

	"connamara/hw-oteron/engine"

	"github.com/stretchr/testify/assert"
)

func TestGetIUPAC(t *testing.T) {
	testCases := map[string]string{
		"CCC(CC)CC": "3-ethylpentane",
		"CCC(CC(C)C)CCC": "4-ethyl-2-methylheptane",
		"CC(C)CCCCC": "2-methylheptane",
		"CCCC(C)(CC)CC(C)": "4-ethyl-4-methylheptane",
		"CCC(C)(CC)CC(C)C": "4-ethyl-2,4-dimethylhexane",
		"CC(C)(C)CC(C)(C)C": "2,2,4,4-tetramethylpentane",
		"CCC(CC)CC(C)CC": "3-ethyl-5-methylheptane",
		"CC(C)C(C)CCC(CCCCC)CC(C)C": "2,3-dimethyl-6-(2-methylpropyl)undecane",
		"CCCCCC(CC(C)(C)CC)CC(CC)CC": "5-(2-ethylbutyl)-3,3-dimethyldecane",
		"CCC(C(C)C)C(C)C(C)C": "3-ethyl-2,4,5-trimethylhexane",
		"CCCC(CC)C(C)C(C)C": "4-ethyl-2,3-dimethylheptane",
		"CC(C)CC(CC)CC": "4-ethyl-2-methylhexane",
		"CCCC(CC)C(C)(C)CC": "4-ethyl-3,3-dimethylheptane",
		"CCC(C(C))CC(C(C)C(C)C)CCCCC": "5-(1,2-dimethylpropyl)-3-ethyldecane",
		"CCCCCC(C(CCC)(C)C)(C(CCC)(C)C)CCC(C)CC": "6,6-bis(1,1-dimethylbutyl)-3-methylundecane",
	}

	t.Run("should convert to IUPAC correctly", func (t *testing.T) {
		for smiles := range testCases {
			assert.Equal(t, testCases[smiles], engine.GetIUPAC(smiles))
		}
	})
}