package storage

import (
	"fmt"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) String() string {
	return fmt.Sprintf("id: %d; name: %s; email: %s; password: %s;",
		u.ID, u.Name, u.Email, u.Password)
}

func (storage *Storage) GetUsers() ([]User, error) {
	const op = "users.GetUsers"
	var users []User

	tx, err := storage.DB.Begin()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback()

	q := `select * from Users`
	rows, err := tx.QueryContext(Context, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return users, nil
}

func (storage *Storage) InsertUser(user User) error {
	const op = "users.InsertUser"
	tx, err := storage.DB.Begin()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback()

	q := `insert into Users(id, name, email, password) values($1, $2, $3, $4)`
	_, err = tx.ExecContext(Context, q,
		user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (storage *Storage) UpdateUser(user User) error {
	const op = "users.UpdateUser"
	tx, err := storage.DB.Begin()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback()

	q := `update Users
	set name=$2, email=$3, password=$4
	where id=$1;`
	_, err = tx.ExecContext(Context, q,
		user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (storage *Storage) DeleteUser(userID string) error {
	const op = "users.DeleteUser"
	tx, err := storage.DB.Begin()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback()

	q := `delete from Users
	where id=$1;`
	_, err = tx.ExecContext(Context, q, userID)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (storage *Storage) GetUserStat(userID int) (float64, float64, error) {
	const op = "users.GetUserStat"
	tx, err := storage.DB.Begin()
	if err != nil {
		return -1, -1, fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback()

	q := `select sum(Orders.total_amount) as total_order_amount,
    avg(Products.price) as avg_product_price
	from Users
	join Orders on $1 = Orders.user_id
	join OrderProducts on Orders.id = OrderProducts.order_id
	join Products on OrderProducts.product_id = Products.id
	group by $1`
	rows, err := tx.QueryContext(Context, q, userID)
	if err != nil {
		return -1, -1, err
	}
	defer rows.Close()

	var sum, avg float64
	for rows.Next() {
		if err = rows.Scan(&sum, &avg); err != nil {
			return -1, -1, err
		}
	}
	if err = rows.Err(); err != nil {
		return -1, -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, -1, fmt.Errorf("%s: %w", op, err)
	}
	return sum, avg, err
}
