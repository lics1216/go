package cache

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	"github.com/google/wire"

	"eshop/cmd/user/internal/conf"
	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewRedisClient, NewData, NewUserCache)

// Data .
type Data struct {
	cache *redis.Client
}

func NewRedisClient(conf *conf.Data, logger log.Logger) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: conf.Redis.Addr,
	})
}

// NewData .
func NewData(cache *redis.Client, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "user-service/data/cache"))

	d := &Data{
		cache: cache,
	}
	return d, func() {
		if err := d.cache.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
