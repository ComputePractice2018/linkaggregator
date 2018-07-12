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

var subscriptions []Subscription

func GetSubscriptionById(id int) Subscription {
	if id < 0 || id >= len(subscriptions) {
		fmt.Errorf("incorrect subscription id")
	}
	return subscriptions[id]
}

func GetSubscriptions() []Subscription {
	return subscriptions
}

func AddSubscription(url string) int {
	id := len(subscriptions)
	subscription, posts := getSubscriptionByUrl(url)
	subscription.Id = id
	subscriptions = append(subscriptions, subscription)
	AddPostsToFeedById(id, posts)
	return len(subscriptions)
}

func EditSubscription(subscription Subscription, id int) error {
	if id < 0 || id >= len(subscriptions) {
		return fmt.Errorf("incorrect subscription id")
	}
	subscriptions[id] = subscription
	return nil
}

func RemoveSubscription(id int) error {
	if id < 0 || id >= len(subscriptions) {
		return fmt.Errorf("incorrect subscription id")
	}
	subscriptions = append(subscriptions[:id], subscriptions[id+1:]...)
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
