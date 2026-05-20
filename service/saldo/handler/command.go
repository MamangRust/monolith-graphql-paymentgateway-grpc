package handler

import (
	"context"
	"fmt"
	"time"

	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/saldo"
	"github.com/MamangRust/monolith-graphql-payment-gateway-saldo/service"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"github.com/go-playground/validator/v10"
)

type saldoCommandHandleGrpc struct {
	pb.UnimplementedSaldoCommandServiceServer

	service service.SaldoCommandService
}

func NewSaldoCommandHandleGrpc(query service.SaldoCommandService) SaldoCommandHandleGrpc {
	return &saldoCommandHandleGrpc{
		service: query,
	}
}

func (s *saldoCommandHandleGrpc) CreateSaldo(ctx context.Context, req *pb.CreateSaldoRequest) (*pb.ApiResponseSaldo, error) {
	request := requests.CreateSaldoRequest{
		CardNumber:   req.GetCardNumber(),
		TotalBalance: int(req.GetTotalBalance()),
	}

	if err := request.Validate(); err != nil {
		validations := s.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	saldo, err := s.service.CreateSaldo(ctx, &request)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoSaldo := &pb.SaldoResponse{
		SaldoId:        int32(saldo.SaldoID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int32(saldo.TotalBalance),
		WithdrawTime:   saldo.WithdrawTime.Time.Format(time.RFC3339),
		WithdrawAmount: Int32Value(saldo.WithdrawAmount),
		CreatedAt:      saldo.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:      saldo.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseSaldo{
		Status:  "success",
		Message: "Successfully created saldo record",
		Data:    protoSaldo,
	}, nil
}

func (s *saldoCommandHandleGrpc) UpdateSaldo(ctx context.Context, req *pb.UpdateSaldoRequest) (*pb.ApiResponseSaldo, error) {
	id := int(req.GetSaldoId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "saldo_id",
				Message: "Saldo ID is required",
			},
		}))
	}

	request := requests.UpdateSaldoRequest{
		SaldoID:      &id,
		CardNumber:   req.GetCardNumber(),
		TotalBalance: int(req.GetTotalBalance()),
	}

	if err := request.Validate(); err != nil {
		validations := s.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}
	saldo, err := s.service.UpdateSaldo(ctx, &request)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoSaldo := &pb.SaldoResponse{
		SaldoId:        int32(saldo.SaldoID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int32(saldo.TotalBalance),
		WithdrawTime:   saldo.WithdrawTime.Time.Format(time.RFC3339),
		WithdrawAmount: Int32Value(saldo.WithdrawAmount),
		CreatedAt:      saldo.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:      saldo.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseSaldo{
		Status:  "success",
		Message: "Successfully updated saldo record",
		Data:    protoSaldo,
	}, nil
}

func (s *saldoCommandHandleGrpc) UpdateSaldoBalance(ctx context.Context, req *pb.UpdateSaldoBalanceRequest) (*pb.ApiResponseSaldo, error) {
	request := requests.UpdateSaldoBalance{
		CardNumber:   req.GetCardNumber(),
		TotalBalance: int(req.GetTotalBalance()),
	}

	if err := request.Validate(); err != nil {
		validations := s.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	saldo, err := s.service.UpdateSaldoBalance(ctx, &request)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoSaldo := &pb.SaldoResponse{
		SaldoId:        int32(saldo.SaldoID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int32(saldo.TotalBalance),
		WithdrawTime:   saldo.WithdrawTime.Time.Format(time.RFC3339),
		WithdrawAmount: Int32Value(saldo.WithdrawAmount),
		CreatedAt:      saldo.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:      saldo.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseSaldo{
		Status:  "success",
		Message: "Successfully updated saldo balance record",
		Data:    protoSaldo,
	}, nil
}

func (s *saldoCommandHandleGrpc) UpdateSaldoWithdraw(ctx context.Context, req *pb.UpdateSaldoWithdrawRequest) (*pb.ApiResponseSaldo, error) {
	withdrawTime := time.Now()
	withdrawAmount := int(req.GetWithdrawAmount())

	request := requests.UpdateSaldoWithdraw{
		CardNumber:     req.GetCardNumber(),
		TotalBalance:   int(req.GetTotalBalance()),
		WithdrawTime:   &withdrawTime,
		WithdrawAmount: &withdrawAmount,
	}

	if err := request.Validate(); err != nil {
		validations := s.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	saldo, err := s.service.UpdateSaldoWithdraw(ctx, &request)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoSaldo := &pb.SaldoResponse{
		SaldoId:        int32(saldo.SaldoID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int32(saldo.TotalBalance),
		WithdrawTime:   saldo.WithdrawTime.Time.Format(time.RFC3339),
		WithdrawAmount: Int32Value(saldo.WithdrawAmount),
		CreatedAt:      saldo.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:      saldo.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseSaldo{
		Status:  "success",
		Message: "Successfully updated saldo withdraw record",
		Data:    protoSaldo,
	}, nil
}

func (s *saldoCommandHandleGrpc) parseValidationErrors(err error) []errors.ValidationError {
	var validationErrs []errors.ValidationError

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			validationErrs = append(validationErrs, errors.ValidationError{
				Field:   fe.Field(),
				Message: s.getValidationMessage(fe),
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

func (s *saldoCommandHandleGrpc) getValidationMessage(fe validator.FieldError) string {
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

func (s *saldoCommandHandleGrpc) TrashedSaldo(ctx context.Context, req *pb.FindByIdSaldoRequest) (*pb.ApiResponseSaldoDeleteAt, error) {
	id := int(req.GetSaldoId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "saldo_id",
				Message: "Saldo ID is required",
			},
		}))
	}

	saldo, err := s.service.TrashSaldo(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoSaldo := &pb.SaldoResponseDeleteAt{
		SaldoId:        int32(saldo.SaldoID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int32(saldo.TotalBalance),
		WithdrawTime:   saldo.WithdrawTime.Time.Format(time.RFC3339),
		WithdrawAmount: Int32Value(saldo.WithdrawAmount),
		CreatedAt:      saldo.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:      saldo.UpdatedAt.Time.Format(time.RFC3339),
		DeletedAt:      wrapperspb.String(saldo.DeletedAt.Time.Format(time.RFC3339)),
	}

	return &pb.ApiResponseSaldoDeleteAt{
		Status:  "success",
		Message: "Successfully trashed saldo record",
		Data:    protoSaldo,
	}, nil
}

func (s *saldoCommandHandleGrpc) RestoreSaldo(ctx context.Context, req *pb.FindByIdSaldoRequest) (*pb.ApiResponseSaldoDeleteAt, error) {
	id := int(req.GetSaldoId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "saldo_id",
				Message: "Saldo ID is required",
			},
		}))
	}

	saldo, err := s.service.RestoreSaldo(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoSaldo := &pb.SaldoResponseDeleteAt{
		SaldoId:        int32(saldo.SaldoID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int32(saldo.TotalBalance),
		WithdrawTime:   saldo.WithdrawTime.Time.Format(time.RFC3339),
		WithdrawAmount: Int32Value(saldo.WithdrawAmount),
		CreatedAt:      saldo.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:      saldo.UpdatedAt.Time.Format(time.RFC3339),
		DeletedAt:      wrapperspb.String(saldo.DeletedAt.Time.Format(time.RFC3339)),
	}

	return &pb.ApiResponseSaldoDeleteAt{
		Status:  "success",
		Message: "Successfully restored saldo record",
		Data:    protoSaldo,
	}, nil
}

func (s *saldoCommandHandleGrpc) DeleteSaldoPermanent(ctx context.Context, req *pb.FindByIdSaldoRequest) (*pb.ApiResponseSaldoDelete, error) {
	id := int(req.GetSaldoId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "saldo_id",
				Message: "Saldo ID is required",
			},
		}))
	}

	_, err := s.service.DeleteSaldoPermanent(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseSaldoDelete{
		Status:  "success",
		Message: "Successfully deleted saldo record",
	}, nil
}

func (s *saldoCommandHandleGrpc) RestoreAllSaldo(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseSaldoAll, error) {
	_, err := s.service.RestoreAllSaldo(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseSaldoAll{
		Status:  "success",
		Message: "Successfully restore all saldo",
	}, nil
}

func (s *saldoCommandHandleGrpc) DeleteAllSaldoPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseSaldoAll, error) {
	_, err := s.service.DeleteAllSaldoPermanent(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseSaldoAll{
		Status:  "success",
		Message: "delete saldo permanent",
	}, nil
}
