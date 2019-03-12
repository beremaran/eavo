package genetic

import "testing"

func TestIndividual_Copy(t *testing.T) {
	individual := &Individual{}
	individual2 := individual.Copy()

	if individual2 == nil {
		t.Errorf("copy returned nil")
	}

	if individual == individual2 {
		t.Errorf("copy returned same instance")
	}
}
