package web

import (
	"net/http"
	"encoding/json"
	"github.com/ComputePractice2018/linkaggregator/backend/domain"
	"fmt"
	"github.com/gorilla/mux"
	"strconv"
)

func GetFeed(subList domain.Editable) func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {

		binaryData, err := json.Marshal(domain.GetFeed(subList))

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

func GetFeedById(subList domain.Editable) func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {

		vars := mux.Vars(request)
		id, err := strconv.Atoi(vars["id"])

		if err != nil {
			ProcessError(responseWriter, fmt.Sprintf("Incorrect subscribtion id: %v", err), http.StatusBadRequest)
			return
		}

		posts, err := domain.GetFeedPostsById(id, subList)

		binaryData, err := json.Marshal(posts)

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
