package main

import (
	"fmt"
	"github.com/martinre2/go-ea"
	"math"
	"math/rand"
	"time"
)

func fitness(i ea.Individual) float32 {
	i.Fitness = float32(math.Pow(i.XBin[0], 2))
	return i.Fitness
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	nbits := []int{8}
	minbinvar := []float64{0.0}
	maxbinvar := []float64{5.5}

	param := ea.EAParameter{
		NGen:        2,
		NBin:        1,
		NBits:       nbits,
		MinBinVar:   minbinvar,
		MaxBinVar:   maxbinvar,
		PCross_bin:  0.4,
		PMut_bin:    0.2,
		Initializer: new(ea.EARandomInitializer),
		Selector:    new(ea.RouleteSelector),
		Crosser:     new(ea.EA2PCrossover),
		Mutator:     new(ea.BMutator),
	}

	eai := ea.NewEA(param)
	eai.Init(20, fitness)
	eai.EvalPop(eai.ParentPop)

	eai.Run()

	eai.Report(eai.ParentPop)
}
