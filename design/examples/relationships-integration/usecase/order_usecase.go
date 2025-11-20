package usecase

import (
"github.com/dong-tran/docs/integration-example/domain/order"
"github.com/dong-tran/docs/integration-example/shared/patterns"
)

// OrderUseCase - Application Service (Clean Architecture Use Case Layer)
// Orchestrates domain objects and publishes events
type OrderUseCase struct {
	orderRepo      order.OrderRepository
	paymentFactory *patterns.PaymentFactory
	eventPublisher *patterns.EventPublisher
}

func NewOrderUseCase(
orderRepo order.OrderRepository,
paymentFactory *patterns.PaymentFactory,
eventPublisher *patterns.EventPublisher,
) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:      orderRepo,
		paymentFactory: paymentFactory,
		eventPublisher: eventPublisher,
	}
}

// CreateOrderDTO - Input DTO
type CreateOrderDTO struct {
	CustomerID string
	Items      []OrderItemDTO
}

type OrderItemDTO struct {
	ProductID   string
	ProductName string
	Quantity    int
	Price       float64
	Currency    string
}

// CreateOrder - Use case method
func (uc *OrderUseCase) CreateOrder(dto CreateOrderDTO) (*order.Order, error) {
	// Convert DTOs to domain objects
	customerID := order.NewCustomerID(dto.CustomerID)
	
	items := make([]order.OrderItem, 0, len(dto.Items))
	for _, itemDTO := range dto.Items {
		price, err := order.NewMoney(itemDTO.Price, itemDTO.Currency)
		if err != nil {
			return nil, err
		}
		
		item, err := order.NewOrderItem(
itemDTO.ProductID,
itemDTO.ProductName,
itemDTO.Quantity,
price,
)
		if err != nil {
			return nil, err
		}
		items = append(items, *item)
	}

	// Create order using domain logic
	newOrder, err := order.NewOrder(customerID, items)
	if err != nil {
		return nil, err
	}

	// Persist
	if err := uc.orderRepo.Save(newOrder); err != nil {
		return nil, err
	}

	// Publish domain event
	uc.eventPublisher.Publish(patterns.Event{
Type: "OrderCreated",
Data: order.OrderCreatedEvent{
OrderID:    newOrder.ID().String(),
			CustomerID: newOrder.CustomerID().String(),
			Total:      newOrder.TotalAmount().Amount(),
		},
	})

	return newOrder, nil
}

// ProcessPayment - Use case using Strategy pattern
func (uc *OrderUseCase) ProcessPayment(orderID string, paymentMethod string) error {
	// Get order
	ord, err := uc.orderRepo.FindByID(order.OrderID{})
	if err != nil {
		return err
	}

	// Use Factory to create payment strategy (Factory + Strategy patterns)
	paymentStrategy, err := uc.paymentFactory.CreatePayment(paymentMethod)
	if err != nil {
		return err
	}

	// Process payment using strategy
	if err := paymentStrategy.ProcessPayment(ord.TotalAmount().Amount(), ord.ID().String()); err != nil {
		return err
	}

	// Update order status (domain logic)
	if err := ord.MarkAsPaid(); err != nil {
		return err
	}

	// Persist changes
	if err := uc.orderRepo.Update(ord); err != nil {
		return err
	}

	// Publish event
	uc.eventPublisher.Publish(patterns.Event{
Type: "OrderPaid",
Data: order.OrderPaidEvent{
OrderID:       ord.ID().String(),
			PaymentMethod: paymentStrategy.GetName(),
			Amount:        ord.TotalAmount().Amount(),
		},
	})

	return nil
}

// GetOrder - Query use case
func (uc *OrderUseCase) GetOrder(orderID string) (*order.Order, error) {
	return uc.orderRepo.FindByID(order.OrderID{})
}

// GetCustomerOrders - Query use case
func (uc *OrderUseCase) GetCustomerOrders(customerID string) ([]*order.Order, error) {
	return uc.orderRepo.FindByCustomerID(order.NewCustomerID(customerID))
}

// ShipOrder - Use case
func (uc *OrderUseCase) ShipOrder(orderID string, trackingNumber string) error {
	ord, err := uc.orderRepo.FindByID(order.OrderID{})
	if err != nil {
		return err
	}

	if err := ord.Ship(); err != nil {
		return err
	}

	if err := uc.orderRepo.Update(ord); err != nil {
		return err
	}

	uc.eventPublisher.Publish(patterns.Event{
Type: "OrderShipped",
Data: order.OrderShippedEvent{
OrderID:        ord.ID().String(),
			TrackingNumber: trackingNumber,
		},
	})

	return nil
}
