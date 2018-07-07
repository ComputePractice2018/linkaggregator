package userRepository

import (
	"gopkg.in/mgo.v2/bson"
	"../../model"
	"../core"
)

const (
	COLLECTION = "users"
)

func FindById(id string) (model.User, error) {
	var user model.User
	err := core.GetDBConnection().DBInstance.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func FindByLogin(login string) (model.User, error) {
	var user model.User
	err := core.GetDBConnection().DBInstance.C(COLLECTION).Find(bson.M{"login": login}).One(&user)
	return user, err
}

func Insert(user model.User) error {
	err := core.GetDBConnection().DBInstance.C(COLLECTION).Insert(&user)
	return err
}

func Delete(user model.User) error {
	err := core.GetDBConnection().DBInstance.C(COLLECTION).Remove(&user)
	return err
}

func Update(user model.User) error {
	err := core.GetDBConnection().DBInstance.C(COLLECTION).UpdateId(user.Id, &user)
	return err
}

func IsUserExist(email string, login string) bool {
	count, err := core.GetDBConnection().DBInstance.C(COLLECTION).Find(bson.M{"$or": []bson.M{{"login": login}, {"email": email}}}).Count()
	if err != nil {
		panic(err)
	}

	return count > 0
}
