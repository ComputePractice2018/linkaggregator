package domain

import (
	"time"
)

type Subscription struct {
	Id      string    `json:"id"`
	Url     string    `json:"url"`
	Title   string    `json:"title"`
	Updated time.Time `json:"updated"`
}

var Subscriptions = []Subscription{{
	Id:      "00db4f604a1068ed4dfc",
	Url:     "https://habr.com/rss",
	Title:   "Хабр / Интересные публикации",
	Updated: time.Now()}}
