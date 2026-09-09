package vendingmachine

type Product struct {
	Name  string
	Price float64
}

type Slot struct {
	Code     string
	Product  *Product
	Quantity int
}
