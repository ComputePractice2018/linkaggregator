package job

import (
	"../../src/model"
	"../../src/domain"
	"../dao/subscription"
	"../../src/dao/post"
	"gopkg.in/mgo.v2/bson"
	"encoding/xml"
	"net/http"
	"time"
)

func RetrieveRSSFeedInfo(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	rss := domain.Rss{}

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)

	if err != nil {
		panic("Cannot decode rss response")

	}

	feedObjectId := bson.NewObjectId()
	subscription.Insert(model.Subscription{Id: feedObjectId, Url: url, Title: rss.Channel.Title, Updated: time.Now()})

	go func() {
		for _, item := range rss.Channel.Items {
			timeObject, _ := time.Parse(time.RFC1123, item.PubDate)
			post.Insert(model.FeedPost{bson.NewObjectId(), feedObjectId, item.Title, item.Thumbnail, item.Desc, item.Link, timeObject})
		}
	}()
}
