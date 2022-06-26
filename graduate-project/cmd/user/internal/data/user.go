package data

import (
	"context"
	"eshop/cmd/user/internal/biz"
	"eshop/cmd/user/internal/data/cache"
	"eshop/cmd/user/internal/data/db"
	"eshop/cmd/user/internal/pkg/util"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	db    *db.UserRepo
	cache *cache.UserCache
	log   *log.Helper
}

func NewUserRepo(db *db.UserRepo, cache *cache.UserCache, logger log.Logger) biz.UserRepo {
	return &userRepo{
		db:    db,
		cache: cache,
		log:   log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (u2 *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	user, err := u2.db.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	} else {
		// todo : sending to kafka, to init cache write

		return &biz.User{Id: user.ID, Username: user.Username}, err
	}
}

func (u2 *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	u, err := u2.cache.GetUser(ctx, id)
	if u == nil || err != nil {
		u, err = u2.db.GetUser(ctx, id)
	}
	if u == nil || err != nil {
		return nil, err
	}
	_, _ = u2.cache.CreateUser(ctx, u)
	return &biz.User{Id: u.ID, Username: u.Username}, err
}

func (u2 *userRepo) VerifyPassword(ctx context.Context, u *biz.User) (bool, error) {
	user, err := u2.GetUser(ctx, u.Id)
	if user == nil || err != nil {
		return false, err
	}
	return util.CheckPasswordHash(u.Password, user.Password), nil

}
