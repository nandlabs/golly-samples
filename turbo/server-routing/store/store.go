package store

import "github.com/nandlabs/golly-samples/turbo/server-routing/models"

type Store struct {
	data map[string]models.Item
}

func NewStore() *Store {
	return &Store{data: make(map[string]models.Item)}
}

func (s *Store) GetAll() []models.Item {
	items := []models.Item{}
	for _, item := range s.data {
		items = append(items, item)
	}
	return items
}

func (s *Store) GetById(id string) (item models.Item, exists bool) {
	item, exists = s.data[id]
	return
}

func (s *Store) Put(id string, item models.Item) {
	s.data[id] = item
}

func (s *Store) Delete(id string) {
	delete(s.data, id)
}
