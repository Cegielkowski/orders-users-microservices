package structs

import (
	"time"
)

type OrderResponse struct {
	OrderID         uint32    `json:"id"`
	UserID          uint32    `json:"user_id"`
	ItemDescription string    `json:"item_description"`
	ItemQuantity    uint32    `json:"item_quantity"`
	ItemPrice       uint32    `json:"item_price"`
	TotalValue      uint32    `json:"total_value"`
	User			UserResponse
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// PostOrderRequest Get the request body of a Order POST.
type PostOrderRequest struct {
	UserID          uint32 `json:"user_id"`
	ItemDescription string `json:"item_description"`
	ItemQuantity    uint32 `json:"item_quantity"`
	ItemPrice       uint32 `json:"item_price"`
	TotalValue      uint32 `json:"total_value"`
}
