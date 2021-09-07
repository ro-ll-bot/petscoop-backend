package product

type ProductRepository interface {
	Save(product *Product) error
	Update(product *Product) error
	Delete(product *Product) error
	Get(id string) *Product
	Paginate(pageIndex, pageSize int) ([]Product, error)
}

