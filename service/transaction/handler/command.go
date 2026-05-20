package handler

import (
	"context"
	"fmt"
	"time"

	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/transaction"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/errors"
	"github.com/MamangRust/monolith-graphql-payment-gateway-transaction/service"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"github.com/go-playground/validator/v10"
)

type transactionCommandHandleGrpc struct {
	pb.UnimplementedTransactionCommandServiceServer

	service service.TransactionCommandService
}

func NewTransactionCommandHandleGrpc(service service.TransactionCommandService) TransactionCommandHandleGrpc {
	return &transactionCommandHandleGrpc{
		service: service,
	}
}

func (t *transactionCommandHandleGrpc) CreateTransaction(ctx context.Context, request *pb.CreateTransactionRequest) (*pb.ApiResponseTransaction, error) {
	transactionTime := request.GetTransactionTime().AsTime()
	merchantID := int(request.GetMerchantId())

	req := requests.CreateTransactionRequest{
		CardNumber:      request.GetCardNumber(),
		Amount:          int(request.GetAmount()),
		PaymentMethod:   request.GetPaymentMethod(),
		MerchantID:      &merchantID,
		TransactionTime: transactionTime,
	}

	if err := req.Validate(); err != nil {
		validations := t.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	res, err := t.service.Create(ctx, request.ApiKey, &req)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransaction{
		Status:  "success",
		Message: "Successfully created transaction",
		Data: &pb.TransactionResponse{
			Id:              int32(res.TransactionID),
			CardNumber:      res.CardNumber,
			TransactionNo:   res.TransactionNo.String(),
			Amount:          int32(res.Amount),
			PaymentMethod:   res.PaymentMethod,
			MerchantId:      int32(res.MerchantID),
			TransactionTime: res.TransactionTime.Format(time.RFC3339),
			CreatedAt:       res.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       res.UpdatedAt.Time.Format(time.RFC3339),
		},
	}, nil
}

func (t *transactionCommandHandleGrpc) UpdateTransaction(ctx context.Context, request *pb.UpdateTransactionRequest) (*pb.ApiResponseTransaction, error) {
	id := int(request.GetTransactionId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "transaction_id",
				Message: "Transaction ID is required",
			},
		}))
	}

	transactionTime := request.GetTransactionTime().AsTime()
	merchantID := int(request.GetMerchantId())

	req := requests.UpdateTransactionRequest{
		TransactionID:   &id,
		CardNumber:      request.GetCardNumber(),
		Amount:          int(request.GetAmount()),
		PaymentMethod:   request.GetPaymentMethod(),
		MerchantID:      &merchantID,
		TransactionTime: transactionTime,
	}

	if err := req.Validate(); err != nil {
		validations := t.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	res, err := t.service.Update(ctx, request.ApiKey, &req)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransaction{
		Status:  "success",
		Message: "Successfully updated transaction",
		Data: &pb.TransactionResponse{
			Id:              int32(res.TransactionID),
			CardNumber:      res.CardNumber,
			TransactionNo:   res.TransactionNo.String(),
			Amount:          int32(res.Amount),
			PaymentMethod:   res.PaymentMethod,
			MerchantId:      int32(res.MerchantID),
			TransactionTime: res.TransactionTime.Format(time.RFC3339),
			CreatedAt:       res.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       res.UpdatedAt.Time.Format(time.RFC3339),
		},
	}, nil
}

func (t *transactionCommandHandleGrpc) parseValidationErrors(err error) []errors.ValidationError {
	var validationErrs []errors.ValidationError

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			validationErrs = append(validationErrs, errors.ValidationError{
				Field:   fe.Field(),
				Message: t.getValidationMessage(fe),
			})
		}
		return validationErrs
	}

	return []errors.ValidationError{
		{
			Field:   "general",
			Message: err.Error(),
		},
	}
}

func (t *transactionCommandHandleGrpc) getValidationMessage(fe validator.FieldError) string {
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

func (t *transactionCommandHandleGrpc) TrashedTransaction(ctx context.Context, request *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDeleteAt, error) {
	id := int(request.GetTransactionId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "transaction_id",
				Message: "Transaction ID is required",
			},
		}))
	}

	res, err := t.service.TrashedTransaction(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransactionDeleteAt{
		Status:  "success",
		Message: "Successfully trashed transaction",
		Data: &pb.TransactionResponseDeleteAt{
			Id:              int32(res.TransactionID),
			CardNumber:      res.CardNumber,
			TransactionNo:   res.TransactionNo.String(),
			Amount:          int32(res.Amount),
			PaymentMethod:   res.PaymentMethod,
			MerchantId:      int32(res.MerchantID),
			TransactionTime: res.TransactionTime.Format(time.RFC3339),
			CreatedAt:       res.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       res.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:       &wrapperspb.StringValue{Value: res.DeletedAt.Time.Format(time.RFC3339)},
		},
	}, nil
}

func (t *transactionCommandHandleGrpc) RestoreTransaction(ctx context.Context, request *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDeleteAt, error) {
	id := int(request.GetTransactionId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "transaction_id",
				Message: "Transaction ID is required",
			},
		}))
	}

	res, err := t.service.RestoreTransaction(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransactionDeleteAt{
		Status:  "success",
		Message: "Successfully restored transaction",
		Data: &pb.TransactionResponseDeleteAt{
			Id:              int32(res.TransactionID),
			CardNumber:      res.CardNumber,
			TransactionNo:   res.TransactionNo.String(),
			Amount:          int32(res.Amount),
			PaymentMethod:   res.PaymentMethod,
			MerchantId:      int32(res.MerchantID),
			TransactionTime: res.TransactionTime.Format(time.RFC3339),
			CreatedAt:       res.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:       res.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:       &wrapperspb.StringValue{Value: res.DeletedAt.Time.Format(time.RFC3339)},
		},
	}, nil
}

func (t *transactionCommandHandleGrpc) DeleteTransactionPermanent(ctx context.Context, request *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDelete, error) {
	id := int(request.GetTransactionId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "transaction_id",
				Message: "Transaction ID is required",
			},
		}))
	}

	_, err := t.service.DeleteTransactionPermanent(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransactionDelete{
		Status:  "success",
		Message: "Successfully deleted transaction",
	}, nil
}

func (t *transactionCommandHandleGrpc) RestoreAllTransaction(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTransactionAll, error) {
	_, err := t.service.RestoreAllTransaction(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransactionAll{
		Status:  "success",
		Message: "Successfully restore all transaction",
	}, nil
}

func (t *transactionCommandHandleGrpc) DeleteAllTransactionPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTransactionAll, error) {
	_, err := t.service.DeleteAllTransactionPermanent(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransactionAll{
		Status:  "success",
		Message: "Successfully delete transaction permanent",
	}, nil
}
