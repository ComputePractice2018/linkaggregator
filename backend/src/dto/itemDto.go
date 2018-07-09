package dto

type ItemDto struct {
	Success bool `json:"success"`
	Item    interface{} `json:"item"`
}

func Item(item interface{}) ItemDto {
	return ItemDto{true, item}
}
