package main

import (
	"github.com/ComputePractice2018/linkaggregator/backend/web"
	"net/http"
	"log"
	"github.com/ComputePractice2018/linkaggregator/backend/domain"
)

func main() {
	subscriptionList := domain.NewSubscriptionList()
	log.Fatal(http.ListenAndServe(":8080", web.NewRouter(subscriptionList)))
}
