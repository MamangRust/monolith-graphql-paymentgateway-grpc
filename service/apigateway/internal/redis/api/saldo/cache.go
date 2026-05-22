package saldo_cache

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/cache"
	saldo_stats_cache "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/redis/api/saldo/stats"
)

type SaldoMencache interface {
	SaldoQueryCache
	SaldoCommandCache
	saldo_stats_cache.SaldoStatsCache
}

type saldomencache struct {
	SaldoQueryCache
	SaldoCommandCache
	saldo_stats_cache.SaldoStatsCache
}

func NewSaldoMencache(cacheStore *cache.CacheStore) SaldoMencache {
	return &saldomencache{
		SaldoQueryCache:   NewSaldoQueryCache(cacheStore),
		SaldoCommandCache: NewSaldoCommandCache(cacheStore),
		SaldoStatsCache:   saldo_stats_cache.NewSaldoStatsCache(cacheStore),
	}
}
