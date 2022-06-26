package service

import (
	"context"
	"eshop/api/shop/service/v1"
	"eshop/cmd/shop/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"golang.org/x/sync/errgroup"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewShopInterface)

type ShopInterface struct {
	v1.UnimplementedShopInterfaceServer

	uc       *biz.UserUseCase
	cc       *biz.CatalogUseCase
	cartcase *biz.CartUseCase

	log *log.Helper
}

func NewShopInterface(uc *biz.UserUseCase, cc *biz.CatalogUseCase, logger log.Logger) *ShopInterface {
	return &ShopInterface{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		uc:  uc,
		cc:  cc,
	}
}

func (s *ShopInterface) Index(context context.Context, req *v1.IndexReq) (*v1.IndexReply, error) {
	group, ctx := errgroup.WithContext(context)
	reply := &v1.IndexReply{}
	group.Go(func() error {
		user, err := s.uc.GetUser(ctx, req.UserId)
		if err != nil {
			return err
		}
		reply.User = &v1.User{Id: user.Id, Username: user.Username}
		return nil
	})

	group.Go(func() error {
		beers, err := s.ListBeer(ctx, &v1.ListBeerReq{PageNum: 0, PageSize: 10})
		if err != nil {
			return err
		}
		reply.Beer = beers
		return nil
	})

	group.Go(func() error {
		carts, err := s.ListCartItem(ctx, &v1.ListCartItemReq{PageSize: 10, PageNum: 0})
		if err != nil {
			return err
		}
		reply.Cart = carts
		return nil
	})

	return reply, group.Wait()
}
