package db

import "errors"

type ProductPrice struct {
	ID int64
	ProductID int64
	Price float64
}

var ErrProductNotFound = errors.New("product not found")