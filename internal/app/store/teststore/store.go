package teststore

import (
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"
)

// Store ...
type Store struct {
	userRepository *UserRepository
}

// New store
func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}
