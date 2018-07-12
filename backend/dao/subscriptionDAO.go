package dao

import (
	"github.com/ComputePractice2018/linkaggregator/backend/domain"
	"github.com/jinzhu/gorm"
	"fmt"
)

type SubscriptionDAO struct {
	connection *gorm.DB
}

func NewSubscriptionDAO(dconnection *gorm.DB) *SubscriptionDAO {
	return &SubscriptionDAO{connection: dconnection}
}

func (subDAO *SubscriptionDAO) GetSubscriptions() []domain.Subscription {
	var subscriptions []domain.Subscription
	subDAO.connection.Find(&subscriptions)
	return subscriptions
}

func (subDAO *SubscriptionDAO) GetSubscriptionById(id int) domain.Subscription {
	var subscription domain.Subscription
	subDAO.connection.Where("id = ?", id).Find(&subscription)
	return subscription
}

func (subDAO *SubscriptionDAO) AddSubscription(subscription domain.Subscription) int {
	subDAO.connection.Create(&subscription)
	return subscription.Id
}

func (subDAO *SubscriptionDAO) EditSubscription(subscription domain.Subscription, id int) error {
	var subscriptions []domain.Subscription
	subDAO.connection.Where("id = ?", id).Find(&subscriptions)

	if len(subscriptions) == 0 {
		return fmt.Errorf("can't find record #%d", id)
	}

	subscription.Id = subscriptions[0].Id
	subDAO.connection.Save(&subscription)

	return subDAO.connection.Error
}

func (subDAO *SubscriptionDAO) RemoveSubscription(id int) error {
	var subscriptions []domain.Subscription
	subDAO.connection.Where("id = ?", id).Find(&subscriptions)

	if len(subscriptions) == 0 {
		return fmt.Errorf("can't find record #%d", id)
	}

	subDAO.connection.Delete(&subscriptions[0])
	return subDAO.connection.Error
}
