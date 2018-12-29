package models

type Response struct {
	Message string
	Status string
}

func NewResponse(message string, status string) map[string]string {
	return map[string]string{
		"message": message,
		"status":  status,
	}
}
