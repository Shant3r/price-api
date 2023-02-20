package api

type AddProductPriceRequest struct {
	ProductID int64   `json:"productID"`
	Price     float64 `json:"price"`
}
