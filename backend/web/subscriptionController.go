package web

import (
	"net/http"
	"encoding/json"
	"github.com/ComputePractice2018/linkaggregator/backend/domain"
	"log"
	"fmt"
)

func GetSubscriptions(responseWriter http.ResponseWriter, request *http.Request) {

	binaryData, err := json.Marshal(domain.Subscriptions)

	if err != nil {
		ProcessBadRequest(responseWriter, err)
		return
	}

	responseWriter.Header().Add("Content-Type", "application/json; charset=utf-8")
	responseWriter.WriteHeader(http.StatusOK)

	_, err = responseWriter.Write(binaryData)
	if err != nil {
		ProcessError(responseWriter, err)
	}
}

func AddSubscription(responseWriter http.ResponseWriter, request *http.Request) {
	var url interface{}

	err := json.NewDecoder(request.Body).Decode(&url)

	if err != nil {
		ProcessBadRequest(responseWriter, err)
	}

	log.Printf("%+v", url)
	responseWriter.WriteHeader(http.StatusCreated)
}

func ProcessError(responseWriter http.ResponseWriter, err error) {
	errorMessage := fmt.Sprintf("An error occurred on server: %v", err)
	http.Error(responseWriter, errorMessage, http.StatusInternalServerError)
	log.Print(errorMessage)
}

func ProcessBadRequest(responseWriter http.ResponseWriter, err error) {
	errorMessage := fmt.Sprintf("Bad request: %v", err)
	http.Error(responseWriter, errorMessage, http.StatusBadRequest)
	log.Print(errorMessage)
}
