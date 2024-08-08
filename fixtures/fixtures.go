//go:build fixtures

package fixtures

import "database/sql"

func Truncate(db *sql.DB) {}

func NoData(db *sql.DB) {}
