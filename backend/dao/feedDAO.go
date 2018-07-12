package dao

import (
	"github.com/ComputePractice2018/linkaggregator/backend/domain"
	"github.com/jinzhu/gorm"
)

type FeedDAO struct {
	connection *gorm.DB
}

func NewFeedDAO(connection *gorm.DB) *FeedDAO {
	return &FeedDAO{connection: connection}
}

func (feedDao *FeedDAO) GetFeed() []domain.FeedPost {
	var feed []domain.FeedPost
	feedDao.connection.Find(&feed)
	return feed
}

func (feedDao *FeedDAO) GetFeedBySubscriptionId(id int) []domain.FeedPost {
	var feed []domain.FeedPost
	feedDao.connection.Where("subscription_id = ?", id).Find(&feed)
	return feed
}

func (feedDao *FeedDAO) AddPostsToFeed(id int, posts []domain.FeedPost) {
	for _, value := range posts {
		value.SubscriptionId = id
		feedDao.connection.Save(&value)
	}
}
