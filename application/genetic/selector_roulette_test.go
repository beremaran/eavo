package genetic

import (
	"gitlab.com/beremaran/eavo/domain/genetic"
	"testing"
)

func TestRouletteWheelSelector_Select(t *testing.T) {
	selector := RouletteWheelSelector{}
	population := []*genetic.Individual{{}, {}, {}, {}, {}, {}}

	individual, err := selector.Select(population)
	if err != nil {
		t.Error(err)
	}

	if individual == nil {
		t.Errorf("nil individual selected")
	}

	population = nil
	individual, err = selector.Select(population)
	if err == nil {
		t.Errorf("empty population does not caused error")
	}
}
