package service

import (
	"context"
	v1 "eshop/api/shop/service/v1"
)

func (s *ShopInterface) ListCartItem(ctx context.Context, req *v1.ListCartItemReq) (*v1.ListCartItemReply, error) {
	var item []*v1.ListCartItemReply_Item
	reply := &v1.ListCartItemReply{}
	result, err := s.cartcase.ListCartItem(ctx, req.UserId, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}
	for _, cartItem := range result {
		item = append(item, &v1.ListCartItemReply_Item{ItemId: cartItem.ItemId, Quantity: cartItem.Quantity})
	}
	reply.Items = item
	return reply, nil
}

func (s *ShopInterface) AddCartItem(ctx context.Context, req *v1.AddCartItemReq) (*v1.AddCartItemReply, error) {
	var item []*v1.AddCartItemReply_Item
	reply := &v1.AddCartItemReply{}
	result, err := s.cartcase.ListCartItem(ctx, req.UserId, req.ItemId, req.Quantity)
	if err != nil {
		return nil, err
	}
	for _, cartItem := range result {
		item = append(item, &v1.AddCartItemReply_Item{ItemId: cartItem.ItemId, Quantity: cartItem.Quantity})
	}
	reply.Items = item
	return reply, nil
}
