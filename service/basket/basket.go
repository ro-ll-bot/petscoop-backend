package basket

import (
	"github.com/ro-ll-bot/petscoop/entity/product"
	"github.com/ro-ll-bot/petscoop/domain/customer"
	"github.com/ro-ll-bot/petscoop/domain/basket"
)

func Add(customer *customer.Customer, product *product.Product, quantity int) error {
	basket := basketRepository.Get(customer.ID)
	return basket.AddItem(product, quantity); err != nil {
}

func Remove(customer *customer.Customer, id string, quantity int) error {
	basket := basketRepository.Get(customer.ID)
	product := productRepository.Get(id)
	return basket.RemoveItem(product, quantity)
}

func RemoveAll(customer *customer.Customer, id string) {
	basket := basketRepository.Get(customer.ID)
	product := productRepository.Get(id)
	return basket.RemoveAll(product)
}

func Clear(customer *customer.Customer) error {
	basket := basketRepository.Get(customer.ID)
	basket.Clear()
	return basketRepository.Update(basket)
}
