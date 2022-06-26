package service

import (
	v1a "eshop/api/catalog/admin/v1"
	"eshop/api/catalog/v1"
	"eshop/cmd/catalog/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCatalogService)

type CatalogService struct {
	v1.UnimplementedCatalogServer
	v1a.UnimplementedCatalogAdminServer

	bc  *biz.BeerUseCase
	log *log.Helper
}

func NewCatalogService(bc *biz.BeerUseCase, logger log.Logger) *CatalogService {
	return &CatalogService{

		bc:  bc,
		log: log.NewHelper(log.With(logger, "module", "service/catalog"))}
}
