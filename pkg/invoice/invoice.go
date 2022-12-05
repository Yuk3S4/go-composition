package invoice

import "github.com/Yuk3S4/go-composition/pkg/customer"

type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
}
