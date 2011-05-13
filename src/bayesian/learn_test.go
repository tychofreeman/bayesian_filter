package bayesian

import (
	. "testing"
	"big"
	"strings"
)


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

var always = big.NewRat(1, 1)
var usually = big.NewRat(2, 3)
var mostly = big.NewRat(1, 2)
var sometimes = big.NewRat(1, 3)
var never = big.NewRat(0, 1)

func makeTweet(text string) Tweet {
	return Tweet{strings.Fields(text), false}
}

func TestFilter(t *T) {
	b := makeTweets()

	assertFilters(t, b, makeTweet("grinnell has a president"), always)
	assertNotFilters(t, b, makeTweet("five six seven"), sometimes)
	assertFilters(t, b, makeTweet("one four"), sometimes)
}

func assertFilters(t *T, b *Bayesian, data Data, expected *big.Rat) {
	assertFilterPred(t, b, data, expected, true)
}
func assertNotFilters(t *T, b *Bayesian, data Data, expected *big.Rat) {
	assertFilterPred(t, b, data, expected, false)
}
func assertFilterPred(t *T, b *Bayesian, data Data, expected *big.Rat, pred bool) {
	if pred != b.Filter(data, expected) {
		t.Errorf("Data [%v] expected to filter at or above %v, but did not.\n", data, expected)
	}
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
