package structs

import "time"

type UserResponse struct {
	UserID    uint32    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Phone     string    `json:"phone_number"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserApi struct {
	Url     string
}



