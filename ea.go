package ea

import (
	"fmt"
)

type EAParameter struct {
	Popsize   int
	NGen      int
	NBin      int
	NBits     []int
	MinBinVar []float64
	MaxBinVar []float64

	PCross_bin float32
	PMut_bin   float32

	// Executors EA
	Initializer EAInitializer
	Selector    EASelector
	Crosser     EACrosser
	Mutator     EAMutator
}

type EA struct {
	ParentPop Population
	ChildPop  Population
	Parameter EAParameter
}

func NewEA(parameters EAParameter) *EA {
	ea := new(EA)
	ea.Parameter = parameters
	return ea
}

func (ea *EA) Init(popsize int, ffunc func(i Individual) float32) {

	nbin := ea.Parameter.NBin
	nbits := ea.Parameter.NBits

	ea.ParentPop = ea.Parameter.Initializer.Initialize(popsize,
		nbin,
		nbits,
		ffunc)
	ea.ChildPop = ea.InitializeEmpty(popsize,
		nbin,
		nbits,
		ffunc)

	ea.Parameter.Popsize = popsize
	ea.DecodePop()
}

func (ea *EA) Run() {
	for i := 0; i < ea.Parameter.NGen; i++ {

		/* Selection and Crossover */
		for p := 0; p < len(ea.ParentPop.Individuals); p += 2 {
			ea.ChildPop.Individuals[p], ea.ChildPop.Individuals[p+1] = ea.Parameter.Crosser.Crossover(
				ea.Parameter.Selector.Select(ea.ParentPop),
				ea.Parameter.Selector.Select(ea.ParentPop),
				ea.Parameter,
			)
		}

		/* Mutation */
		ea.Parameter.Mutator.Mutate(ea.ChildPop, ea.Parameter)
		ea.DecodeChildPop()
		ea.EvalPop(ea.ChildPop)

		ea.ParentPop = ea.ChildPop
		ea.ChildPop = ea.InitializeEmpty(ea.Parameter.Popsize,
			ea.Parameter.NBin,
			ea.Parameter.NBits,
			ea.ParentPop.Individuals[0].FFunc)

	}
}

func (ea *EA) Report(p Population) {
	for _, i := range p.Individuals {
		fmt.Printf("%s", "Gene ")
		for g := 0; g < ea.Parameter.NBin; g++ {
			fmt.Printf("%v", i.Gene[g])
		}
		fmt.Printf("%s %v", "\tXBin ", i.XBin)
		fmt.Printf("%s %f", "\tFitness ", i.Fitness)
		fmt.Printf("\n")
	}

}
