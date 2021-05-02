package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "Bob", "New Yourk", "192256", "1980-01-01", "active"},
		{"2", "Alice", "Denver", "102216", "1982-02-02", "active"},
	}
	return CustomerRepositoryStub{customers}
}
