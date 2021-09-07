package basket

type Basket struct {
	ID string

	Items map[string]*BasketItem

	RawPrice float64
	TotalPrice float64
}

type BasketItem struct {
	Product
	Quantity int
}

func NewBasketItem(product *Product, quantity uint) (*BasketItem, error) {
	if product == nil {
		return nil, errors.New("product can not be nil")
	}

	item := BasketItem{
		Product: product,
		Quantity: quantity,
	}

	return item
}

func (b *Basket) AddItem(product *Product, quantity int) error {
	if product == nil {
		return errors.New("product can not be nil")
	}

	if quantity <= 0{
		return errors.New("quantity should bigger than 0")
	}

	item := NewBasketItem(product, quantity)
	if ok, oldBasketItem := b.Items[item.ID]; ok {
		oldBasketItem.quantity += quantity
	} else {
		b.Items[item.ID] = item
	}

	return nil
}

