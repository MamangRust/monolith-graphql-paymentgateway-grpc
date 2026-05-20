package graph

import (
	"time"

	authgraphqlmapper "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/mapper/auth"
	cardgraphqlmapper "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/mapper/card"
	merchantgraphqlmapper "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/mapper/merchant"
	merchantdocumentgraphqlmapper "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/mapper/merchant_document"
	rolegraphqlmapper "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/mapper/role"
	saldographqlmapper "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/mapper/saldo"
	topupgraphqlmapper "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/mapper/topup"
	transactiongraphqlmapper "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/mapper/transaction"
	transfergraphqlmapper "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/mapper/transfer"
	usergraphqlmapper "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/mapper/user"
	withdrawgraphqlmapper "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/mapper/withdraw"
	merchantpermission "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/permission/merchant"
	rolepermission "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/permission/role"
	mencache "github.com/MamangRust/monolith-graphql-payment-gateway-apigateway/internal/redis"
	authpb "github.com/MamangRust/monolith-graphql-payment-gateway-pb"
	cardpb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/card"
	cardstatpb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/card/stats"
	merchantpb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/merchant"
	merchantstatpb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/merchant/stats"
	merchantdocumentpb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/merchant_document"
	rolepb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/role"
	saldopb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/saldo"
	saldostatspb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/saldo/stats"
	topuppb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/topup"
	topupstatpb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/topup/stats"
	transactionpb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/transaction"
	transactionstatpb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/transaction/stats"
	transferpb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/transfer"
	transferstatpb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/transfer/stats"
	userpb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/user"
	withdrawpb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/withdraw"
	withdrawstatpb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/withdraw/stats"
	"github.com/MamangRust/monolith-graphql-payment-gateway-pkg/kafka"
	"github.com/MamangRust/monolith-graphql-payment-gateway-pkg/logger"
	"google.golang.org/grpc"
	"fmt"
	"github.com/go-playground/validator/v10"
	sharedErrors "github.com/MamangRust/monolith-graphql-payment-gateway-shared/errors"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthGraphql             AuthHandleGraphql
	RoleGraphql             RoleHandleGraphql
	UserGraphql             UserHandleGraphql
	CardGraphql             CardHandleGraphql
	MerchantGraphql         MerchantHandleGraphql
	MerchantDocumentGraphql MerchantDocumentHandleGraphql
	SaldoGraphql            SaldoHandleGraphql
	TopupGraphql            TopupHandleGraphql
	TransactionGraphql      TransactionHandleGraphql
	TransferGraphql         TransferHandleGraphql
	WithdrawGraphql         WithdrawHandleGraphql
}

type UserClient struct {
	UserQueryClient   userpb.UserQueryServiceClient
	UserCommandClient userpb.UserCommandServiceClient
}

type RoleClient struct {
	RoleQueryClient   rolepb.RoleServiceClient
	RoleCommandClient rolepb.RoleCommandServiceClient
}

type CardClient struct {
	CardQueryClient                  cardpb.CardQueryServiceClient
	CardCommandClient                cardpb.CardCommandServiceClient
	CardDashboardClient              cardpb.CardDashboardServiceClient
	CardStatsBalanceClient           cardstatpb.CardStatsBalanceServiceClient
	CardStatsTopupAmountClient       cardstatpb.CardStatsTopupServiceClient
	CardStatsTransactionAmountClient cardstatpb.CardStatsTransactionServiceClient
	CardStatsWithdrawAmountClient    cardstatpb.CardStatsWithdrawServiceClient
	CardStatsTransferAmountClient    cardstatpb.CardStatsTransferServiceClient
}

type MerchantClient struct {
	MerchantQuery            merchantpb.MerchantQueryServiceClient
	MerchantCommand          merchantpb.MerchantCommandServiceClient
	MerchantTransaction      merchantpb.MerchantTransactionServiceClient
	MerchantStatsAmount      merchantstatpb.MerchantStatsAmountServiceClient
	MerchantStatsTotalAmount merchantstatpb.MerchantStatsTotalAmountServiceClient
	MerchantStatsMethod      merchantstatpb.MerchantStatsMethodServiceClient
}

type MerchantDocumentClient struct {
	MerchantDocumentQueryClient   merchantdocumentpb.MerchantDocumentQueryServiceClient
	MerchantDocumentCommandClient merchantdocumentpb.MerchantDocumentCommandServiceClient
}

type SaldoClient struct {
	SaldoQueryClient             saldopb.SaldoQueryServiceClient
	SaldoCommandClient           saldopb.SaldoCommandServiceClient
	SaldoStatsBalanceClient      saldostatspb.SaldoStatsBalanceServiceClient
	SaldoStatsTotalBalanceClient saldostatspb.SaldoStatsTotalBalanceClient
}

type TopupClient struct {
	TopupQueryClient   topuppb.TopupQueryServiceClient
	TopupCommandClient topuppb.TopupCommandServiceClient
	TopupStatsAmount   topupstatpb.TopupStatsAmountServiceClient
	TopupStatsMethod   topupstatpb.TopupStatsMethodServiceClient
	TopupStatsStatus   topupstatpb.TopupStatsStatusServiceClient
}

type TransactionClient struct {
	TransactionQueryClient   transactionpb.TransactionQueryServiceClient
	TransactionCommandClient transactionpb.TransactionCommandServiceClient
	TransactionStatsAmount   transactionstatpb.TransactionStatsAmountServiceClient
	TransactionStatsMethod   transactionstatpb.TransactionStatsMethodServiceClient
	TransactionStatsStatus   transactionstatpb.TransactionStatsStatusServiceClient
}

type TransferClient struct {
	TransferQueryClient   transferpb.TransferQueryServiceClient
	TransferCommandClient transferpb.TransferCommandServiceClient
	TransferStatsAmount   transferstatpb.TransferStatsAmountServiceClient
	TransferStatsStatus   transferstatpb.TransferStatsStatusServiceClient
}
type WithdrawClient struct {
	WithdrawQueryClient   withdrawpb.WithdrawQueryServiceClient
	WithdrawCommandClient withdrawpb.WithdrawCommandServiceClient
	WithdrawStatsAmount   withdrawstatpb.WithdrawStatsAmountServiceClient
	WithdrawStatsStatus   withdrawstatpb.WithdrawStatsStatusServiceClient
}

type AuthHandleGraphql struct {
	AuthClient authpb.AuthServiceClient
	Logger     logger.LoggerInterface
	Mapping    authgraphqlmapper.AuthGraphqlMapper
}

type RoleHandleGraphql struct {
	RoleClient RoleClient
	Logger     logger.LoggerInterface
	Mapping    rolegraphqlmapper.RoleGraphqlMapper
	Kafka      *kafka.Kafka
	CacheRole  mencache.RoleCache
	Permission rolepermission.RolePermission
}

type UserHandleGraphql struct {
	UserClient UserClient
	Logger     logger.LoggerInterface
	Mapping    usergraphqlmapper.UserGraphqlMapper
}

type CardHandleGraphql struct {
	CardClient CardClient
	Logger     logger.LoggerInterface
	Mapping    cardgraphqlmapper.CardGraphqlMapper
}

type MerchantHandleGraphql struct {
	MerchantClient MerchantClient
	Logger         logger.LoggerInterface
	Mapping        merchantgraphqlmapper.MerchantGraphqlMapper
}

type MerchantDocumentHandleGraphql struct {
	MerchantClient MerchantDocumentClient
	Logger         logger.LoggerInterface
	Mapping        merchantdocumentgraphqlmapper.MerchantDocumentGraphqlMapper
}

type SaldoHandleGraphql struct {
	SaldoClient SaldoClient
	Logger      logger.LoggerInterface
	Mapping     saldographqlmapper.SaldoGraphqlMapper
}

type TopupHandleGraphql struct {
	TopupClient TopupClient
	Logger      logger.LoggerInterface
	Mapping     topupgraphqlmapper.TopupGraphqlMapper
}

type TransactionHandleGraphql struct {
	TransactionClient TransactionClient
	Logger            logger.LoggerInterface
	Mapping           transactiongraphqlmapper.TransactionGraphqlMapper
	Permission        merchantpermission.MerchantPermission
	CacheMerchant     mencache.MerchantCache
}

type TransferHandleGraphql struct {
	TransferClient TransferClient
	Logger         logger.LoggerInterface
	Mapping        transfergraphqlmapper.TransferGraphqlMapper
}

type WithdrawHandleGraphql struct {
	WithdrawClient WithdrawClient
	Logger         logger.LoggerInterface
	Mapping        withdrawgraphqlmapper.WithdrawGraphqlMapper
}

type ServiceConnections struct {
	AuthClient        *grpc.ClientConn
	RoleClient        *grpc.ClientConn
	UserClient        *grpc.ClientConn
	CardClient        *grpc.ClientConn
	MerchantClient    *grpc.ClientConn
	SaldoClient       *grpc.ClientConn
	TopupClient       *grpc.ClientConn
	TransactionClient *grpc.ClientConn
	TransferClient    *grpc.ClientConn
	WithdrawClient    *grpc.ClientConn
}

func NewResolver(
	clients *ServiceConnections,
	logger logger.LoggerInterface,
	kafka *kafka.Kafka,
	mencache mencache.CacheApiGateway,
) *Resolver {
	return &Resolver{
		AuthGraphql: AuthHandleGraphql{
			AuthClient: authpb.NewAuthServiceClient(clients.AuthClient),
			Logger:     logger,
			Mapping:    authgraphqlmapper.NewAuthGraphqlMapper(),
		},
		RoleGraphql: RoleHandleGraphql{
			RoleClient: RoleClient{
				RoleQueryClient:   rolepb.NewRoleServiceClient(clients.RoleClient),
				RoleCommandClient: rolepb.NewRoleCommandServiceClient(clients.RoleClient),
			},
			Kafka:      kafka,
			Logger:     logger,
			Mapping:    rolegraphqlmapper.NewRoleGraphqlMapper(),
			Permission: rolepermission.NewRolePermission(kafka, "request-role", "response-role", 5*time.Second, logger, mencache),
		},
		UserGraphql: UserHandleGraphql{
			UserClient: UserClient{
				UserQueryClient:   userpb.NewUserQueryServiceClient(clients.UserClient),
				UserCommandClient: userpb.NewUserCommandServiceClient(clients.UserClient),
			},
			Logger:  logger,
			Mapping: usergraphqlmapper.NewUserGraphqlMapper(),
		},
		CardGraphql: CardHandleGraphql{
			CardClient: CardClient{
				CardQueryClient:                  cardpb.NewCardQueryServiceClient(clients.CardClient),
				CardCommandClient:                cardpb.NewCardCommandServiceClient(clients.CardClient),
				CardDashboardClient:              cardpb.NewCardDashboardServiceClient(clients.CardClient),
				CardStatsBalanceClient:           cardstatpb.NewCardStatsBalanceServiceClient(clients.CardClient),
				CardStatsTopupAmountClient:       cardstatpb.NewCardStatsTopupServiceClient(clients.CardClient),
				CardStatsTransactionAmountClient: cardstatpb.NewCardStatsTransactionServiceClient(clients.CardClient),
				CardStatsWithdrawAmountClient:    cardstatpb.NewCardStatsWithdrawServiceClient(clients.CardClient),
				CardStatsTransferAmountClient:    cardstatpb.NewCardStatsTransferServiceClient(clients.CardClient),
			},
			Logger:  logger,
			Mapping: cardgraphqlmapper.NewCardResponseMapper(),
		},
		MerchantGraphql: MerchantHandleGraphql{
			MerchantClient: MerchantClient{
				MerchantQuery:            merchantpb.NewMerchantQueryServiceClient(clients.MerchantClient),
				MerchantCommand:          merchantpb.NewMerchantCommandServiceClient(clients.MerchantClient),
				MerchantTransaction:      merchantpb.NewMerchantTransactionServiceClient(clients.MerchantClient),
				MerchantStatsAmount:      merchantstatpb.NewMerchantStatsAmountServiceClient(clients.MerchantClient),
				MerchantStatsTotalAmount: merchantstatpb.NewMerchantStatsTotalAmountServiceClient(clients.MerchantClient),
				MerchantStatsMethod:      merchantstatpb.NewMerchantStatsMethodServiceClient(clients.MerchantClient),
			},
			Logger:  logger,
			Mapping: merchantgraphqlmapper.NewMerchantResponseMapper(),
		},
		MerchantDocumentGraphql: MerchantDocumentHandleGraphql{
			MerchantClient: MerchantDocumentClient{
				MerchantDocumentQueryClient:   merchantdocumentpb.NewMerchantDocumentQueryServiceClient(clients.MerchantClient),
				MerchantDocumentCommandClient: merchantdocumentpb.NewMerchantDocumentCommandServiceClient(clients.MerchantClient),
			},
			Logger:  logger,
			Mapping: merchantdocumentgraphqlmapper.NewMerchantDocumentGraphqlMapper(),
		},
		SaldoGraphql: SaldoHandleGraphql{
			SaldoClient: SaldoClient{
				SaldoQueryClient:             saldopb.NewSaldoQueryServiceClient(clients.SaldoClient),
				SaldoCommandClient:           saldopb.NewSaldoCommandServiceClient(clients.SaldoClient),
				SaldoStatsBalanceClient:      saldostatspb.NewSaldoStatsBalanceServiceClient(clients.SaldoClient),
				SaldoStatsTotalBalanceClient: saldostatspb.NewSaldoStatsTotalBalanceClient(clients.SaldoClient),
			},
			Logger:  logger,
			Mapping: saldographqlmapper.NewSaldoGraphqlMapper(),
		},
		TopupGraphql: TopupHandleGraphql{
			TopupClient: TopupClient{
				TopupQueryClient:   topuppb.NewTopupQueryServiceClient(clients.TopupClient),
				TopupCommandClient: topuppb.NewTopupCommandServiceClient(clients.TopupClient),
				TopupStatsAmount:   topupstatpb.NewTopupStatsAmountServiceClient(clients.TopupClient),
				TopupStatsMethod:   topupstatpb.NewTopupStatsMethodServiceClient(clients.TopupClient),
				TopupStatsStatus:   topupstatpb.NewTopupStatsStatusServiceClient(clients.TopupClient),
			},
			Logger:  logger,
			Mapping: topupgraphqlmapper.NewTopupGraphqlMapper(),
		},
		TransactionGraphql: TransactionHandleGraphql{
			TransactionClient: TransactionClient{
				TransactionQueryClient:   transactionpb.NewTransactionQueryServiceClient(clients.TransactionClient),
				TransactionCommandClient: transactionpb.NewTransactionCommandServiceClient(clients.TransactionClient),
				TransactionStatsAmount:   transactionstatpb.NewTransactionStatsAmountServiceClient(clients.TransactionClient),
				TransactionStatsMethod:   transactionstatpb.NewTransactionStatsMethodServiceClient(clients.TransactionClient),
				TransactionStatsStatus:   transactionstatpb.NewTransactionStatsStatusServiceClient(clients.TransactionClient),
			},
			Logger:     logger,
			Mapping:    transactiongraphqlmapper.NewTransactionGraphqlMapper(),
			Permission: merchantpermission.NewMerchantPermission(kafka, "request-transaction", "response-transaction", 5*time.Second, logger),
		},
		TransferGraphql: TransferHandleGraphql{
			TransferClient: TransferClient{
				TransferQueryClient:   transferpb.NewTransferQueryServiceClient(clients.TransferClient),
				TransferCommandClient: transferpb.NewTransferCommandServiceClient(clients.TransferClient),
				TransferStatsAmount:   transferstatpb.NewTransferStatsAmountServiceClient(clients.TransferClient),
				TransferStatsStatus:   transferstatpb.NewTransferStatsStatusServiceClient(clients.TransferClient),
			},
			Logger:  logger,
			Mapping: transfergraphqlmapper.NewTransferGraphqlMapper(),
		},
		WithdrawGraphql: WithdrawHandleGraphql{
			WithdrawClient: WithdrawClient{
				WithdrawQueryClient:   withdrawpb.NewWithdrawQueryServiceClient(clients.WithdrawClient),
				WithdrawCommandClient: withdrawpb.NewWithdrawCommandServiceClient(clients.WithdrawClient),
				WithdrawStatsAmount:   withdrawstatpb.NewWithdrawStatsAmountServiceClient(clients.WithdrawClient),
				WithdrawStatsStatus:   withdrawstatpb.NewWithdrawStatsStatusServiceClient(clients.WithdrawClient),
			},
			Logger:  logger,
			Mapping: withdrawgraphqlmapper.NewWithdrawGraphqlMapper(),
		},
	}
}

func (r *Resolver) parseValidationErrors(err error) []sharedErrors.ValidationError {
	var validationErrs []sharedErrors.ValidationError

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			validationErrs = append(validationErrs, sharedErrors.ValidationError{
				Field:   fe.Field(),
				Message: r.getValidationMessage(fe),
			})
		}
		return validationErrs
	}

	return []sharedErrors.ValidationError{
		{
			Field:   "general",
			Message: err.Error(),
		},
	}
}

func (r *Resolver) getValidationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Must be at least %s", fe.Param())
	case "max":
		return fmt.Sprintf("Must be at most %s", fe.Param())
	case "gte":
		return fmt.Sprintf("Must be greater than or equal to %s", fe.Param())
	case "lte":
		return fmt.Sprintf("Must be less than or equal to %s", fe.Param())
	case "oneof":
		return fmt.Sprintf("Must be one of: %s", fe.Param())
	default:
		return fmt.Sprintf("Validation failed on '%s' tag", fe.Tag())
	}
}
