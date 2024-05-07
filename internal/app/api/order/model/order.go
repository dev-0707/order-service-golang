package model

// CreateOrderDto defines model for CreateOrderDto.
type Order struct {
	Id         string      `json:"id,omitempty"`
	CustomerId string      `json:"customerId,omitempty"`
	Lines      []OrderLine `json:"lines,omitempty"`
}

// OrderLine defines model for OrderLine.
type OrderLine struct {
	ProductId string `json:"productId,omitempty"`
	Quantity  int64  `json:"quantity,omitempty"`
}
