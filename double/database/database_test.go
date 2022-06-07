package database

import "testing"

type fakeDB struct{}

func (fakeDB) SetDBName(dbName string) {

}
func (fakeDB) DBName() string {
	return ""
}
func (fakeDB) Insert(collection string, data interface{}) error {
	return nil
}
func (fakeDB) Count(collection string, query interface{}) (int, error) {
	return 0, nil

}
func (fakeDB) FindOne(collection string, query interface{}, data interface{}) error {
	return nil
}
func (fakeDB) FindOnePrimary(collection string, query interface{}, data interface{}) error {
	return nil

}
func (fakeDB) Find(collection string, query interface{}, data interface{}) error {
	return nil

}
func (fakeDB) FindAndSelect(collection string, query interface{}, selector interface{}, data interface{}) error {
	return nil

}
func (fakeDB) FindLimit(collection string, query interface{}, limit int, data interface{}) error {
	return nil

}
func (fakeDB) FindLimitAndOrderBy(collection string, query interface{}, limit int, orderBy string, data interface{}) error {
	return nil

}
func (fakeDB) FindLimitAndOrderByM(collection string, query interface{}, limit int, orderBy []string, data interface{}) error {
	return nil

}
func (fakeDB) FindOrderByOne(collection string, query interface{}, orderby string, data interface{}) error {
	return nil

}
func (fakeDB) FindByOrder(collection string, query interface{}, orderBy []string, data interface{}) error {
	return nil

}
func (fakeDB) FindOrderAndUpdate(collection string, query interface{}, orderBy []string, changeQuery interface{}, afterItem interface{}) error {
	return nil

}
func (fakeDB) Upsert(collection string, query interface{}, data interface{}) error {
	return nil

}
func (fakeDB) Update(collection string, query interface{}, data interface{}) error {
	return nil

}
func (fakeDB) UpdateAll(collection string, query interface{}, data interface{}) error {
	return nil

}
func (fakeDB) Aggregate(collection string, query interface{}, data interface{}) error {
	return nil
}
func (fakeDB) RemoveAll(collection string, query interface{}) error {
	return nil
}
func TestInsert(t *testing.T) {
	mock := &fakeDB{}

	err := Insert(mock, "product", `{}`)

	if err != nil {
		t.Error(err.Error())
	}
}
