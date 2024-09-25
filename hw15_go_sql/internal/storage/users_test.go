package storage

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var (
	User1 = User{
		ID:       1,
		Name:     "test-user1",
		Email:    "test-email1@example.com",
		Password: "Test123",
	}

	User2 = User{
		ID:       2,
		Name:     "test-user2",
		Email:    "test-email2@example.com",
		Password: "Test12",
	}

	User3 = User{
		ID:       3,
		Name:     "test-user3",
		Email:    "test-email3@example.com",
		Password: "Test1",
	}
)

func TestGetUsers(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	var s Storage
	s.AddDB(db)

	assert.NoError(t, err)
	defer db.Close()

	existsRows := sqlmock.NewRows([]string{"id", "name", "email", "password"}).
		AddRow(User1.ID, User1.Name, User1.Email, User1.Password).
		AddRow(User2.ID, User2.Name, User2.Email, User2.Password).
		AddRow(User3.ID, User3.Name, User3.Email, User3.Password)

	mock.ExpectBegin()
	mock.ExpectQuery(`select * from Users`).
		WillReturnRows(existsRows)
	mock.ExpectCommit()

	var expectUsers []User
	expectUsers = append(expectUsers, User1)
	expectUsers = append(expectUsers, User2)
	expectUsers = append(expectUsers, User3)

	users, err := s.GetUsers()
	assert.Nil(t, err)
	assert.Equal(t, expectUsers, users)
}

func TestInsertUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	var s Storage
	s.AddDB(db)

	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(`insert into Users(id, name, email, password) values($1, $2, $3, $4)`).
		WithArgs(User1.ID, User1.Name, User1.Email, User1.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = s.InsertUser(User1)
	assert.Nil(t, err)
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	var s Storage
	s.AddDB(db)
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(`update Users
	set name=$2, email=$3, password=$4
	where id=$1;`).
		WithArgs(User2.ID, User2.Name, User2.Email, User2.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()
	err = s.UpdateUser(User2)
	assert.Nil(t, err)
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	var s Storage
	s.AddDB(db)
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(`delete from Users
	where id=$1;`).
		WithArgs("test-user2").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = s.DeleteUser("test-user2")
	assert.Nil(t, err)
}

func TestGetUserStat(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	var s Storage
	s.AddDB(db)

	assert.NoError(t, err)
	defer db.Close()

	existsRows := sqlmock.NewRows([]string{"sum", "avg"}).
		AddRow(25, 5)

	mock.ExpectBegin()
	mock.ExpectQuery(
		`select sum(Orders.total_amount) as total_order_amount,
		avg(Products.price) as avg_product_price
		from Users
		join Orders on $1 = Orders.user_id
		join OrderProducts on Orders.id = OrderProducts.order_id
		join Products on OrderProducts.product_id = Products.id
		group by $1`).
		WithArgs(2).
		WillReturnRows(existsRows)
	mock.ExpectCommit()

	var expectSum, expectAvg float64
	expectSum = 25
	expectAvg = 5

	sum, avg, err := s.GetUserStat(2)
	assert.Nil(t, err)
	assert.Equal(t, expectSum, sum)
	assert.Equal(t, expectAvg, avg)
}
