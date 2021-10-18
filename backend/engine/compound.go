package engine

type Compound struct {
	iupac        string     // IUPAC of compound
	substituents []Compound // Substituents (side chains) of compound
}

func (a *Compound) GetIUPAC() string {
	return (*a).iupac
}

func (a *Compound) SetIUPAC(iupac string) {
	(*a).iupac = iupac
}

/*
 * Check if the compound has substituents
 */
func (a *Compound) IsComplex() bool {
	return len((*a).substituents) > 0
}

func (a *Compound) GetSubstituents() []Compound {
	return (*a).substituents
}

func (a *Compound) AppendSubstituent(s Compound) {
	(*a).substituents = append((*a).substituents, s)
}

func (a *Compound) NumberOfSubstituents() int {
	return len((*a).substituents)
}

/*
 * Return the number of complex substituents
 * Complex substituent means the substituent compound that has its substituents too
 */
func (a *Compound) NumberOfComplexSubstituents() int {
	count := 0

	for _, substituent := range (*a).substituents {
		if substituent.IsComplex() {
			count++
		}
	}

	return count
}