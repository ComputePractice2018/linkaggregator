package dao

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"../../src/model"
	"../../src/config"
	"sync"
)

type UserDAO struct {
	Server   string
	Database string
}

var instance *UserDAO
var once sync.Once

func GetUserDao() *UserDAO {
	once.Do(func() {
		instance = &UserDAO{}
		dbConfig := config.DatabaseConfig{}
		dbConfig.Read()
		instance.Server = "localhost"    /*dbConfig.Server*/
		instance.Database = "aggregator" /*dbConfig.Database*/
		instance.Connect()
	})
	return instance
}

var db *mgo.Database

const (
	COLLECTION = "users"
)

func (m *UserDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *UserDAO) FindById(id string) (model.User, error) {
	var user model.User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (m *UserDAO) FindByLogin(login string) (model.User, error) {
	var user model.User
	err := db.C(COLLECTION).Find(bson.M{"login": login}).One(&user)
	return user, err
}

func (m *UserDAO) Insert(user model.User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

func (m *UserDAO) Delete(user model.User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

func (m *UserDAO) Update(user model.User) error {
	err := db.C(COLLECTION).UpdateId(user.Id, &user)
	return err
}

func (m *UserDAO) IsUserExist(email string, login string) bool {
	count, err := db.C(COLLECTION).Find(bson.M{"$or": []bson.M{{"login": login}, {"email": email}}}).Count()
	if err != nil {
		panic(err)
	}

	return count > 0
}
