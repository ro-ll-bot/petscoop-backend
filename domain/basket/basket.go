package basket

type Basket struct {
	ID string

	Products []*entity.Product

	RawPrice float64
	TotalPrice float64
}