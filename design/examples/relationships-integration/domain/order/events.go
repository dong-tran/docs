package order

// Domain Events (DDD pattern)
type OrderCreatedEvent struct {
	OrderID    string
	CustomerID string
	Total      float64
}

type OrderPaidEvent struct {
	OrderID       string
	PaymentMethod string
	Amount        float64
}

type OrderShippedEvent struct {
	OrderID        string
	TrackingNumber string
}
