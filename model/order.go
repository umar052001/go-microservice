package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderID    uint64     `json:"order_id"`
	CustomerID uuid.UUID  `json:"customer_id"`
	LineItem   []LineItem `json:"line_item"`
	CreatedAt  *time.Time `json:"created_at"`
	ShippedAt  *time.Time `json:"shipped_at"`
	CompltedAt *time.Time `json:"complted_at"`
}

type LineItem struct {
	ItemID   uuid.UUID `json:"item_id"`
	Quantity uint      `json:"quantity"`
	Price    uint      `json:"price"`
}
