package models

// UserResponse is the response structure for the user endpoint
// that suppresses the password field
type UserResponse struct {
	ID            int64  `json:"id"`
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	Status        string `json:"status"`
	StatusID      int64  `json:"status_id"`
	PhoneVerified bool   `json:"phone_verified"`
}
