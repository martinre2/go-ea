package ea

import (
	"math/rand"
)

type EACrosser interface {
	Crossover(a, b Individual, params EAParameter) (Individual, Individual)
}

/* two point binary crossover*/
type EA2PCrossover struct{}

func (bx *EA2PCrossover) Crossover(a, b Individual, params EAParameter) (Individual, Individual) {
	/* Initialize childs*/

	ca := new(Individual)

	ca.Gene = make([][]int, params.NBin)
	ca.XBin = make([]float64, params.NBin)

	for g := range ca.Gene {
		ca.Gene[g] = make([]int, params.NBits[g])
		ca.FFunc = a.FFunc
	}
	cb := ca.Copy()

	for i := 0; i < params.NBin; i++ {
		rn := rand.Float32()

		if rn <= params.PCross_bin {

			site1 := rand.Intn(params.NBits[i] - 1)
			site2 := rand.Intn(params.NBits[i] - 1)

			if site1 > site2 {
				temp := site1
				site1 = site2
				site2 = temp
			}

			for j := 0; j < site1; j++ {
				ca.Gene[i][j] = a.Gene[i][j]
				cb.Gene[i][j] = b.Gene[i][j]
			}
			for j := site1; j < site2; j++ {
				ca.Gene[i][j] = b.Gene[i][j]
				cb.Gene[i][j] = a.Gene[i][j]
			}
			for j := site2; j < params.NBits[i]; j++ {
				ca.Gene[i][j] = a.Gene[i][j]
				cb.Gene[i][j] = b.Gene[i][j]
			}
		} else {
			for j := 0; j < params.NBits[i]; j++ {
				ca.Gene[i][j] = a.Gene[i][j]
				ca.Gene[i][j] = b.Gene[i][j]
			}
		}
	}

	return *ca, *cb
}
