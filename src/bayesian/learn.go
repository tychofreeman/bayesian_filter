package bayesian

import "fmt"
import s "strings"
import u "unicode"

type division struct {
	numWhenIs int
	numWhenNotIs int
}

type Bayesian struct {
	totalNum int
	appearances map[string]*division
}

func trim(word string) string {
	return s.TrimFunc(word, func(rune int) bool { return !u.IsLetter(rune) })
}

func (b *Bayesian) Add(data Data) {
	words := data.GetWords()
	if b.appearances == nil {
		b.appearances = map[string]*division{}
	}
	for _, rawWord := range words {
		w := trim(rawWord)
		b.totalNum++
		d, err := b.appearances[w]
		if err == false {
			d = new(division)
			b.appearances[w] = d

		}
		if data.Is() {
			d.numWhenIs++
		} else {
			d.numWhenNotIs++
		}
	}
}

func (b *Bayesian) Learn() {
}

func (b *Bayesian) PrintWordProbs() {
	fmt.Printf("Words:\n")
	for k, v := range b.appearances {
		fmt.Printf("%v - %v !%v\n", k, v.numWhenIs, v.numWhenNotIs)
	}
}
