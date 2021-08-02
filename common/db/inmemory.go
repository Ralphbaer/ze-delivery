package db

// InMemoryDB implements a in-memory database keeping items in a map[string]interface.
type InMemoryDB struct {
	data map[string]interface{}
}

// NewInMemoryDB creates a new inMemoryDB instance. initialData is optional.
func NewInMemoryDB(initialData map[string]interface{}) *InMemoryDB {
	if initialData != nil {
		return &InMemoryDB{initialData}
	}
	return &InMemoryDB{make(map[string]interface{})}
}

// Save stores any value inside InMemoryRepo
func (r *InMemoryDB) Save(id string, value interface{}) error {
	r.data[id] = value
	return nil
}

// Find finds an entry containing given id inside InMemoryRepo
func (r *InMemoryDB) Find(id string) (interface{}, error) {
	v, ok := r.data[id]
	if !ok {
		return nil, ErrRecordNotFound
	}
	return v, nil
}

// FindAll returns all entries.
func (r *InMemoryDB) FindAll() map[string]interface{} {
	return r.data
}

// Delete delete a specific entry given an id.
func (r *InMemoryDB) Delete(id string) {
	delete(r.data, id)
}

// Count returns the quantity of entries at InMemoryRepo
func (r *InMemoryDB) Count() int {
	if len(r.data) == 0 {
		return 0
	}

	return len(r.data)
}
