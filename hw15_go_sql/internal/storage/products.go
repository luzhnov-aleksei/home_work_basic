package storage

import (
	"fmt"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *Product) String() string {
	return fmt.Sprintf("id: %d; name: %s; price: %v;",
		p.ID, p.Name, p.Price)
}

func (storage *Storage) InsertProduct(product Product) error {
	const op = "products.InsertProduct"

	tx, err := storage.DB.Begin()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback()

	q := `insert into Products(id, name, price) values($1, $2, $3)`
	_, err = tx.ExecContext(Context, q,
		product.ID, product.Name, product.Price)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (storage *Storage) UpdateProduct(product Product) error {
	const op = "products.UpdateProduct"
	tx, err := storage.DB.Begin()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer tx.Rollback()
	q := `update Products
	set name=$2, price=$3
	where id=$1;`
	_, err = tx.ExecContext(Context, q,
		product.ID, product.Name, product.Price)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (storage *Storage) DeleteProduct(productID int) error {
	const op = "products.DeleteProduct"

	tx, err := storage.DB.Begin()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback()

	q := `delete from OrderProducts
	where product_id=$1;`
	_, err = tx.ExecContext(Context, q, productID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	q = `delete from Products
	where id=$1;`
	_, err = tx.ExecContext(Context, q, productID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (storage *Storage) GetProductPrice(productID int) (float64, error) {
	const op = "products.GetProductPrice"
	var price float64
	q := `select price from Products
	where id=$1`

	rows, err := storage.DB.QueryContext(Context, q, productID)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&price); err != nil {
			return -1, fmt.Errorf("%s: %w", op, err)
		}
	}
	if price <= 0 {
		return -1, fmt.Errorf("%s: Invalid value of the price field", op)
	}
	return price, nil
}

func (storage *Storage) GetProducts() ([]Product, error) {
	const op = "products.GetProducts"
	tx, err := storage.DB.Begin()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback()

	var products []Product

	q := `select * from Products`

	rows, err := tx.QueryContext(Context, q)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()
	for rows.Next() {
		var product Product
		if err = rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return products, nil
}
