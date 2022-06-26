package db

import (
	"context"
	"eshop/cmd/user/internal/biz"
	"eshop/cmd/user/internal/data/db/ent"
	"eshop/cmd/user/internal/pkg/util"
	"github.com/go-kratos/kratos/v2/log"
)

// var _ biz.UserRepo = (*UserRepo)(nil)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) *UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, u *biz.User) (*ent.User, error) {
	ph, err := util.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	po, err := r.data.db.User.
		Create().
		SetUsername(u.Username).
		SetPasswordHash(ph).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return po, err
	//	return &biz.User{Id: po.ID, Username: po.Username}, err
}

func (r *UserRepo) GetUser(ctx context.Context, id int64) (*ent.User, error) {
	return r.data.db.User.Get(ctx, id)
	//if err != nil {
	//	return nil, err
	//}
	//return &biz.User{Id: po.ID, Username: po.Username}, err
}
