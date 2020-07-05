package structs

import "time"
// UserResponse Get the request body of a User Response.
type UserResponse struct {
	UserID    uint32    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Phone     string    `json:"phone_number"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PostUserRequest Get the request body of a User POST.
type PostUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Cpf   string `json:"cpf"`
	Phone string `json:"phone_number"`
}

// UserResponse Store the order-api info.
type OrderApi struct {
	StartOfUrl    string
	EndOfUrl    string
}