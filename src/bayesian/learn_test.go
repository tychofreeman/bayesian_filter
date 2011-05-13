package bayesian

import . "testing"
import "big"


type Tweet struct {
	text []string
	isGrinnell bool
}
func (t Tweet) GetWords() []string { return t.text }
func (t Tweet) Is() bool { return t.isGrinnell }

func makeTweets() *Bayesian {
	b := new(Bayesian)
	b.Add(Tweet{[]string{"one", "two"}, false})
	b.Add(Tweet{[]string{"three", "four"}, false})
	b.Add(Tweet{[]string{"one", "two"}, false})
	b.Add(Tweet{[]string{"three", "four"}, false})
	b.Add(Tweet{[]string{"five", "six"}, false})
	b.Add(Tweet{[]string{"seven", "seven"}, false})
	b.Add(Tweet{[]string{"one", "two"}, false})
	b.Add(Tweet{[]string{"one", "two"}, false})
	b.Add(Tweet{[]string{"one", "two"}, false})
	b.Add(Tweet{[]string{"one", "two"}, false})
	b.Add(Tweet{[]string{"one", "two"}, false})
	b.Add(Tweet{[]string{"one", "other"}, false})
	b.Add(Tweet{[]string{"grinnell", "other", "thing"}, true})
	b.Add(Tweet{[]string{"grinnell", "grinnell"}, true})
	b.Add(Tweet{[]string{"grinnell", "grinnell", "four"}, true})
	b.Add(Tweet{[]string{"president", "grinnell", "one"}, true})
	b.Add(Tweet{[]string{"president", "grinnell", "two"}, true})
	b.Add(Tweet{[]string{"president", "grinnell", "three"}, true})

	return b
}

func TestTweets(t *T) {
	b := makeTweets()

	assertProb(t, b, big.NewRat(1, 1), "grinnell")
	assertProb(t, b, big.NewRat(1, 1), "president")
	assertProb(t, b, big.NewRat(1, 9), "one")
	assertProb(t, b, big.NewRat(1, 3), "four")
	assertProb(t, b, big.NewRat(1, 2), "other")
}

func assertProb(t *T, b *Bayesian, expected *big.Rat, word string) {
	actual := b.ProbAWhenB(word)
	if expected.Cmp(actual) != 0 {
		t.Errorf("'%s' should be correlated %v:%v with Grinnell tweets, but was %v:%v\n", 
				word,
				expected.Num(), expected.Denom(),
				actual.Num(), actual.Denom())
	}
}
