package ea

import (
	"math/rand"
)

type EAInitializer interface {
	Initialize(popsize int, nbin int, nbits []int,
		ffunc func(i Individual) float32) Population
}

type EARandomInitializer struct {
}

func (ri *EARandomInitializer) Initialize(popsize int, nbin int, nbits []int,
	ffunc func(i Individual) float32) Population {
	var pop Population

	pop.Individuals = make([]Individual, popsize)

	for i := 0; i < popsize; i++ {
		pop.Individuals[i].Gene = make([][]int, nbin)
		pop.Individuals[i].XBin = make([]float64, nbin)

		for g := range pop.Individuals[i].Gene {
			pop.Individuals[i].Gene[g] = make([]int, nbits[g])
			pop.Individuals[i].FFunc = ffunc

			for k := 0; k < nbits[g]; k++ {

				if rand.Float32() <= 0.5 {
					pop.Individuals[i].Gene[g][k] = 0
				} else {
					pop.Individuals[i].Gene[g][k] = 1
				}

			}

		}
	}

	return pop
}

func (ri *EA) InitializeEmpty(popsize int, nbin int, nbits []int,
	ffunc func(i Individual) float32) Population {
	var pop Population

	pop.Individuals = make([]Individual, popsize)

	for i := 0; i < popsize; i++ {
		pop.Individuals[i].Gene = make([][]int, nbin)
		pop.Individuals[i].XBin = make([]float64, nbin)

		for g := range pop.Individuals[i].Gene {
			pop.Individuals[i].Gene[g] = make([]int, nbits[g])
			pop.Individuals[i].FFunc = ffunc
		}
	}

	return pop
}
