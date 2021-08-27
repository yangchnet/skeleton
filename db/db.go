package db

import "database/sql"

type Databases struct {
	*sql.DB
}
