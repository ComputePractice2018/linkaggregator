package userSubscription

import (
	"gopkg.in/mgo.v2/bson"
	"../../model"
	"../core"
)

const (
	COLLECTION = "user_subscriptions"
)

func FindById(id bson.ObjectId) (model.UserSubscription, error) {
	var source model.UserSubscription
	err := core.GetDBConnection().DBInstance.C(COLLECTION).FindId(id).One(&source)
	return source, err
}

func Insert(feed model.UserSubscription) error {
	err := core.GetDBConnection().DBInstance.C(COLLECTION).Insert(&feed)
	return err
}

func AddFeedSubscribtion(userId bson.ObjectId, feedInfo model.UserSubscriptionInfo) error {
	err := core.GetDBConnection().DBInstance.C(COLLECTION).Update(bson.M{"_id": userId}, bson.M{"$push": bson.M{"feeds": feedInfo}})
	return err
}

func Delete(feed model.UserSubscription) error {
	err := core.GetDBConnection().DBInstance.C(COLLECTION).Remove(&feed)
	return err
}

func IsUserFeedExist(uid bson.ObjectId) bool {
	count, err := core.GetDBConnection().DBInstance.C(COLLECTION).FindId(uid).Count()
	if err != nil {
		panic(err)
	}

	return count > 0
}
