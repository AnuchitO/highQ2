package sql

import (
	"database/sql"
	"testing"
)

type mockDB struct {
}

func (mockDB) LastInsertId() (int64, error) {
	return 1, nil
}
func (mockDB) RowsAffected() (int64, error) {
	return 32, nil
}

func (m *mockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return m, nil
}

func TestExecQuery(t *testing.T) {
	mock := &mockDB{}

	r, _ := execQuery(mock, "SELECT * FROM sql", nil)

	if r != 32 {
		t.Errorf("should return row effect %d but it got %d.", 32, r)
	}
}
