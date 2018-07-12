package domain

import (
	"time"
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
