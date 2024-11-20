package dto

type RequestError struct {
	Error string `json:"error"`
}

func Error(msg string) RequestError {
	return RequestError{Error: msg}
}
