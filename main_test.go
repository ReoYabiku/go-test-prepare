package main

import (
	"database/sql"
	"go-test/fixtures"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	db *sql.DB
)

func TestMain(m *testing.M) {
	db = sql.OpenDB(nil)
	defer db.Close()
	m.Run()
}

func TestAdd(t *testing.T) {
	tests := map[string]struct {
		prepare func(db *sql.DB)
		a       int
		b       int
		check   func(t *testing.T, got int, err error)
	}{
		"ok": {
			prepare: fixtures.NoData,
			a:       10,
			b:       10,
			check: func(t *testing.T, got int, err error) {
				assert.NoError(t, err)
				assert.Equal(t, got, 20)
			},
		},
	}

	for name, tt := range tests {
		fixtures.Truncate(db)
		tt.prepare(db)

		t.Run(name, func(t *testing.T) {
			got, err := Add(tt.a, tt.b)
			tt.check(t, got, err)
		})
	}
}
