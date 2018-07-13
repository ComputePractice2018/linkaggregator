package web

import (
	"net/http"
	"encoding/xml"
	"time"
	"github.com/ComputePractice2018/linkaggregator/backend/domain"
)

func GetSubscriptionByUrl(url string) (domain.Subscription, []domain.FeedPost) {
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

	var posts []domain.FeedPost

	for _, post := range rss.Channel.Items {
		timeObject, _ := time.Parse(time.RFC1123, post.PubDate)
		posts = append(posts, domain.FeedPost{Title: post.Title, Thumbnail: post.Thumbnail, Description: post.Desc, Link: post.Link, PubDate: timeObject})
	}

	return domain.Subscription{Url: url, Title: rss.Channel.Title, Updated: time.Now()}, posts
}
