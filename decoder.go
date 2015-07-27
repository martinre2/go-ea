package ea

import (
	"math"
)

func (ea *EA) DecodePop() {
	for _, p := range ea.ParentPop.Individuals {
		decodeInd(p, ea.Parameter)
	}
}

func (ea *EA) DecodeChildPop() {
	for _, p := range ea.ChildPop.Individuals {
		decodeInd(p, ea.Parameter)
	}
}

func decodeInd(ind Individual, parameter EAParameter) {
	var sum float64

	for j := 0; j < parameter.NBin; j++ {
		sum = 0

		for k := 0; k < parameter.NBits[j]; k++ {
			if ind.Gene[j][k] == 1 {
				sum += math.Pow(2, float64(parameter.NBits[j]-1-k))
			}
		}
		ind.XBin[j] = parameter.MinBinVar[j] + sum*(parameter.MaxBinVar[j]-parameter.MinBinVar[j])/(math.Pow(2, float64(parameter.NBits[j]))-1)
	}
}
