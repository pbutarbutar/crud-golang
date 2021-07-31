package entity

//ApiResponse is for response
type ApiResponse struct {
	Success bool                   `json:"success"`
	Status  int                    `json:"statusCode"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Errors  map[string]string      `json:"errors,omitempty"`
}

//Error is struct error
type Error struct {
	Field   string `json:"field,omitempty"`
	Value   string `json:"value,omitempty"`
	Message string `json:"message,omitempty"`
}
