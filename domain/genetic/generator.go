package genetic

//GeneratorFunction generates an individual for initial population
type GeneratorFunction interface {
	SetContext(ctx *Context)
	Generate() (*Individual, error)
}
