package utils

import "github.com/DATA-DOG/go-sqlmock"

func NewUserQuery() string {
	return "SELECT * FROM \"users\"  WHERE (phone_number = $1) ORDER BY \"users\".\"id\" ASC LIMIT 1"
}

func NewUserTableRow() *sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "name", "phone_number", "email"})
}
