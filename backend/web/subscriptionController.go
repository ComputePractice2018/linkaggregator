package web

import (
	"net/http"
	"encoding/json"
	"github.com/ComputePractice2018/linkaggregator/backend/domain"
	"log"
)

func GetSubscriptions(responseWriter http.ResponseWriter, request *http.Request) {
	binaryData, err := json.Marshal(domain.Subscriptions)

	if err != nil {
		responseWriter.Header().Add("Content-Type", "application/json;charset=UTF-8")
		responseWriter.WriteHeader(http.StatusInternalServerError)
		unsuccessfulResponseBody, _ := json.Marshal(domain.Error("Unexpected error"))
		responseWriter.Write(unsuccessfulResponseBody)
		return
	}

	responseWriter.Header().Add("Content-Type", "application/json; charset=utf-8")
	responseWriter.WriteHeader(http.StatusOK)

	_, err = responseWriter.Write(binaryData)
	if err != nil {
		panic("Cannot write response")
	}
}

func AddSubscription(responseWriter http.ResponseWriter, request *http.Request) {
	var subscriptionUrl domain.Url

	err := json.NewDecoder(request.Body).Decode(&subscriptionUrl)

	if err != nil {
		unsuccessfulResponseBody, _ := json.Marshal(domain.Error("Cannot parse data"))
		responseWriter.WriteHeader(http.StatusUnsupportedMediaType)
		responseWriter.Write(unsuccessfulResponseBody)
		return
	}

	log.Printf("%+v", subscriptionUrl)
	responseWriter.WriteHeader(http.StatusCreated)

	successResponseBody, _ := json.Marshal(domain.Success())

	_, err = responseWriter.Write(successResponseBody)

	if err != nil {
		panic("Cannot write response")
	}
}
