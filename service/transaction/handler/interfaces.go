package handler

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/transaction"
)

type TransactionQueryHandleGrpc interface {
	pb.TransactionQueryServiceServer
}

type TransactionCommandHandleGrpc interface {
	pb.TransactionCommandServiceServer
}
