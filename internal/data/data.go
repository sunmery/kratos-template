package data

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"

	"github.com/jackc/pgx/v5/pgxpool"
	"payment/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewCache, NewGreeterRepo)

// Data .
type Data struct {
	db  *modules.Queries
	rdb *redis.Client
}

// NewData .
func NewData(
	logger log.Logger,
	pgx *pgxpool.Pool,
	rdb *redis.Client,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db:  modules.New(pgx),
		rdb: rdb,
	}, cleanup, nil
}

func NewDB(c *conf.Data) *pgxpool.Pool {

	conn, err := pgxpool.New(context.Background(), c.Database.Source)
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v", err))
	}

	return conn
}

func NewCache(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Protocol: 3,
		Addr:     c.Redis.Addr,
		// Username:     c.Redis.Username,
		// Password:     c.Redis.Password,
		// DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	})

	return rdb
}
