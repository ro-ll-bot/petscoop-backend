package entity

type Product struct {
	ID string // maybe just barcode for that
	Barcode string
	Name string
	ProductDetails
}

type ProductDetails struct {
}
