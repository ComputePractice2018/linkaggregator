package main

import (
	"github.com/ComputePractice2018/linkaggregator/backend/web"
	"net/http"
	"log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ComputePractice2018/linkaggregator/backend/dao"
	"flag"
)

func main() {
	connection := flag.String("connection", "linkaggregator:SuperSecretPassword@tcp(localhost:3306)/linkaggregator?parseTime=true", "mysql connection string")
	flag.Parse()

	dbConnection := dao.InitDbConnection(*connection)
	defer dbConnection.Close()

	subscriptionList := dao.NewSubscriptionDAO(dbConnection)
	feedList := dao.NewFeedDAO(dbConnection)

	log.Fatal(http.ListenAndServe(":8080", web.NewRouter(subscriptionList, feedList)))
}
