package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type FeedPost struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	FeedId      bson.ObjectId `bson:"feed_id" json:"feedId"`
	Title       string        `bson:"title" json:"title"`
	Thumbnail   string        `bson:"thumbnail" json:"thumbnail"`
	Description string        `bson:"description" json:"description"`
	Link        string        `bson:"link" json:"link"`
	PubDate     time.Time     `bson:"pubDate" json:"pubDate"`
}
