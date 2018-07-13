package dao

import (
	"github.com/ComputePractice2018/linkaggregator/backend/domain"
	"github.com/jinzhu/gorm"
	"fmt"
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

func (feedDao *FeedDAO) RemoveFeedPostsBySubscriptionId(id int) error {
	var feedPosts []domain.FeedPost
	feedDao.connection.Where("subscription_id = ?", id).Find(&feedPosts)

	if len(feedPosts) == 0 {
		return fmt.Errorf("can't find record #%d", id)
	}

	feedDao.connection.Delete(&feedPosts)
	return nil
}
