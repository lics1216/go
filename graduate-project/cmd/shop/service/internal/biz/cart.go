package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type CartItem struct {
	ItemId   int64
	Quantity int64
}

type CartRepo interface {
	ListCartItem(context context.Context, userId int64, pageNum, pageCount int64) ([]CartItem, error)
	AddCartItem(context context.Context, userId int64, itemId int64, quantity int64) ([]CartItem, error)
}

type CartUseCase struct {
	repo CartRepo
	log  *log.Helper
}

func NewCartUseCase(repo CartRepo, logger log.Logger) *CartUseCase {
	return &CartUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "cartcase/interface")),
	}
}

func (c *CartUseCase) ListCartItem(context context.Context, userId, pageNum, pageCount int64) ([]CartItem, error) {
	return c.repo.ListCartItem(context, userId, pageNum, pageCount)
}

func (c *CartUseCase) AddCartItem(context context.Context, userId, itemId, quantity int64) ([]CartItem, error) {
	return c.repo.AddCartItem(context, userId, itemId, quantity)
}
