package genetic

import (
	"errors"
	"fmt"
	"gitlab.com/beremaran/eavo/domain/aotree"
	"gitlab.com/beremaran/eavo/domain/entities"
	"log"
	"math/rand"
)

//Optimizer runs the genetic process, optimizing its population
//generation by generation
type Optimizer struct {
	Context          Context
	Selector         Selector
	Mutations        []Mutation
	FitnessFunctions []FitnessFunction
	Generator        GeneratorFunction
	Crossover        CrossoverFunction

	population     []*Individual
	bestIndividual *Individual
}

//NewOptimizer creates a new Optimizer instance with given configuration
func NewOptimizer(context Context, selector Selector, generator GeneratorFunction, crossover CrossoverFunction) *Optimizer {
	optimizer := &Optimizer{Context: context, Selector: selector, Generator: generator, Crossover: crossover}
	optimizer.Generator.SetContext(&optimizer.Context)

	return optimizer
}

//Optimize optimizes its Context.Problem by using genetic processes
func (o *Optimizer) Optimize() ([]*entities.Box, error) {
	var err error
	var generation int

	err = o.generatePopulation()
	if err != nil {
		return nil, err
	}

	err = o.calculatePopulationFitness()
	if err != nil {
		return nil, err
	}

	generation = 1
	o.updateBestIndividual()
	for ; generation <= o.Context.MaxGenerations; generation++ {
		err = o.step()
		if err != nil {
			return nil, err
		}

		err = o.calculatePopulationFitness()
		if err != nil {
			return nil, err
		}

		o.updateBestIndividual()
		o.logGeneration(generation)
	}

	return o.decode(o.bestIndividual), nil
}

func (o *Optimizer) generatePopulation() error {
	if o.Generator == nil {
		return fmt.Errorf("no generator found")
	}

	for len(o.population) <= o.Context.PopulationSize {
		generated, err := o.Generator.Generate()
		if err != nil {
			return err
		}

		o.population = append(o.population, generated)
	}

	return nil
}

func (o *Optimizer) calculatePopulationFitness() error {
	totalWeight := o.getTotalFitnessScoreWeight()

	for i := 0; i < len(o.population); i++ {
		o.population[i].FitnessScore = 0

		for j := 0; j < len(o.FitnessFunctions); j++ {
			fitnessScore, err := o.FitnessFunctions[j].Calculate(o.population[i])
			if err != nil {
				return err
			}

			o.population[i].FitnessScore += fitnessScore
		}

		o.population[i].FitnessScore /= totalWeight
	}

	return nil
}

func (o *Optimizer) step() error {
	var offspring []*Individual

	if o.Context.NumberOfElites > 0 {
		if o.Context.NumberOfElites > len(o.population) {
			return errors.New("number of elites is greater than population size")
		}

		offspring = append(offspring, o.getBestIndividuals(o.Context.NumberOfElites)...)
	}

	for len(offspring) < o.Context.PopulationSize {
		var newIndividual *Individual
		var err error

		if rand.Float64() <= o.Context.CrossoverRate {
			newIndividual, err = o.stepWithCrossover()
			if err != nil {
				return err
			}
		} else {
			newIndividual, err = o.stepWithoutCrossover()
			if err != nil {
				return err
			}
		}

		offspring = append(offspring, newIndividual)
	}

	o.population = offspring
	return nil
}

func (o *Optimizer) updateBestIndividual() {
	for i := 0; i < len(o.population); i++ {
		if o.bestIndividual == nil ||
			o.population[i].FitnessScore > o.bestIndividual.FitnessScore {
			o.bestIndividual = o.population[i].Copy()
		}
	}
}

func (o *Optimizer) logGeneration(generation int) {
	log.Printf("Generation #%04d: %.4f", generation, o.bestIndividual.FitnessScore)
}

func (o *Optimizer) getTotalFitnessScoreWeight() float64 {
	totalWeight := 0.0

	for i := 0; i < len(o.FitnessFunctions); i++ {
		totalWeight += o.FitnessFunctions[i].GetWeight()
	}

	return totalWeight
}

//AddMutation adds a new mutation to optimizer's mutation set
func (o *Optimizer) AddMutation(mutation Mutation) {
	o.Mutations = append(o.Mutations, mutation)
}

//AddFitnessFunction adds a new fitness function to optimizer's fitness function set
func (o *Optimizer) AddFitnessFunction(function FitnessFunction) {
	function.SetContext(&o.Context)
	o.FitnessFunctions = append(o.FitnessFunctions, function)
}

func (o *Optimizer) stepWithCrossover() (*Individual, error) {
	parent1, err := o.Selector.Select(o.population)
	if err != nil {
		return nil, err
	}

	parent2, err := o.Selector.Select(o.population)
	if err != nil {
		return nil, err
	}

	child, err := o.Crossover.Crossover(parent1, parent2)
	if err != nil {
		return nil, err
	}

	err = o.mutate(child)
	if err != nil {
		return nil, err
	}

	return child, nil
}

func (o *Optimizer) stepWithoutCrossover() (*Individual, error) {
	individual, err := o.Selector.Select(o.population)
	if err != nil {
		return nil, err
	}

	err = o.mutate(individual)
	if err != nil {
		return nil, err
	}

	return individual, nil
}

func (o *Optimizer) mutate(individual *Individual) error {
	if rand.Float64() <= o.Context.MutationRate {
		mutation := o.Mutations[rand.Intn(len(o.Mutations))]
		err := mutation.Mutate(individual)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *Optimizer) decode(individual *Individual) []*entities.Box {
	boxes, err := aotree.MapToBoxes(individual.Genome)
	if err != nil {
		panic(err)
	}

	return boxes
}

func (o *Optimizer) getBestIndividuals(i int) []*Individual {
	var bestIndividuals []*Individual

	for len(bestIndividuals) < i {
		bestIndividuals = append(bestIndividuals, o.getBestIndividual(bestIndividuals))
	}

	return bestIndividuals
}

func (o *Optimizer) getBestIndividual(except []*Individual) *Individual {
	var bestUntil *Individual

	for i := 0; i < len(o.population); i++ {
		if bestUntil == nil ||
			(o.population[i].FitnessScore > bestUntil.FitnessScore && !ContainsIndividual(except, o.population[i])) {
			bestUntil = o.population[i]
		}
	}

	return bestUntil
}
