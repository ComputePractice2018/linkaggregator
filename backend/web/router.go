package web

import (
	"github.com/gorilla/mux"
	"github.com/ComputePractice2018/linkaggregator/backend/domain"
)

func NewRouter(subscriptionList domain.Editable, feedList domain.Readable) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/linkaggregator/subscriptions", GetSubscriptions(subscriptionList)).Methods("GET")
	router.HandleFunc("/api/linkaggregator/subscriptions", AddSubscription(subscriptionList, feedList)).Methods("POST")
	router.HandleFunc("/api/linkaggregator/subscriptions/{id}", EditSubscription(subscriptionList)).Methods("PUT")
	router.HandleFunc("/api/linkaggregator/subscriptions/{id}", DeleteSubscription(subscriptionList)).Methods("DELETE")

	router.HandleFunc("/api/linkaggregator/feed", GetFeed(subscriptionList, feedList)).Methods("GET")
	router.HandleFunc("/api/linkaggregator/feed/{id}", GetFeedById(subscriptionList, feedList)).Methods("GET")
	return router
}
