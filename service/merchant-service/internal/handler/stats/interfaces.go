package merchantstatshandler

import (
	pbmerchant "github.com/MamangRust/monolith-graphql-payment-gateway-pb/merchant/stats"
)

type MerchantStatsAmountHandleGrpc interface {
	pbmerchant.MerchantStatsAmountServiceServer
}

type MerchantStatsMethodHandleGrpc interface {
	pbmerchant.MerchantStatsMethodServiceServer
}

type MerchantStatsTotalAmountHandleGrpc interface {
	pbmerchant.MerchantStatsTotalAmountServiceServer
}
