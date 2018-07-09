package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Subscription struct {
	Id      bson.ObjectId `bson:"_id" json:"id"`
	Url     string        `bson:"url" json:"url"`
	Title   string        `bson:"title" json:"title"`
	Updated time.Time     `bson:"updated" json:"updated"`
}
