package payment

type InMemoryStorage struct {
	payments map[int]Record
}

func getNextID(payments map[int]Record) int {
	max := -1
	for key := range payments {
		if key > max {
			max = key
		}
	}
	return max + 1
}

func (storage *InMemoryStorage) Create(payment Record) Identifier {
	id := getNextID(storage.payments)
	storage.payments[id] = payment
	return Identifier(id)
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{payments: make(map[int]Record)}
}
