package customer

// Estructura de cliente
type Customer struct {
	name    string
	address string
	phone   string
}

// New retorna un nuevo cliente (Customer)
func New(name, address, phone string) Customer {
	return Customer{name, address, phone}
}
