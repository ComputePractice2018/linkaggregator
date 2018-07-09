package dto

import (
	"../../src/model"
)

type UserDto struct {
	Id    string `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
}

func UserInfo(user model.User) UserDto {
	return UserDto{user.Id.String(), user.Login, user.Email}
}
