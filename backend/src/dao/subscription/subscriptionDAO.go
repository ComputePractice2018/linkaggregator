package subscription

import (
	"gopkg.in/mgo.v2/bson"
	"../../model"
	"../core"
	"time"
)

const (
	COLLECTION = "subscriptions"
)

func FindById(id bson.ObjectId) (model.Subscription, error) {
	var source model.Subscription
	err := core.GetDBConnection().DBInstance.C(COLLECTION).FindId(id).One(&source)
	return source, err
}

func FindByUrl(login string) (model.Subscription, error) {
	var source model.Subscription
	err := core.GetDBConnection().DBInstance.C(COLLECTION).Find(bson.M{"url": login}).One(&source)
	return source, err
}

func Insert(feed model.Subscription) error {
	err := core.GetDBConnection().DBInstance.C(COLLECTION).Insert(&feed)
	return err
}

func Delete(feed model.Subscription) error {
	err := core.GetDBConnection().DBInstance.C(COLLECTION).Remove(&feed)
	return err
}

func UpdateLastUpdateDate(id bson.ObjectId) error {
	err := core.GetDBConnection().DBInstance.C(COLLECTION).Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"added": time.Now()}})
	return err
}

func IsSourceExist(url string) bool {
	count, err := core.GetDBConnection().DBInstance.C(COLLECTION).Find(bson.M{"url": url}).Count()

	if err != nil {
		panic(err)
	}

	return count > 0
}
