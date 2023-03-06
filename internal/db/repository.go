package db

import (
	"context"
	"database/sql"
)

type Repository struct {
	database *sql.DB
}

func New(database *sql.DB) *Repository {
	return &Repository{
		database: database,
	}
}

func (r *Repository) AddProductPrice(ctx context.Context, productID int64, price float64) error {
	_, err := r.database.ExecContext(ctx, `
	insert into product_price (product_id, price)
	values ($1, $2)
`, productID, price)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetProductPrice(ctx context.Context, id int64) (float64, error) {
	var result float64
	err := r.database.QueryRowContext(ctx, "SELECT price FROM product_price WHERE product_id = $1", id).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, ErrProductNotFound
		}
		return 0, err
	}
	return result, nil
}
