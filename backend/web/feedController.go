package web

import (
	"net/http"
	"github.com/ComputePractice2018/linkaggregator/backend/domain"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"strconv"
)

type FullFeed struct {
	Feeds []AssembledFeed `json:"feed"`
}

type AssembledFeed struct {
	Subscription domain.Subscription `json:"subscription"`
	FeedPosts    []domain.FeedPost   `json:"posts"`
}

func GetFeed(subList domain.Editable, feedList domain.Readable) func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {

		var result FullFeed

		subs := subList.GetSubscriptions()

		for _, value := range subs {
			result.Feeds = append(result.Feeds, AssembledFeed{
				Subscription: value,
				FeedPosts:    feedList.GetFeedBySubscriptionId(value.Id),
			})
		}

		binaryData, err := json.Marshal(result)

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

func GetFeedById(subList domain.Editable, feedList domain.Readable) func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])

		if err != nil {
			ProcessError(responseWriter, fmt.Sprintf("Incorrect subscribtion id: %v", err), http.StatusBadRequest)
			return
		}

		binaryData, err := json.Marshal(AssembledFeed{
			Subscription: subList.GetSubscriptionById(id),
			FeedPosts:    feedList.GetFeedBySubscriptionId(id),
		})

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
