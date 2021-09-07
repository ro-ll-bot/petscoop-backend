package customer

import (
	"github.com/ro-ll-bot/petscoop/entity"
)

type Customer struct {
	*entity.User	

	basket *entity.Basket
}

func NewCustomer() *Customer {
	return nil
}