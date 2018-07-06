package dto

type ErrorDto struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Error(string string) ErrorDto {
	return ErrorDto{false,
		string}
}
