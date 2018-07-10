package main

import (
	"github.com/ComputePractice2018/linkaggregator/backend/web"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/api/subscription/list", web.GetSubscriptions)
	http.HandleFunc("/api/subscription/add", web.AddSubscription)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
