package domain

import (
	"time"
	"fmt"
)

type EditableFeed interface {
	GetFeed() []FeedPost
	GetFeedBySubscriptionId(id int) []FeedPost
	AddPostsToFeed(id int, posts []FeedPost)
	RemoveFeedPostsBySubscriptionId(id int) error
}

type FeedPost struct {
	Id             int       `json:"id"`
	SubscriptionId int       `json:"subscription_id"`
	Title          string    `json:"title" sql:"type:text;"`
	Thumbnail      string    `json:"thumbnail"`
	Description    string    `json:"description" sql:"type:text;"`
	Link           string    `json:"link" sql:"type:text;"`
	PubDate        time.Time `json:"pubDate"`
}

type FeedList struct {
	feedList [][]FeedPost
}

func NewFeedList() *FeedList {
	return &FeedList{}
}

func (fl *FeedList) GetFeed() []FeedPost {
	var result []FeedPost

	for _, value := range fl.feedList {
		for _, value := range value {
			result = append(result, value)
		}
	}

	return result
}

func (fl *FeedList) GetFeedBySubscriptionId(id int) []FeedPost {
	if id < 0 || id >= len(fl.feedList) {
		fmt.Errorf("incorrect subscription id")
	}
	return fl.feedList[id]
}

func (fl *FeedList) AddPostsToFeed(id int, posts []FeedPost) {
	fl.feedList = append(fl.feedList, posts)
}

func (fl *FeedList) RemoveFeedPostsBySubscriptionId(id int) error {

	if id < 0 || id >= len(fl.feedList) {
		return fmt.Errorf("incorrect subscription id")
	}

	fl.feedList = append(fl.feedList[:id], fl.feedList[id+1:]...)
	return nil
}
