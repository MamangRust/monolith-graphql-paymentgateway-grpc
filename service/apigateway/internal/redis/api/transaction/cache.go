package transaction_cache

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/cache"
	transaction_stats_cache "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/redis/api/transaction/stats"
	transaction_stats_bycard_cache "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/redis/api/transaction/statsbycard"
)

type TransactionMencache interface {
	TransactionQueryCache
	TransactionCommandCache
	transaction_stats_cache.TransactionStatsCache
	transaction_stats_bycard_cache.TransactionStatsByCardCache
}

type transactionmencache struct {
	TransactionQueryCache
	TransactionCommandCache
	transaction_stats_cache.TransactionStatsCache
	transaction_stats_bycard_cache.TransactionStatsByCardCache
}

func NewTransactionMencache(cacheStore *cache.CacheStore) TransactionMencache {
	return &transactionmencache{
		TransactionQueryCache:       NewTransactionQueryCache(cacheStore),
		TransactionCommandCache:     NewTransactionCommandCache(cacheStore),
		TransactionStatsCache:       transaction_stats_cache.NewTransactionStatsCache(cacheStore),
		TransactionStatsByCardCache: transaction_stats_bycard_cache.NewTransactionStatsByCardCache(cacheStore),
	}
}
