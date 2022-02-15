package db

import "database/sql"

// Databases hold the database connection.
type Databases struct {
	*sql.DB
}
