package models

// AboutResponse defines the structure of our JSON response
type AboutResponse struct {
	Message   string `json:"message"`
	Version   string `json:"version"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}
