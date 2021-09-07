package basket

import (
	"github.com/ro-ll-bot/petscoop/domain/customer"
	"github.com/ro-ll-bot/petscoop/domain/product"
)

func Add(customer *customer.Customer, product *product.Product, quantity int) error {
	return nil
}

func Remove(customer *customer.Customer, id string, quantity int) error {
	return nil
}

func Clear(customer *customer.Customer) error {
	return nil
}