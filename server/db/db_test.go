package db

import "testing"

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "virtual_machine",
		User:       "postgres",
		Password:   "postgres",
		Host:       "localhost",
		DisableSSL: true,
	}
	if conn.ConnectionURL() != "postgres://postgres:postgres@localhost/virtual_machine?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
