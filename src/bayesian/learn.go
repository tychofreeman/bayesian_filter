package bayesian

import "fmt"
import s "strings"
import u "unicode"
import "big"

type division struct {
	numWhenIs int64
	numWhenNotIs int64
	msgsWithWord int64
	grinnellMsgsWithWord int64
	probKAppears *big.Rat
	probKAppearsWhenIsGrinnell *big.Rat
	probAWhenB *big.Rat
}

type Bayesian struct {
	totalWords int64
	totalWordsWhenIsGrinnell int64
	totalMsgs int64
	totalMsgsWhenIsGrinnell int64
	probIsGrinnell *big.Rat
	appearances map[string]*division
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
	}[word]
}

func (b *Bayesian) Add(data Data) {
	words := data.GetWords()
	b.totalMsgs++
	if data.Is() {
		b.totalMsgsWhenIsGrinnell++
	}
	if b.appearances == nil {
		b.appearances = map[string]*division{}
	}
	uniqueWords := make(map[string]bool)
	for _, rawWord := range words {
		w := lowerTrim(rawWord)
		if shouldIgnore(w) {
			break
		}
		b.totalWords++
		d, err := b.appearances[w]
		if err == false {
			d = new(division)
			b.appearances[w] = d

		}
		if _, defd := uniqueWords[w]; defd == false {
			d.msgsWithWord++
			if  data.Is() {
				d.grinnellMsgsWithWord++
			}
		}
		if data.Is() {
			d.numWhenIs++
			b.totalWordsWhenIsGrinnell++
		} else {
			d.numWhenNotIs++
		}
		uniqueWords[w] = true
	}
}

func (b *Bayesian) Learn() {
	b.probIsGrinnell = big.NewRat(b.totalMsgsWhenIsGrinnell, b.totalMsgs)
	for k, v := range b.appearances {
		probKAppears := big.NewRat(v.numWhenIs, v.numWhenIs + v.numWhenNotIs)
		probIsGrinnellWhenKAppears := big.NewRat(v.grinnellMsgsWithWord, v.msgsWithWord)
		br := &BayesRule{probIsGrinnellWhenKAppears, probKAppears, b.probIsGrinnell}
		v.probAWhenB = br.Calculate()
		fmt.Printf("\t%v - %v = (%v*%v)/%v\n", k, v.probAWhenB, probIsGrinnellWhenKAppears, probKAppears, b.probIsGrinnell)
	}
}

func (b *Bayesian) PrintWordProbs() {
	fmt.Printf("Words:\n")
	fmt.Printf("Prob Is? %v\n", b.probIsGrinnell)
	for k, v := range b.appearances {
		if v.numWhenIs > 0 {
			fmt.Printf("%v - %v\n", k, v.probAWhenB)
		}
	}
}

func (b *Bayesian) Filter(data Data, threshold *big.Rat) bool {
	total := big.NewRat(1, 1)
	for _, word := range data.GetWords() {
		w := lowerTrim(word)
		v, has := b.appearances[w]
		if has {
			total = total.Mul(total, v.probAWhenB)
		}

	}

	return total.Cmp(threshold) >= 0
}
