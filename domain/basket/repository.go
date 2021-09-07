package basket

type BasketRepository struct {
	Get(id string) *Basket
	Save(Basket) error
	Update(Basket) error
}