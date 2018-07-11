package main

import (
	"github.com/gorilla/mux"
	"github.com/ComputePractice2018/linkaggregator/backend/web"
	"net/http"
	"log"
	"github.com/ComputePractice2018/linkaggregator/backend/domain"
)

func main() {
	router := mux.NewRouter()

	subscriptionList := domain.NewSubscriptionList()

	router.HandleFunc("/api/subscriptions", web.GetSubscriptions(subscriptionList)).Methods("GET")
	router.HandleFunc("/api/subscriptions", web.AddSubscription(subscriptionList)).Methods("POST")
	router.HandleFunc("/api/subscriptions/{id}", web.EditSubscription(subscriptionList)).Methods("PUT")
	router.HandleFunc("/api/subscriptions/{id}", web.DeleteSubscription(subscriptionList)).Methods("DELETE")

	router.HandleFunc("/api/feed", web.GetFeed(subscriptionList)).Methods("GET")
	router.HandleFunc("/api/feed/{id}", web.GetFeedById(subscriptionList)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
