package ea

import ()

type Individual struct {
	Gene    [][]int
	XBin    []float64
	Fitness float32
	FFunc   func(i Individual) float32
	PSelect float32
}

type Population struct {
	Individuals []Individual
}

func (i *Individual) Copy() *Individual {
	ni := new(Individual)
	ni.XBin = make([]float64, len(i.XBin))
	ni.FFunc = i.FFunc

	ni.Gene = make([][]int, len(i.Gene))
	for g := 0; g < len(i.Gene); g++ {
		ni.Gene[g] = make([]int, len(i.Gene[g]))
	}

	return ni
}
