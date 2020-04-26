package repositories

import (
	"testing"
)

type EntityMock struct {
	id string
}

func (e EntityMock) Id() string {
	return e.id
}

func TestInMemoryRepository_Create(t *testing.T) {
	repository := NewInMemoryCreateRepository()
	e1 := EntityMock{"1"}

	// save first payment, no errors expected
	err := repository.Create(e1)
	if err != nil {
		t.Errorf("Unepxected error during payment creation: %s", err)
	}

	// try to save payment with the same id
	err = repository.Create(e1)
	expectedErr := EntityAlreadyExistsError{"1"}
	if err != expectedErr {
		t.Error("Payment with id already exists error expected")
	}
}

func TestInMemoryRepository_CreateBulk(t *testing.T) {
	e1 := EntityMock{"1"}
	e2 := EntityMock{"2"}
	repository := NewInMemoryCreateRepository()

	createdQuantity, err := repository.CreateBulk([]Entity{e1, e2})
	if err != nil {
		t.Errorf("Unpected error")
		if createdQuantity != 2 {
			t.Errorf("2 payments expected to be created")
		}
	}

	repository = NewInMemoryCreateRepository()
	expectedErr := EntityAlreadyExistsError{"1"}
	createdQuantity, err = repository.CreateBulk([]Entity{e1, e1})
	if err != expectedErr {
		t.Errorf("Payment alread exists error expected")
	}
	if createdQuantity != 1 {
		t.Errorf("One payment expected to be created before error")
	}
}
