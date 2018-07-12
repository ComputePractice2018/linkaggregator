package domain

import (
	"time"
	"fmt"
	"net/http"
	"encoding/xml"
)

type Subscription struct {
	Id      int       `json:"id"`
	Url     string    `json:"url"`
	Title   string    `json:"title"`
	Updated time.Time `json:"updated"`
}

type SubscriptionList struct {
	subscriptions []Subscription
}

type Editable interface {
	GetSubscriptions() []Subscription
	GetSubscriptionById(id int) Subscription
	AddSubscription(url string) int
	EditSubscription(subscription Subscription, id int) error
	RemoveSubscription(id int) error
}

func NewSubscriptionList() *SubscriptionList {
	return &SubscriptionList{}
}

func (subList *SubscriptionList) GetSubscriptionById(id int) Subscription {
	if id < 0 || id >= len(subList.subscriptions) {
		fmt.Errorf("incorrect subscription id")
	}
	return subList.subscriptions[id]
}

func (subList *SubscriptionList) GetSubscriptions() []Subscription {
	return subList.subscriptions
}

func (subList *SubscriptionList) AddSubscription(url string) int {
	id := len(subList.subscriptions)
	subscription, posts := getSubscriptionByUrl(url)
	subscription.Id = id
	subList.subscriptions = append(subList.subscriptions, subscription)
	AddPostsToFeedById(posts)
	return len(subList.subscriptions) - 1
}

func (subList *SubscriptionList) EditSubscription(subscription Subscription, id int) error {
	if id < 0 || id >= len(subList.subscriptions) {
		return fmt.Errorf("incorrect subscription id")
	}
	subList.subscriptions[id] = subscription
	return nil
}

func (subList *SubscriptionList) RemoveSubscription(id int) error {
	if id < 0 || id >= len(subList.subscriptions) {
		return fmt.Errorf("incorrect subscription id")
	}
	subList.subscriptions = append(subList.subscriptions[:id], subList.subscriptions[id+1:]...)
	return nil
}

func getSubscriptionByUrl(url string) (Subscription, []FeedPost) {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	rss := Rss{}

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)

	if err != nil {
		panic("Cannot decode rss response")
	}

	var posts []FeedPost

	for _, post := range rss.Channel.Items {
		timeObject, _ := time.Parse(time.RFC1123, post.PubDate)
		posts = append(posts, FeedPost{Title: post.Title, Thumbnail: post.Thumbnail, Description: post.Desc, Link: post.Link, PubDate: timeObject})
	}

	return Subscription{Url: url, Title: rss.Channel.Title, Updated: time.Now()}, posts
}
