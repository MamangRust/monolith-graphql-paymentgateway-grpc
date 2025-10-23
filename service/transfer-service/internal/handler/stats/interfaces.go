package transferstatshandler

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/transfer/stats"
)

type TransferStatsAmountHandleGrpc interface {
	pb.TransferStatsAmountServiceServer
}

type TransferStatsStatusHandleGrpc interface {
	pb.TransferStatsStatusServiceServer
}
