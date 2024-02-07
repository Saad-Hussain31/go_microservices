// Contains structs representing data models. In this case, Customer and Health structs are defined for representing customer data and service health status, respectively.

package models

type Customer struct {
	CustomerID string `gorm:"primaryKey" json:"customerId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"emailAddress"`
	Phone      string `json:"phoneNumber"`
	Address    string `json:"address"`
}
