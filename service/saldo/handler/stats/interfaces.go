package saldostatshandler

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/saldo/stats"
)

type SaldoStatsBalanceHandleGrpc interface {
	pb.SaldoStatsBalanceServiceServer
}

type SaldoStatsTotalBalanceHandleGrpc interface {
	pb.SaldoStatsTotalBalanceServer
}
