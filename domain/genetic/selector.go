package genetic

//Selector selects an Individual from given population according to its own rules
type Selector interface {
	Select(population []*Individual) (*Individual, error)
}
