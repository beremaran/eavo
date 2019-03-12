package genetic

//Mutation applies mutation operator to an Individual
type Mutation interface {
	GetProbability() float64
	Mutate(individual *Individual) error
}
