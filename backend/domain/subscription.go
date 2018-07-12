package domain

import (
	"time"
	"fmt"
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
	AddSubscription(subscription Subscription) int
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

func (subList *SubscriptionList) AddSubscription(subscription Subscription) int {
	id := len(subList.subscriptions)
	subscription.Id = id
	subList.subscriptions = append(subList.subscriptions, subscription)
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
