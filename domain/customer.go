package domain

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// acts as a port: any component implementing this interface is becoming adapter
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, error)
}
