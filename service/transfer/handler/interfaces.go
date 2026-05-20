package handler

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/transfer"
)

type TransferQueryHandleGrpc interface {
	pb.TransferQueryServiceServer
}

type TransferCommandHandleGrpc interface {
	pb.TransferCommandServiceServer
}
