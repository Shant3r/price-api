package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shant3r/price-api/internal/db"
)

type Handler struct {
	repo *db.Repository
}

func New(repository *db.Repository) *Handler {
	return &Handler{repo: repository}
}

func (h *Handler) AddProductPrice(ctx context.Context, c *gin.Context) {
	req := new(AddProductPriceRequest)
	err := c.BindJSON(req)
	if err != nil {
		internalError(c, err)
		return
	}
	if req.ProductID <= 0 {
		badRequest(c)
		return
	}
	if req.Price <= 0 {
		badRequest(c)
		return
	}
	err = h.repo.AddProductPrice(ctx, req.ProductID, req.Price)
	if err == db.ErrProductNotFound {
		badRequest(c)
		return
	}
}

func (h *Handler) GetProductPrice(ctx context.Context, c *gin.Context) {

}


func internalError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, fmt.Sprintf("internal error: %s", err))
}

func badRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, "bad request")
}