package post

import (
	"gopkg.in/mgo.v2/bson"
	"../../model"
	"../core"
)

const (
	COLLECTION = "posts"
)

func FindById(id bson.ObjectId) (model.FeedPost, error) {
	var source model.FeedPost
	err := core.GetDBConnection().DBInstance.C(COLLECTION).FindId(id).One(&source)
	return source, err
}

func FindByUrl(login string) (model.FeedPost, error) {
	var source model.FeedPost
	err := core.GetDBConnection().DBInstance.C(COLLECTION).Find(bson.M{"url": login}).One(&source)
	return source, err
}

func Insert(feed model.FeedPost) error {
	err := core.GetDBConnection().DBInstance.C(COLLECTION).Insert(&feed)
	return err
}

func Delete(feed model.FeedPost) error {
	err := core.GetDBConnection().DBInstance.C(COLLECTION).Remove(&feed)
	return err
}

func Update(feed model.FeedPost) error {
	err := core.GetDBConnection().DBInstance.C(COLLECTION).UpdateId(feed.Id, &feed)
	return err
}
