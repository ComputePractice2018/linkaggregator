package core

import (
	"sync"
	"gopkg.in/mgo.v2"
	"../../config"
	"log"
)

type SharedDbConnection struct {
	Server     string
	Database   string
	DBInstance *mgo.Database
}

var instance *SharedDbConnection
var once sync.Once

func GetDBConnection() *SharedDbConnection {
	once.Do(func() {
		instance = &SharedDbConnection{}
		dbConfig := config.DatabaseConfig{}
		dbConfig.Read()
		instance.Server = "localhost"    /*dbConfig.Server*/
		instance.Database = "aggregator" /*dbConfig.Database*/
		instance.Connect()
	})
	return instance
}

func (m *SharedDbConnection) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	m.DBInstance = session.DB(m.Database)
}
