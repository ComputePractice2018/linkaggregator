package domain

import (
	"testing"
	"time"
)

var testSubscriptions = []Subscription{
	{
		Id:      0,
		Url:     "https://habr.com/rss/interesting",
		Title:   "Хабр / Интересные публикации",
		Updated: time.Now(),
	},
	{
		Id:      1,
		Url:     "http://www.dailymail.co.uk/home/index.rss",
		Title:   "Home | Mail Online",
		Updated: time.Now(),
	},
}

func TestAddSubscription(t *testing.T) {
	subList := NewSubscriptionList()

	subList.AddSubscription("https://habr.com/rss/interesting")

	subList.GetSubscriptions()[0].Updated = time.Time{}
	testSubscriptions[0].Updated = time.Time{}

	if subList.GetSubscriptions()[0] != testSubscriptions[0] {
		t.Errorf("AddSubscription is not working")
	}
}

func TestEditSubscription(t *testing.T) {
	subList := NewSubscriptionList()
	subList.AddSubscription("https://habr.com/rss/interesting")

	err := subList.EditSubscription(testSubscriptions[1], 0)

	if subList.GetSubscriptions()[0] != testSubscriptions[1] {
		t.Errorf("EditSubscription is not working")
	}
	if err != nil {
		t.Errorf("Unexpected EditSubscription error")
	}

	err = subList.EditSubscription(testSubscriptions[1], -1)
	if err == nil {
		t.Errorf("Nothandled out of range error")
	}

}

func TestRemoveSubscription(t *testing.T) {
	subList := NewSubscriptionList()
	subList.AddSubscription(testSubscriptions[0].Url)
	subList.AddSubscription(testSubscriptions[1].Url)

	err := subList.RemoveSubscription(0)


	subList.GetSubscriptions()[0].Updated = time.Time{}
	testSubscriptions[1].Updated = time.Time{}

	if subList.GetSubscriptions()[0] != testSubscriptions[1] {
		t.Errorf("RemoveSubscription is not working")
	}
	if err != nil {
		t.Errorf("Unexpected RemoveSubscription error")
	}

	err = subList.RemoveSubscription(-1)
	if err == nil {
		t.Errorf("Nothandled out of range error")
	}

}
