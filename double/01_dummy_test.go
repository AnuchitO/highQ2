package double

import "testing"

type DummySearcher struct{}

func (ds DummySearcher) Search(people []*Person, firstName, lastName string) *Person {
	return &Person{}
}

func TestFindReturnsError(t *testing.T) {
	phonebook := &Phonebook{}

	_, err := phonebook.Find(DummySearcher{}, "", "")

	want := ErrMissingArgs
	if err != want {
		t.Errorf("Want '%s', err '%s'", want, err)
	}
}
