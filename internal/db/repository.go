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
	var res int64
	err := r.database.QueryRowContext(ctx, "SELECT id FROM product WHERE id = $1", productID).Scan(&res)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrProductNotFound
		}
		return err
	}

	_, err = r.database.ExecContext(ctx, `
	insert into product_price (product_id, price)
	values ($1, $2)
`, productID, price)
	if err != nil {
		return err
	}
	return nil
}
