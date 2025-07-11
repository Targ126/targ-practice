package main

import (
	"context"
	"fmt"
	"github.com/NFTGalaxy/rick/pkg/database/cache"
	rickWpgx "github.com/NFTGalaxy/rick/pkg/database/wpgx"
	"sqlc_test/space"

	"github.com/stumble/dcache"
	"github.com/stumble/wpgx"
)

func main() {

	var (
		dbConnPool *wpgx.Pool
		dCache     *dcache.DCache
	)

	ctx := context.Background()

	pool, err := rickWpgx.NewWPGXPool(ctx, "postgres")
	if err != nil {
		panic(fmt.Errorf("failed to create db connection pool: %w", err))
	}
	dbConnPool = pool

	redisConn := cache.NewRedisClient("redis")
	// create dcache instance
	dCache, err = cache.NewDCache("app_name", "dcache", redisConn)
	if err != nil {
		panic(fmt.Errorf("failed to create dcache: %w", err))
	}

	spaceDBClient := space.New(dbConnPool.WConn(), dCache)

	data, err := spaceDBClient.GetById(ctx, 1)
	if err != nil {
		panic(fmt.Errorf("failed to get space by id: %w", err))
	}
	fmt.Println(data)

}

// 1. 本地docker 部署 postgres和redis
// 2. 在postgres中创建table
// 3. 能够使用goland代码查询出db中的数据
