package payment

type InMemoryStorage struct {
	payments map[int]Payment
}

func getNextID(payments map[int]Payment) int {
	max := -1
	for key := range payments {
		if key > max {
			max = key
		}
	}
	return max + 1
}

func (storage *InMemoryStorage) Create(payment Payment) Identifier {
	id := getNextID(storage.payments)
	storage.payments[id] = payment
	return Identifier(id)
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{payments: make(map[int]Payment)}
}
