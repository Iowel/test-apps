package cache

import (
	"context"

	"github.com/Iowel/test-apps/internal/store"

	"github.com/stretchr/testify/mock"
)

type MockUserStore struct {
	mock.Mock
}

func NewMockStore() Storage {
	return Storage{
		Users: &MockUserStore{},
	}
}

func (m MockUserStore) Get(ctx context.Context, userID int64) (*store.User, error) {
	args := m.Called(userID)

	return nil, args.Error(1)
}

func (m MockUserStore) Set(ctx context.Context, user *store.User) error {
	args := m.Called(user)

	return args.Error(0)
}
