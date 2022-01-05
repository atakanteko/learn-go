package domain

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	var customer []Customer

	c1 := Customer{Id: "1001", Name: "Ata", City: "Istanbul", Zipcode: "11011", DateOfBirth: "2000-01-01", Status: "1"}
	c2 := Customer{Id: "1002", Name: "Erhan", City: "Kocaeli", Zipcode: "53222", DateOfBirth: "1998-01-01", Status: "0"}
	c3 := Customer{Id: "1003", Name: "Melike", City: "Antalya", Zipcode: "42211", DateOfBirth: "1999-01-01", Status: "1"}

	customer = append(customer, c1)
	customer = append(customer, c2)
	customer = append(customer, c3)

	return CustomerRepositoryStub{customers: customer}
}
