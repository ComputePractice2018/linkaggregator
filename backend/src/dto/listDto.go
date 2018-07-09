package dto

type ListDto struct {
	Success bool          `json:"success"`
	Items   []interface{} `json:"items"`
}

func List(items []interface{}) ListDto {
	return ListDto{true, items}
}
