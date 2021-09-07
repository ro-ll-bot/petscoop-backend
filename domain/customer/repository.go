package customer

type CustomerRepository interface {
	Get(id string) *Customer
	Delete(id string) error
	Save(Customer) error
	Update(Customer) error
}