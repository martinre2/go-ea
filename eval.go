package ea

import ()

func (ea *EA) EvalPop(p Population) {
	for a, i := range p.Individuals {
		p.Individuals[a].Fitness = i.FFunc(i)
	}
}
