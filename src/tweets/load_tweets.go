package tweets

import (
	. "json"
	. "fmt"
	"os"
)

type Tweet struct {
	Topic string "topic"
	Favorited bool "favorited"
	Truncated bool "truncated"
	Text string "text"
	InReplyToStatusId int64 "in_reply_to_status_id"
	User string "user"
	Id int64 "id"
	Author string "author"
	Source string "source"
	InReplyToScreenName string "in_reply_to_screen_name"
	InReplyToUserId int64 "in_reply_to_user_id"
}

type Tweets struct {
	Data []Tweet "tweets"
}

func GetTweets(fileName string) (tweets *Tweets, err os.Error) {
	var file *os.File
	file, err = os.Open(fileName)
	if err != nil {
		Fprintf(os.Stderr, "Could not open file %v - %v\n", fileName, err)
		return
	}

	decoder := NewDecoder(file)
	tweets = new(Tweets)
	err = decoder.Decode(tweets)
	if err != nil {
		Fprintf(os.Stderr, "Could not decode tweets: %v\n", err)
	}
	return
}

func (t *Tweet) Print(prefix string) {
	actualPrefix := ""
	if len(prefix) > 0 {
		actualPrefix = prefix + " "
	}
	Printf("%s%v - %v\n", actualPrefix, t.Author, t.Text)
}

func (t *Tweets) Print() {
	for i, t := range t.Data {
		t.Print(Sprintf("%d", i))
	}
}
