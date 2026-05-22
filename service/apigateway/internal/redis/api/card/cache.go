package card_cache

import (
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/cache"
	card_dashboard_cache "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/redis/api/card/dashboard"
	card_stats_cache "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/redis/api/card/stats"
	card_stats_bycard_cache "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/redis/api/card/statsbycard"
)

type CardMencache interface {
	CardQueryCache
	CardCommandCache
	card_stats_cache.CardStatsCache
	card_stats_bycard_cache.CardStatsByCardCache
	card_dashboard_cache.CardDashboardCache
}

type cardmencache struct {
	CardQueryCache
	CardCommandCache
	card_stats_cache.CardStatsCache
	card_stats_bycard_cache.CardStatsByCardCache
	card_dashboard_cache.CardDashboardCache
}

func NewCardMencache(cacheStore *cache.CacheStore) CardMencache {

	return &cardmencache{
		CardCommandCache:     NewCardCommandCache(cacheStore),
		CardQueryCache:       NewCardQueryCache(cacheStore),
		CardStatsCache:       card_stats_cache.NewMencacheStats(cacheStore),
		CardStatsByCardCache: card_stats_bycard_cache.NewMencacheStatsByCard(cacheStore),
		CardDashboardCache:   card_dashboard_cache.NewMencacheDashboard(cacheStore),
	}
}
