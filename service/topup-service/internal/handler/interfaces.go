package handler

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/topup"
)

type TopupQueryHandleGrpc interface {
	pb.TopupQueryServiceServer
}

type TopupCommandHandleGrpc interface {
	pb.TopupCommandServiceServer
}
