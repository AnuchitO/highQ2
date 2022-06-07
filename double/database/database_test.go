package database

import "testing"

type fakeDB struct {
	Database
}

func (f fakeDB) Insert(collection string, data interface{}) error {
	return nil
}

func TestInsert(t *testing.T) {
	mock := &fakeDB{}

	err := Insert(mock, "product", `{}`)

	if err != nil {
		t.Error(err.Error())
	}
}
