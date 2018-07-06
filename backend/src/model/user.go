package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id           bson.ObjectId `bson:"_id" json:"id"`
	Login        string        `bson:"login" json:"login"`
	Password     string        `bson:"password" json:"password"`
	Email        string        `bson:"email" json:"email"`
	Registration time.Time     `bson:"reg" json:"reg"`
}
