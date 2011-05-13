package bayesian

import "fmt"
import s "strings"
import u "unicode"
import "big"

type wordStats struct {
	countB int64
	countAWithB int64
	probAWhenB *big.Rat
}

type Bayesian struct {
	totalCount int64
	countA int64
	words map[string]*wordStats
}

func lowerTrim(word string) string {
	return s.ToLower(s.TrimFunc(word, func(rune int) bool { return !u.IsLetter(rune) }))
}

func shouldIgnore(word string) bool {
	return map[string]bool {
		"to": true,
		"as": true,
		"from": true,
		"and": true,
		"or": true,
		"in": true,
		"out": true,
		"you": true,
		"your": true,
		"it's": true,
		"on": true,
		"upon": true,
		"the": true,
	}[word]
}

func (b *Bayesian) Add(data Data) {
	words := data.GetWords()
	b.totalCount++
	if data.Is() {
		b.countA++
	}
	if b.words == nil {
		b.words = map[string]*wordStats{}
	}
	uniqueWords := make(map[string]bool)
	for _, rawWord := range words {
		w := lowerTrim(rawWord)
		if shouldIgnore(w) {
			continue
		}

		d, exists := b.words[w]
		if exists == false {
			d = new(wordStats)
			b.words[w] = d

		}
		if _, defd := uniqueWords[w]; defd == false {
			d.countB++
			if  data.Is() {
				d.countAWithB++
			}
		}
		uniqueWords[w] = true
	}
}

func (b *Bayesian) ProbAWhenB(word string) *big.Rat{
	probA := big.NewRat(b.countA, b.totalCount)
	v := b.words[word]
	probBWhenA := big.NewRat(v.countAWithB, b.countA)
	probB := big.NewRat(v.countB, b.totalCount)
	br := &BayesRule{probBWhenA, probA, probB}
	probAWhenB := br.Calculate()
	return probAWhenB
}

func (b *Bayesian) PrintWordProbs() {
	fmt.Printf("Words:\n")
	for k := range b.words {
		probAWhenB := b.ProbAWhenB(k)
		if probAWhenB.Num().Int64() > 0 {
			fmt.Printf("%v - %v\n", k, probAWhenB)
		}
	}
}

func (b *Bayesian) Filter(data Data, threshold *big.Rat) bool {
	total := big.NewRat(1, 1)
	for _, word := range data.GetWords() {
		w := lowerTrim(word)
		_, has := b.words[w]
		if has {
			probAWhenB := b.ProbAWhenB(word)
			one := big.NewRat(1, 1)
			total = total.Mul(total, one.Sub(one, probAWhenB))
		}

	}
	one := big.NewRat(1, 1)
	inverseThreshold := one.Sub(one, threshold)
	return total.Cmp(inverseThreshold) <= 0
}
