package tests

import (
	"errors"
	"testing"

	"connamara/hw-oteron/engine"

	"github.com/stretchr/testify/assert"
)

type TestResult struct {
	iupac string
	err error
}

func TestGetIUPAC(t *testing.T) {
	testCases := map[string]TestResult{
		"CCC(CC)CC": {"3-ethylpentane", nil},
		"CCC(CC(C)C)CCC": {"4-ethyl-2-methylheptane", nil},
		"CC(C)CCCCC": {"2-methylheptane", nil},
		"CCCC(C)(CC)CC(C)": {"4-ethyl-4-methylheptane", nil},
		"CCC(C)(CC)CC(C)C": {"4-ethyl-2,4-dimethylhexane", nil},
		"CC(C)(C)CC(C)(C)C": {"2,2,4,4-tetramethylpentane", nil},
		"CCC(CC)CC(C)CC": {"3-ethyl-5-methylheptane", nil},
		"CC(C)C(C)CCC(CCCCC)CC(C)C": {"2,3-dimethyl-6-(2-methylpropyl)undecane", nil},
		"CCCCCC(CC(C)(C)CC)CC(CC)CC": {"5-(2-ethylbutyl)-3,3-dimethyldecane", nil},
		"CCC(C(C)C)C(C)C(C)C": {"3-ethyl-2,4,5-trimethylhexane", nil},
		"CCCC(CC)C(C)C(C)C": {"4-ethyl-2,3-dimethylheptane", nil},
		"CC(C)CC(CC)CC": {"4-ethyl-2-methylhexane", nil},
		"CCCC(CC)C(C)(C)CC": {"4-ethyl-3,3-dimethylheptane", nil},
		"CCC(C(C))CC(C(C)C(C)C)CCCCC": {"5-(1,2-dimethylpropyl)-3-ethyldecane", nil},
		"CCCCCC(C(CCC)(C)C)(C(CCC)(C)C)CCC(C)CC": {"6,6-bis(1,1-dimethylbutyl)-3-methylundecane", nil},
		"CCC((C)CCCC": {"", errors.New("Brackets not closed properly")},
		"CC))((CC": {"", errors.New("Brackets not opened properly")},
		"CCC()CCC": {"", errors.New("Empty content inside brackets")},
	}

	t.Run("should convert to IUPAC correctly", func (t *testing.T) {
		for smiles := range testCases {
			iupac, err := engine.GetIUPAC(smiles)
			assert.Equal(t, testCases[smiles].iupac, iupac)
			assert.Equal(t, testCases[smiles].err, err)
		}
	})
}