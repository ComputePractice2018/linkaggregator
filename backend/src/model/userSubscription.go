package model

import (
	"gopkg.in/mgo.v2/bson"
)

type UserSubscriptionInfo struct {
	SubscriptionId bson.ObjectId `bson:"subscription_id" json:"subscription_id"`
	Viewed         bool          `bson:"viewed" json:"viewed"`
}

type UserSubscription struct {
	Id            bson.ObjectId          `bson:"_id" json:"id"`
	Subscriptions []UserSubscriptionInfo `bson:"subscriptions" json:"subscriptions"`
}
