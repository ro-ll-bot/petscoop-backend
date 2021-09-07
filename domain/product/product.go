package product

import (
	"github.com/ro-ll-bot/petscoop/entity"
	"github.com/ro-ll-bot/petscoop/domain/seller"
)

type Product struct {
	entity.Product
	seller *seller.Seller
}

func NewProduct(seller *seller.Seller) *Product {
	return nil
}