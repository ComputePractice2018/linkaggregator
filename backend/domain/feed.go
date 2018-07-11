package domain

import (
	"time"
	"fmt"
)

type Feed struct {
	Subscription Subscription `json:"subscription"`
	FeedPosts    []FeedPost   `json:"posts"`
}

type FeedPost struct {
	Title       string    `json:"title"`
	Thumbnail   string    `json:"thumbnail"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	PubDate     time.Time `json:"pubDate"`
}

var feedPosts map[int][]FeedPost

func GetFeed() []Feed {
	var result []Feed
	for k, v := range feedPosts {
		result = append(result, Feed{Subscription: GetSubscriptionById(k), FeedPosts: v})
	}
	return result
}

func GetFeedPostsById(id int) (Feed, error) {
	if id < 0 || id >= len(feedPosts) {
		return Feed{}, fmt.Errorf("incorrect subscription id")
	}
	return Feed{Subscription: GetSubscriptionById(id), FeedPosts: feedPosts[id]}, nil
}

func AddPostsToFeedById(subscriptionId int, posts []FeedPost) {
	if len(feedPosts) == 0 {
		feedPosts = make(map[int][]FeedPost)
	}
	feedPosts[subscriptionId] = posts
}
