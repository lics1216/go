package cache

import (
	"context"
	"encoding/json"
	"eshop/cmd/user/internal/data/db/ent"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
	"time"
)

type UserCache struct {
	cache *Data
	log   *log.Helper
}

func NewUserCache(data *Data, logger log.Logger) *UserCache {
	return &UserCache{
		cache: data,
		log:   log.NewHelper(log.With(logger, "module", "cache/user-service")),
	}
}

func (u2 *UserCache) CreateUser(ctx context.Context, u *ent.User) (*ent.User, error) {
	result := u2.cache.cache.WithContext(ctx).Set("user@"+strconv.FormatInt(u.ID, 10), u, 24*time.Hour)
	return u, result.Err()
}

func (u2 *UserCache) GetUser(ctx context.Context, id int64) (*ent.User, error) {
	u := ent.User{}
	result, err := u2.cache.cache.Get("user@" + strconv.FormatInt(id, 10)).Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(result, &u)
	if err != nil {
		return nil, err
	}
	return &u, err
}
