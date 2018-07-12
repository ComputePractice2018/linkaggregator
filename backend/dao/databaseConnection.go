package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/ComputePractice2018/linkaggregator/backend/domain"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDbConnection(connection string) (*gorm.DB) {
	database, err := gorm.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&domain.Subscription{}, &domain.FeedPost{})
	return database
}
