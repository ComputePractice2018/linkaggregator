package service

import (
	"../dao/subscription"
	"../dao/userSubscription"
	"../../src/model"
	"../../src/domain"
	"../../src/dao/userRepository"
	"../../src/job"
)

func SubscribeUserToRSS(userLogin string, rssUrl string) {

	if !subscription.IsSourceExist(rssUrl) {
		job.RetrieveRSSFeedInfo(rssUrl)
	}

	user, _ := userRepository.FindByLogin(userLogin)

	subscriptionEntity, _ := subscription.FindByUrl(rssUrl)

	if userSubscription.IsUserFeedExist(user.Id) {
		userSubscription.AddFeedSubscribtion(user.Id, model.UserSubscriptionInfo{SubscriptionId: subscriptionEntity.Id, Viewed: false})
	} else {
		userSubscription.Insert(model.UserSubscription{Id: user.Id, Subscriptions: []model.UserSubscriptionInfo{{SubscriptionId: subscriptionEntity.Id, Viewed: false}}})
	}
}

func GetUserSubscriptions(userLogin string) []domain.UserSubscription {
	user, _ := userRepository.FindByLogin(userLogin)

	userSubscriptionsEntity, _ := userSubscription.FindById(user.Id)

	var result []domain.UserSubscription

	for _, v := range userSubscriptionsEntity.Subscriptions {
		subscriptionEntity, _ := subscription.FindById(v.SubscriptionId)
		result = append(result, domain.UserSubscription{Url: subscriptionEntity.Url, Title: subscriptionEntity.Title, Updated: subscriptionEntity.Updated})
	}

	return result
}
