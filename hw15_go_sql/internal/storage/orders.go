package storage

import (
	"fmt"
	"time"
)

type Order struct {
	ID          int       `json:"id"`
	UserID      int       `json:"userId"`
	OrderDate   time.Time `json:"orderDate,omitempty"`
	TotalAmount float64   `json:"totalAmount"`
}

type NewOrder struct {
	ID        int `json:"id"`
	UserID    int `json:"userId"`
	ProductID int `json:"productId"`
	Amount    int `json:"amount"`
}

func (o *Order) String() string {
	return fmt.Sprintf(
		"id: %d; user id: %d; order date: %s; total amount: %v;",
		o.ID, o.UserID, o.OrderDate.Format("2006-01-02"), o.TotalAmount)
}

func (storage *Storage) InsertOrder(order NewOrder) error {
	const op = "orders.InsertOrder"
	tx, err := storage.DB.Begin()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback()

	price, err := storage.GetProductPrice(order.ProductID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	date := time.Now()

	q := `insert into Orders (id, user_id, order_date, total_amount)
	values($1, $2, $3, $4)`
	_, err = tx.ExecContext(Context, q,
		order.ID, order.UserID, date.Format("2006-01-02 15:04:05"),
		fmt.Sprintf("%.2f", (float64(order.Amount)*price)))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	q = `insert into OrderProducts (order_id, product_id, amount)
	values($1, $2, $3)`
	_, err = tx.ExecContext(Context, q,
		order.ID, order.ProductID, order.Amount)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (storage *Storage) DeleteOrder(orderID int) error {
	const op = "orders.DeleteOrder"
	tx, err := storage.DB.Begin()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback()

	q := `delete from OrderProducts
	where order_id=$1;`
	_, err = tx.ExecContext(Context, q, orderID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	q = `delete from Orders
	where id=$1;`
	_, err = tx.ExecContext(Context, q, orderID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (storage *Storage) GetOrders(userID int) ([]Order, error) {
	const op = "orders.GetOrders"

	tx, err := storage.DB.Begin()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback()

	var orders []Order

	q := `select * from Orders
	where user_id=$1`

	rows, err := tx.QueryContext(Context, q, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		if err = rows.Scan(&order.ID, &order.UserID, &order.OrderDate, &order.TotalAmount); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return orders, nil
}
