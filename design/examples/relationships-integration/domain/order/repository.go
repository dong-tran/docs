package order

// OrderRepository - Repository interface (DDD pattern)
// Defined in domain layer but implemented in infrastructure (DIP)
type OrderRepository interface {
	Save(order *Order) error
	FindByID(id OrderID) (*Order, error)
	FindByCustomerID(customerID CustomerID) ([]*Order, error)
	Update(order *Order) error
}
