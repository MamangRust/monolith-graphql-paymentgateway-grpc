package handler

import (
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/user"
)

type UserQueryHandleGrpc interface {
	pb.UserQueryServiceServer
}

type UserCommandHandleGrpc interface {
	pb.UserCommandServiceServer
}
