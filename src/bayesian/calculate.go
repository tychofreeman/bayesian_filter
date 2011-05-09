package bayesian

import . "big"

type BayesRule struct {
	probBWhenA *Rat
	probA *Rat
	probB *Rat
}

func (br *BayesRule) Calculate() *Rat {
	return NewRat(1, 1).Quo(NewRat(1, 1).Mul(br.probBWhenA, br.probA), br.probB)
}
