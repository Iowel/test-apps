package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Iowel/test-apps/internal/store"

	"github.com/go-redis/redis/v8"
)

type UserStore struct {
	redisDB *redis.Client
}

const UserExpDuration = time.Minute

func (u *UserStore) Get(ctx context.Context, userID int64) (*store.User, error) {
	cacheKey := fmt.Sprintf("user-%v", userID)

	data, err := u.redisDB.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var user store.User
	if data != "" {
		err := json.Unmarshal([]byte(data), &user)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (u *UserStore) Set(ctx context.Context, user *store.User) error {
	cacheKey := fmt.Sprintf("user-%v", user.ID)

	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return u.redisDB.SetEX(ctx, cacheKey, data, UserExpDuration).Err()
}
