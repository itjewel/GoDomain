package persistence

import (
	"errors"
	"sync"

	"github.com/jewelmia/GoDomain/internal/domain/user"
)

type UserRepoInMemory struct {
	mu    sync.RWMutex
	store map[string]*user.User
}

func NewUserRepoInMemory() user.UserRepository {
	return &UserRepoInMemory{
		store: make(map[string]*user.User),
	}
}

func (r *UserRepoInMemory) GetByID(id string) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.store[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return u, nil
}

func (r *UserRepoInMemory) Save(u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.store[u.ID] = u
	return nil
}
