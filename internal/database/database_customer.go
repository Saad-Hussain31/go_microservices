// Contains the implementation of database operations specific to customers, like retrieving all customers based on email address.

package database

import (
	"context"

	"github.com/Saad-Hussain31/go-microservices/internal/models"
)

func (c Client) GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error) {
	var customers []models.Customer
	result := c.DB.WithContext(ctx).
		Where(models.Customer{Email: emailAddress}).
		Find(&customers)
	return customers, result.Error
}
