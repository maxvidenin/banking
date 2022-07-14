package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
} 
 
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: 1234, Name: "John Doe", City: "New York", Zipcode: "10001", DateOfBirth: "01/02/1989", Status: "active"},
		{Id: 1235, Name: "Jane Doe", City: "New York", Zipcode: "10001", DateOfBirth: "14/01/1989", Status: "active"},
		{Id: 1236, Name: "John Smith", City: "San Francisco", Zipcode: "20002", DateOfBirth: "11/07/1978", Status: "active"},
	}
	return CustomerRepositoryStub{customers}
}
