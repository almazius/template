package repo

import "database/sql"

type AuthPGRepo struct {
	db *sql.DB
}
