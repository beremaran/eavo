package genetic

//CrossoverFunction instances applies crossover operator on two individuals
type CrossoverFunction interface {
	Crossover(parent1, parent2 *Individual) (*Individual, error)
}
