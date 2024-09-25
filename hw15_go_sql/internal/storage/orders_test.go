package storage

import (
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var (
	date   = time.Date(2024, time.May, 6, 5, 0, 0, 0, time.Local)
	Order1 = Order{
		ID:          1,
		UserID:      1,
		OrderDate:   date,
		TotalAmount: 5.12,
	}
	Order2 = Order{
		ID:          2,
		UserID:      3,
		OrderDate:   date,
		TotalAmount: 29,
	}
	Order3 = Order{
		ID:          3,
		UserID:      1,
		OrderDate:   date,
		TotalAmount: 2,
	}
	NewOrder1 = NewOrder{
		ID:        1,
		UserID:    1,
		ProductID: 3,
		Amount:    1,
	}
	Product4 = Product{
		ID:    3,
		Name:  "tomato",
		Price: 5.12,
	}
)

func TestGetOrders(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	var s Storage
	s.AddDB(db)

	assert.NoError(t, err)
	defer db.Close()

	existsRows := sqlmock.NewRows(
		[]string{"id", "user_id", "order_date", "total_amount"}).
		AddRow(Order1.ID, Order1.UserID, Order1.OrderDate, Order1.TotalAmount).
		AddRow(Order3.ID, Order3.UserID, Order3.OrderDate, Order3.TotalAmount)

	mock.ExpectBegin()
	mock.ExpectQuery(`select * from Orders
	where user_id=$1`).
		WithArgs(1).
		WillReturnRows(existsRows)
	mock.ExpectCommit()

	var expectOrders []Order
	expectOrders = append(expectOrders, Order1)
	expectOrders = append(expectOrders, Order3)

	orders, err := s.GetOrders(1)
	assert.Nil(t, err)
	assert.Equal(t, expectOrders, orders)
}

func TestDeleteOrder(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	var s Storage
	s.AddDB(db)
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(`delete from OrderProducts
	where order_id=$1;`).
		WithArgs(Order1.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`delete from Orders
	where id=$1;`).
		WithArgs(Order1.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = s.DeleteOrder(Order1.ID)
	assert.Nil(t, err)
}

func TestInsertOrder(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	var s Storage
	s.AddDB(db)

	assert.NoError(t, err)
	defer db.Close()

	priceRows := sqlmock.NewRows([]string{"price"}).
		AddRow(Product2.Price)

	date = time.Now()
	mock.ExpectBegin()
	mock.ExpectQuery(`select price from Products
	where id=$1`).
		WithArgs(Product4.ID).
		WillReturnRows(priceRows)
	mock.ExpectExec(`insert into Orders (id, user_id, order_date, total_amount)
	values($1, $2, $3, $4)`).
		WithArgs(NewOrder1.ID,
			NewOrder1.UserID,
			date.Format("2006-01-02 15:04:05"),
			fmt.Sprintf("%.2f", (float64(NewOrder1.Amount)*Product2.Price))).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`insert into OrderProducts (order_id, product_id, amount)
		values($1, $2, $3)`).
		WithArgs(NewOrder1.ID,
			NewOrder1.ProductID,
			NewOrder1.Amount).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = s.InsertOrder(NewOrder1)
	assert.Nil(t, err)
}
