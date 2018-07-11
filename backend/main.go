package main

import (

	"fmt"
	//"github.com/ComputePractice2018/linkaggregator/backend/utils"
)

//func main() {
	//firstNumber := flag.Int("first", 2, "first multiply attribute")
	//secondNumber := flag.Int("second", 2, "second multiply attribute")

//	fmt.Print("Hello world \n")
	//fmt.Printf("Sum of two int: %d\n", utils.MultiplyTwoInt(*firstNumber, *secondNumber))

	"github.com/gorilla/mux"
	"github.com/ComputePractice2018/linkaggregator/backend/web"
	"net/http"
	"log"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/subscriptions", web.GetSubscriptions).Methods("GET")
	router.HandleFunc("/api/subscriptions", web.AddSubscription).Methods("POST")
	router.HandleFunc("/api/subscriptions/{id}", web.EditSubscription).Methods("PUT")
	router.HandleFunc("/api/subscriptions/{id}", web.DeleteSubscription).Methods("DELETE")

	router.HandleFunc("/api/feed", web.GetFeed).Methods("GET")
	router.HandleFunc("/api/feed/{id}", web.GetFeedById).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}
