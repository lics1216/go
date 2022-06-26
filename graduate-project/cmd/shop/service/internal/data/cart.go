package data

import (
	"context"
	csV1 "eshop/api/cart/v1"
	"eshop/cmd/shop/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.CartRepo = (*cartRepo)(nil)

type cartRepo struct {
	data *Data
	log  *log.Helper
}

func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
	return &cartRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}

func (c *cartRepo) ListCartItem(context context.Context, userId int64, pageNum, pageCount int) ([]biz.CartItem, error) {
	cart, err := c.data.cc.GetCart(context, &csV1.GetCartReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	var result []biz.CartItem
	for _, item := range cart.Items {
		result = append(result, biz.CartItem{ItemId: item.ItemId, Quantity: item.Quantity})
	}
	return result, nil
}

func (c cartRepo) AddCartItem(context context.Context, userId int64, itemId int64, quantity int64) ([]biz.CartItem, error) {
	cart, err := c.data.cc.AddItem(context, &csV1.AddItemReq{UserId: userId, ItemId: itemId, Quantity: quantity})
	if err != nil {
		return nil, err
	}

	var result []biz.CartItem
	for _, item := range cart.Items {
		result = append(result, biz.CartItem{ItemId: item.ItemId, Quantity: item.Quantity})
	}
	return result, nil
}
