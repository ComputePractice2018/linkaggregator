package web

import (
	"net/http"
	"encoding/json"
	"github.com/ComputePractice2018/linkaggregator/backend/domain"
	"log"
	"fmt"
	"strconv"
	"github.com/gorilla/mux"
)

func GetSubscriptions(subList domain.Editable) func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		binaryData, err := json.Marshal(subList.GetSubscriptions())

		if err != nil {
			ProcessError(responseWriter, fmt.Sprintf("Unable to encode data: %v", err), http.StatusInternalServerError)
			return
		}

		responseWriter.Header().Add("Content-Type", "application/json; charset=utf-8")
		responseWriter.WriteHeader(http.StatusOK)

		_, err = responseWriter.Write(binaryData)
		if err != nil {
			ProcessError(responseWriter, fmt.Sprintf("An error occurred on server: %v", err), http.StatusInternalServerError)
		}
	}
}

func AddSubscription(subList domain.Editable, feedList domain.Readable) func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		var subscription domain.Subscription

		err := json.NewDecoder(request.Body).Decode(&subscription)

		if err != nil {
			ProcessError(responseWriter, "Bad request", http.StatusBadRequest)
		}

		filledSubscription, posts := GetSubscriptionByUrl(subscription.Url)
		subscriptionId := subList.AddSubscription(filledSubscription)
		feedList.AddPostsToFeed(subscriptionId, posts)

		responseWriter.Header().Add("Location", request.URL.String()+"/"+strconv.Itoa(subscriptionId))
		responseWriter.WriteHeader(http.StatusCreated)
	}
}

func EditSubscription(subList domain.Editable) func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		var subscription domain.Subscription
		err := json.NewDecoder(request.Body).Decode(&subscription)
		if err != nil {
			ProcessError(responseWriter, fmt.Sprintf("Unable to decode PUT data: %v", err), http.StatusUnsupportedMediaType)
			return
		}

		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])

		if err != nil {
			ProcessError(responseWriter, fmt.Sprintf("Incorrect subscribtion id: %v", err), http.StatusBadRequest)
			return
		}

		err = subList.EditSubscription(subscription, id)
		if err != nil {
			ProcessError(responseWriter, fmt.Sprintf("Incorrect subscribtion id: %v", err), http.StatusBadRequest)
			return
		}

		responseWriter.Header().Add("Location", request.URL.String())
		responseWriter.WriteHeader(http.StatusAccepted)
	}
}

func DeleteSubscription(subList domain.Editable) func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])

		if err != nil {
			ProcessError(responseWriter, fmt.Sprintf("Incorrect subscribtion id: %v", err), http.StatusBadRequest)
			return
		}

		err = subList.RemoveSubscription(id)

		if err != nil {
			ProcessError(responseWriter, fmt.Sprintf("Incorrect subscribtion id: %v", err), http.StatusBadRequest)
			return
		}

		responseWriter.Header().Add("Location", request.URL.String())
		responseWriter.WriteHeader(http.StatusNoContent)
	}
}

func ProcessError(responseWriter http.ResponseWriter, error string, code int) {
	errorMessage := fmt.Sprintf(error, code)
	http.Error(responseWriter, errorMessage, code)
	log.Print(errorMessage)
}
