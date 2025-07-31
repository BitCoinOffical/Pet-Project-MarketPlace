package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"myapp/infrastructure/cache/storage"
	"myapp/internal/interfaces/http/dto"
	usecase "myapp/internal/usecase/products"
	"time"

	"github.com/redis/go-redis/v9"
)

type GetAllCashe struct {
	usecase *usecase.UseCases
	rdb     *redis.Client
}

func NewGetAllCash(usecase *usecase.UseCases, rdb *redis.Client) *GetAllCashe {
	return &GetAllCashe{usecase: usecase, rdb: rdb}
}

func Redisinit() (*redis.Client, error) {
	ctx := context.Background()
	cfg := storage.Config{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	}

	rdb, err := storage.NewClient(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return rdb, nil

}

func (uc *GetAllCashe) GetDataWithCache(ctx context.Context, category string, search string, page int, min_price float64, max_price float64) (string, []dto.ProductResponse, int, error) {
	key := fmt.Sprint("key: ", category, search, page, min_price, max_price)
	val, err := uc.rdb.Get(ctx, key).Result()
	if err == nil {
		var list dto.ProductList
		if err := json.Unmarshal([]byte(val), &list); err == nil {
			return "", list.Items, list.TotalCount, nil
		}
	} else if err != redis.Nil {
		return "", nil, 0, err
	}

	products, total_count, err := uc.usecase.GetAll.GetAll(ctx, category, search, page, min_price, max_price)
	if err != nil {
		return "", nil, 0, err
	}
	list := dto.ProductList{
		Items:      products,
		TotalCount: total_count,
	}
	bytes, err := json.Marshal(list)
	if err == nil {
		uc.rdb.Set(ctx, key, bytes, time.Minute*5)
	}
	uc.rdb.Set(ctx, fmt.Sprint("key: ", category, search, page, min_price, max_price), &dto.ProductList{Items: products, TotalCount: total_count}, time.Minute*5)
	return "", products, total_count, nil

}
