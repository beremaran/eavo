package genetic

//FitnessFunction calculates an individual's fitness score
type FitnessFunction interface {
	GetWeight() float64
	SetContext(ctx *Context)
	Calculate(individual *Individual) (float64, error)
}
