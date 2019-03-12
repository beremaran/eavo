package genetic

import (
	"gitlab.com/beremaran/eavo/domain/genetic"
	"testing"
)

func TestTournamentSelector_Select(t *testing.T) {
	selector := TournamentSelector{
		TournamentSize: 2,
	}

	population := []*genetic.Individual{{}, {}, {}, {}, {}, {}}

	individual, err := selector.Select(population)
	if err != nil {
		t.Error(err)
	}

	if individual == nil {
		t.Errorf("nil individual selected")
	}
}
