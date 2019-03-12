package genetic

import "testing"

func TestContainsIndividual(t *testing.T) {
	toSearch := &Individual{}
	individuals := []*Individual{{}, toSearch, {}}

	doesContain := ContainsIndividual(individuals, toSearch)
	if !doesContain {
		t.Errorf("individual not found")
	}

	doesContain = ContainsIndividual([]*Individual{}, toSearch)
	if doesContain {
		t.Errorf("individual found in empty slice")
	}
}
