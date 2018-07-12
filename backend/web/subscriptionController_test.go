package web

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
	"github.com/ComputePractice2018/linkaggregator/backend/domain"
)

var testSubscriptionUrl = "{\"url\":\"https://habr.com/rss/interesting/\"}"

var TestSubscription = "{\"url\": \"https: //habr.com/rss/interesting/\",\"title\": \"Хабр / Интересные публикации\",\"updated\": \"2018-07-10T21:56:22.738946318+03:00\"}"

func TestCrudInSubscriptionController(t *testing.T) {
	subList := domain.NewSubscriptionList()
	feedList := domain.NewFeedList()

	router := NewRouter(subList, feedList)

	req, err := http.NewRequest("GET", "/api/linkaggregator/subscriptions", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Error("Expected 200 code")
	}
	if resp.Header.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Error("Expected json MIME-type")
	}

	testData := strings.NewReader(testSubscriptionUrl)
	req, err = http.NewRequest("POST", "/api/linkaggregator/subscriptions", testData)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 code (gotten: %d)", resp.StatusCode)
	}

	if resp.Header.Get("Location") != "/api/linkaggregator/subscriptions/0" {
		t.Error("Expected another location")
	}

	if len(subList.GetSubscriptions()) != 1 {
		t.Error("Expected new value")
	}

	testData = strings.NewReader(TestSubscription)
	req, err = http.NewRequest("PUT", "/api/linkaggregator/subscriptions/0", testData)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Expected 201 code (gotten: %d)", resp.StatusCode)
	}
	if resp.Header.Get("Location") != "/api/linkaggregator/subscriptions/0" {
		t.Error("Expected another location")
	}

	if len(subList.GetSubscriptions()) != 1 {
		t.Error("Expected old value")
	}

	req, err = http.NewRequest("DELETE", "/api/linkaggregator/subscriptions/0", nil)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected 204 code (gotten: %d)", resp.StatusCode)
	}
	if len(subList.GetSubscriptions()) != 0 {
		t.Error("Expected old value")
	}
}
