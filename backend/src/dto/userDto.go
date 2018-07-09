package dto

import (
	"../../src/model"
	"gopkg.in/mgo.v2/bson"
)

type UserDto struct {
	Id    bson.ObjectId `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
}

func UserInfo(user model.User) UserDto {
	return UserDto{user.Id, user.Login, user.Email}
}
