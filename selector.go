package ea

import (
	"math/rand"
)

type EASelector interface {
	Select(p Population) Individual
}

type RouleteSelector struct {
	TotalFitness float32
}

func (rs *RouleteSelector) Select(p Population) Individual {
	if rs.TotalFitness == 0 {
		rs.TotalFitness = 0
		for _, i := range p.Individuals {
			rs.TotalFitness += i.Fitness
		}
	}

	for a, i := range p.Individuals {
		if i.PSelect == 0 {
			p.Individuals[a].PSelect = i.Fitness / rs.TotalFitness
		}
	}

	var r = rand.Float32()

	for i, _ := range p.Individuals {
		if r <= p.Individuals[i].PSelect {
			return p.Individuals[i]
		}
	}

	for i, _ := range p.Individuals {
		if p.Individuals[i].PSelect > p.Individuals[i+1].PSelect {
			return p.Individuals[i]
		}
	}

	return p.Individuals[0]
}
