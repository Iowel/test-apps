package store

import (
	"context"
	"database/sql"
	"time"

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

func (m *MockUserStore) Create(ctx context.Context, tx *sql.Tx, u *User) error {
	return nil
}

func (m *MockUserStore) GetByID(ctx context.Context, userID int64) (*User, error) {
	return &User{ID: 1}, nil
}

func (m *MockUserStore) CreateAndInvite(ctx context.Context, user *User, token string, invitationExp time.Duration) error {
	return nil
}

func (m *MockUserStore) Activate(ctx context.Context, t string) error {
	return nil
}

func (m *MockUserStore) Delete(ctx context.Context, userID int64) error {
	return nil
}

func (m *MockUserStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	return &User{ID: 1}, nil
}
