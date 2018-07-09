package dto

import "fmt"

type ItemDto struct {
	Success bool `json:"success"`
	item    interface{}
}

func Item(item interface{}) ItemDto {
	itemFields := fmt.Sprintf("%v", item)
	return ItemDto{true, itemFields}
}
