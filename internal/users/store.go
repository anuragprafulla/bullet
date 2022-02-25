package users

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type IUserStore interface {
	Get(ctx context.Context, in *GetUserRequest) (*User, error)
	List(ctx context.Context, in *ListUserRequest) ([]*User, error)
	Create(ctx context.Context, in *CreateUserRequest) error
	Update(ctx context.Context, in *UpdateUserRequest) error
	Delete(ctx context.Context, in *DeleteUserRequest) error
}

func init() {
	rand.Seed(time.Now().UTC().Unix())
}

func GenerateUniqueID() string {
	word := []byte("0987654321")
	rand.Shuffle(len(word), func(i, j int) {
		word[i], word[j] = word[j], word[i]
	})
	now := time.Now().UTC()
	return fmt.Sprintf("%010v-%010v-%s", now.Unix(), now.Nanosecond(), string(word))
}
