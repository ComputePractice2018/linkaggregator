package dto

type SuccessDto struct {
	Success bool `json:"success"`
}

func Success() SuccessDto {
	return SuccessDto{true}
}
