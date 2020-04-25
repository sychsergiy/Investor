package ports

import "investor/entities"

type InMemoryStorage struct {
	payments map[int]entities.Payment
}

func getNextID(payments map[int]entities.Payment) int {
	max := -1
	for key := range payments {
		if key > max {
			max = key
		}
	}
	return max + 1
}

func (storage *InMemoryStorage) Create(payment entities.Payment) Identifier {
	id := getNextID(storage.payments)
	storage.payments[id] = payment
	return Identifier(id)
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{payments: make(map[int]entities.Payment)}
}
