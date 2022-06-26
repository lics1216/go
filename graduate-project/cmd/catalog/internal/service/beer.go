package service

import (
	"context"
	"fmt"

	v1a "eshop/api/catalog/admin/v1"
	"eshop/api/catalog/v1"
	"eshop/cmd/catalog/internal/biz"
)

func (s *CatalogService) CreateBeer(ctx context.Context, req *v1a.CreateBeerReq) (*v1a.CreateBeerReply, error) {
	b := &biz.Beer{
		Name:        req.Name,
		Description: req.Description,
		Count:       req.Count,
		Images:      make([]biz.Image, 0),
	}
	for _, x := range req.Image {
		b.Images = append(b.Images, biz.Image{URL: x.Url})
	}
	x, err := s.bc.Create(ctx, b)
	img := make([]*v1a.CreateBeerReply_Image, 0)
	for _, i := range x.Images {
		img = append(img, &v1a.CreateBeerReply_Image{Url: i.URL})
	}
	return &v1a.CreateBeerReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       img,
	}, err
}

func (s *CatalogService) DeleteBear(ctx context.Context, req *v1a.DeleteBeerReq) (*v1a.DeleteBeerReq, error) {

	return nil, nil
}

func (s *CatalogService) GetBeer(ctx context.Context, req *v1.GetBeerReq) (*v1.GetBeerReply, error) {
	x, err := s.bc.Get(ctx, req.Id)
	if x == nil {
		return nil, nil
	}
	fmt.Printf("%v\n", x)
	img := make([]*v1.GetBeerReply_Image, 0)
	for _, i := range x.Images {
		img = append(img, &v1.GetBeerReply_Image{Url: i.URL})
	}
	return &v1.GetBeerReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       img,
	}, err
}

func (s *CatalogService) UpdateBeer(ctx context.Context, req *v1a.UpdateBeerReq) (*v1a.UpdateBeerReply, error) {
	b := &biz.Beer{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Count:       req.Count,
		Images:      make([]biz.Image, 0),
	}
	for _, x := range req.Image {
		b.Images = append(b.Images, biz.Image{URL: x.Url})
	}
	x, err := s.bc.Update(ctx, b)
	img := make([]*v1a.UpdateBeerReply_Image, 0)
	for _, i := range x.Images {
		img = append(img, &v1a.UpdateBeerReply_Image{Url: i.URL})
	}
	return &v1a.UpdateBeerReply{
		Id:          x.Id,
		Name:        x.Name,
		Description: x.Description,
		Count:       x.Count,
		Image:       img,
	}, err
}

func (s *CatalogService) ListBeer(ctx context.Context, req *v1.ListBeerReq) (*v1.ListBeerReply, error) {
	rv, err := s.bc.List(ctx, req.PageNum, req.PageSize)
	rs := make([]*v1.ListBeerReply_Beer, 0)
	for _, x := range rv {
		img := make([]*v1.ListBeerReply_Beer_Image, 0)
		for _, i := range x.Images {
			img = append(img, &v1.ListBeerReply_Beer_Image{Url: i.URL})
		}
		rs = append(rs, &v1.ListBeerReply_Beer{
			Id:          x.Id,
			Name:        x.Name,
			Description: x.Description,
			Image:       img,
		})
	}
	return &v1.ListBeerReply{
		Results: rs,
	}, err
}
