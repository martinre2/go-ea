package ea

import (
	"math/rand"
)

type EAMutator interface {
	Mutate(p Population, pm EAParameter)
}

type BMutator struct {
	Parameters EAParameter
}

func (bm BMutator) Mutate(p Population, params EAParameter) {
	bm.Parameters = params
	for _, i := range p.Individuals {
		bm.MutateInd(i)
	}
}

func (bm BMutator) MutateInd(i Individual) {
	for j := 0; j < bm.Parameters.NBin; j++ {
		for k := 0; k < bm.Parameters.NBits[j]; k++ {

			prob := rand.Float32()
			if prob <= bm.Parameters.PMut_bin {
				if i.Gene[j][k] == 0 {
					i.Gene[j][k] = 1
				} else {
					i.Gene[j][k] = 0
				}
			}
		}
	}
}
