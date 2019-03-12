package genetic

//ContainsIndividual checks if given Individual slice contains the given Individual
func ContainsIndividual(arr []*Individual, ind *Individual) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == ind {
			return true
		}
	}

	return false
}
