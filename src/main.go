package main

import "tweets"
import "bayesian"
import "strings"

type BayesianTweet struct {
	tweets.Tweet
	topic string
}

func (bt *BayesianTweet) GetWords() []string {
	return strings.Fields(bt.Text)
}

func (bt *BayesianTweet) Is() bool {
	return strings.Contains(bt.Topic, bt.topic)
}

func main() {
	t, _ := tweets.GetTweets("data/tweets.txt")

	b := new(bayesian.Bayesian)
	for _, t := range t.Data {
		b.Add(&BayesianTweet{t, "Grinnell"})
	}
	b.PrintWordProbs()
}
