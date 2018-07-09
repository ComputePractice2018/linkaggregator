package domain

import "time"

type UserSubscription struct {
	Url     string
	Title   string
	Updated time.Time
}
