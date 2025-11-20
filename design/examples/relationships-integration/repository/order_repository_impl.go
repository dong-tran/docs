package repository

import (
"database/sql"
"encoding/json"

"github.com/dong-tran/docs/integration-example/domain/order"
"github.com/jmoiron/sqlx"
)

// OrderRepositoryImpl - Infrastructure implementation (Clean Architecture + DIP)
type OrderRepositoryImpl struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) order.OrderRepository {
	return &OrderRepositoryImpl{db: db}
}

type orderDB struct {
	ID          string  `db:"id"`
	CustomerID  string  `db:"customer_id"`
	Items       string  `db:"items"`
	TotalAmount float64 `db:"total_amount"`
	Currency    string  `db:"currency"`
	Status      string  `db:"status"`
	CreatedAt   string  `db:"created_at"`
	UpdatedAt   string  `db:"updated_at"`
}

func (r *OrderRepositoryImpl) Save(ord *order.Order) error {
	itemsJSON, _ := json.Marshal(ord.Items())
	
	query := `
		INSERT INTO orders (id, customer_id, items, total_amount, currency, status, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query,
ord.ID().String(),
		ord.CustomerID().String(),
		itemsJSON,
		ord.TotalAmount().Amount(),
		ord.TotalAmount().Currency(),
		string(ord.Status()),
		ord.CreatedAt(),
		ord.UpdatedAt(),
	)
	return err
}

func (r *OrderRepositoryImpl) FindByID(id order.OrderID) (*order.Order, error) {
	// Implementation details...
	return nil, sql.ErrNoRows
}

func (r *OrderRepositoryImpl) FindByCustomerID(customerID order.CustomerID) ([]*order.Order, error) {
	// Implementation details...
	return nil, nil
}

func (r *OrderRepositoryImpl) Update(ord *order.Order) error {
	query := `
		UPDATE orders
		SET status = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := r.db.Exec(query, string(ord.Status()), ord.UpdatedAt(), ord.ID().String())
	return err
}
