// Defines HTTP server functionality using Echo, a lightweight web framework for Go. It implements the Server interface, which consists of methods for starting the server, checking service readiness, and handling HTTP routes for CRUD operations on customers.

// Defines HTTP route handlers for customer-related operations such as retrieving all customers.

package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *EchoServer) GetAllCustomers(ctx echo.Context) error {
	emailAddress := ctx.QueryParam("emailAddress")

	customers, err := s.DB.GetAllCustomers(ctx.Request().Context(), emailAddress)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, customers)
}
