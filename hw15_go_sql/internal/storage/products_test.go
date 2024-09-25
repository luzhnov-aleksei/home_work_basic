package storage

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var (
	Product1 = Product{
		ID:    1,
		Name:  "milk",
		Price: 14.5,
	}

	Product2 = Product{
		ID:    2,
		Name:  "bread",
		Price: 2,
	}

	Product3 = Product{
		ID:    3,
		Name:  "tomato",
		Price: 5.12,
	}
)

func TestInsertProduct(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	var s Storage
	s.AddDB(db)

	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(`insert into Products(id, name, price) values($1, $2, $3)`).
		WithArgs(Product1.ID, Product1.Name, Product1.Price).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = s.InsertProduct(Product1)
	assert.Nil(t, err)
}

func TestGetProducts(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	var s Storage
	s.AddDB(db)

	assert.NoError(t, err)
	defer db.Close()

	existsRows := sqlmock.NewRows([]string{"id", "name", "price"}).
		AddRow(Product1.ID, Product1.Name, Product1.Price).
		AddRow(Product2.ID, Product2.Name, Product2.Price).
		AddRow(Product3.ID, Product3.Name, Product3.Price)

	mock.ExpectBegin()
	mock.ExpectQuery(`select * from Products`).
		WillReturnRows(existsRows)
	mock.ExpectCommit()

	var expectProducts []Product
	expectProducts = append(expectProducts, Product1)
	expectProducts = append(expectProducts, Product2)
	expectProducts = append(expectProducts, Product3)

	products, err := s.GetProducts()
	assert.Nil(t, err)
	assert.Equal(t, expectProducts, products)
}

func TestGetProductPrice(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	var s Storage
	s.AddDB(db)

	assert.NoError(t, err)
	defer db.Close()

	existsRows := sqlmock.NewRows([]string{"price"}).
		AddRow(Product2.Price)

	mock.ExpectQuery(`select price from Products
	where id=$1`).
		WillReturnRows(existsRows)

	price, err := s.GetProductPrice(2)
	assert.Nil(t, err)
	assert.Equal(t, Product2.Price, price)
}

func TestUpdateProduct(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	var s Storage
	s.AddDB(db)
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(`update Products
	set name=$2, price=$3
	where id=$1;`).
		WithArgs(Product2.ID, Product2.Name, Product2.Price).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = s.UpdateProduct(Product2)
	assert.Nil(t, err)
}

func TestDeleteProduct(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	var s Storage
	s.AddDB(db)
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(`delete from OrderProducts
	where product_id=$1;`).
		WithArgs(Product1.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`delete from Products
	where id=$1;`).
		WithArgs(Product1.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = s.DeleteProduct(Product1.ID)
	assert.Nil(t, err)
}
