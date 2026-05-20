package handler

import (
	"context"
	"fmt"
	"time"

	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/withdraw"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/errors"
	"github.com/MamangRust/monolith-graphql-payment-gateway-withdraw/service"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"github.com/go-playground/validator/v10"
)

type withdrawCommandHandleGrpc struct {
	pb.UnimplementedWithdrawCommandServiceServer

	withdrawCommand service.WithdrawCommandService
}

func NewWithdrawCommandHandleGrpc(
	withdrawCommand service.WithdrawCommandService,
) WithdrawCommandHandlerGrpc {
	return &withdrawCommandHandleGrpc{
		withdrawCommand: withdrawCommand,
	}
}

func (w *withdrawCommandHandleGrpc) CreateWithdraw(ctx context.Context, req *pb.CreateWithdrawRequest) (*pb.ApiResponseWithdraw, error) {
	request := &requests.CreateWithdrawRequest{
		CardNumber:     req.CardNumber,
		WithdrawAmount: int(req.WithdrawAmount),
		WithdrawTime:   req.WithdrawTime.AsTime(),
	}

	if err := request.Validate(); err != nil {
		validations := w.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	withdraw, err := w.withdrawCommand.Create(ctx, request)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseWithdraw{
		Status:  "success",
		Message: "Successfully created withdraw",
		Data: &pb.WithdrawResponse{
			WithdrawId:     int32(withdraw.WithdrawID),
			WithdrawNo:     withdraw.WithdrawNo.String(),
			CardNumber:     withdraw.CardNumber,
			WithdrawAmount: int32(withdraw.WithdrawAmount),
			WithdrawTime:   withdraw.WithdrawTime.Format(time.RFC3339),
			CreatedAt:      withdraw.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:      withdraw.UpdatedAt.Time.Format(time.RFC3339),
		},
	}, nil
}

func (w *withdrawCommandHandleGrpc) UpdateWithdraw(ctx context.Context, req *pb.UpdateWithdrawRequest) (*pb.ApiResponseWithdraw, error) {
	id := int(req.GetWithdrawId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "withdraw_id",
				Message: "Withdraw ID is required",
			},
		}))
	}

	request := &requests.UpdateWithdrawRequest{
		WithdrawID:     &id,
		CardNumber:     req.CardNumber,
		WithdrawAmount: int(req.WithdrawAmount),
		WithdrawTime:   req.WithdrawTime.AsTime(),
	}

	if err := request.Validate(); err != nil {
		validations := w.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	withdraw, err := w.withdrawCommand.Update(ctx, request)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseWithdraw{
		Status:  "success",
		Message: "Successfully updated withdraw",
		Data: &pb.WithdrawResponse{
			WithdrawId:     int32(withdraw.WithdrawID),
			WithdrawNo:     withdraw.WithdrawNo.String(),
			CardNumber:     withdraw.CardNumber,
			WithdrawAmount: int32(withdraw.WithdrawAmount),
			WithdrawTime:   withdraw.WithdrawTime.Format(time.RFC3339),
			CreatedAt:      withdraw.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:      withdraw.UpdatedAt.Time.Format(time.RFC3339),
		},
	}, nil
}

func (w *withdrawCommandHandleGrpc) TrashedWithdraw(ctx context.Context, req *pb.FindByIdWithdrawRequest) (*pb.ApIResponseWithdrawDeleteAt, error) {
	id := int(req.GetWithdrawId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "withdraw_id",
				Message: "Withdraw ID is required",
			},
		}))
	}

	withdraw, err := w.withdrawCommand.TrashedWithdraw(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApIResponseWithdrawDeleteAt{
		Status:  "success",
		Message: "Successfully trashed withdraw",
		Data: &pb.WithdrawResponseDeleteAt{
			WithdrawId:     int32(withdraw.WithdrawID),
			WithdrawNo:     withdraw.WithdrawNo.String(),
			CardNumber:     withdraw.CardNumber,
			WithdrawAmount: int32(withdraw.WithdrawAmount),
			WithdrawTime:   withdraw.WithdrawTime.Format(time.RFC3339),
			CreatedAt:      withdraw.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:      withdraw.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:      &wrapperspb.StringValue{Value: withdraw.DeletedAt.Time.Format(time.RFC3339)},
		},
	}, nil
}

func (w *withdrawCommandHandleGrpc) RestoreWithdraw(ctx context.Context, req *pb.FindByIdWithdrawRequest) (*pb.ApIResponseWithdrawDeleteAt, error) {
	id := int(req.GetWithdrawId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "withdraw_id",
				Message: "Withdraw ID is required",
			},
		}))
	}

	withdraw, err := w.withdrawCommand.RestoreWithdraw(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApIResponseWithdrawDeleteAt{
		Status:  "success",
		Message: "Successfully restored withdraw",
		Data: &pb.WithdrawResponseDeleteAt{
			WithdrawId:     int32(withdraw.WithdrawID),
			WithdrawNo:     withdraw.WithdrawNo.String(),
			CardNumber:     withdraw.CardNumber,
			WithdrawAmount: int32(withdraw.WithdrawAmount),
			WithdrawTime:   withdraw.WithdrawTime.Format(time.RFC3339),
			CreatedAt:      withdraw.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:      withdraw.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:      &wrapperspb.StringValue{Value: withdraw.DeletedAt.Time.Format(time.RFC3339)},
		},
	}, nil
}

func (w *withdrawCommandHandleGrpc) DeleteWithdrawPermanent(ctx context.Context, req *pb.FindByIdWithdrawRequest) (*pb.ApiResponseWithdrawDelete, error) {
	id := int(req.GetWithdrawId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "withdraw_id",
				Message: "Withdraw ID is required",
			},
		}))
	}

	_, err := w.withdrawCommand.DeleteWithdrawPermanent(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseWithdrawDelete{
		Status:  "success",
		Message: "Successfully deleted withdraw permanently",
	}, nil
}

func (s *withdrawCommandHandleGrpc) RestoreAllWithdraw(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseWithdrawAll, error) {
	_, err := s.withdrawCommand.RestoreAllWithdraw(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseWithdrawAll{
		Status:  "success",
		Message: "Successfully restore all withdraw",
	}, nil
}

func (s *withdrawCommandHandleGrpc) DeleteAllWithdrawPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseWithdrawAll, error) {
	_, err := s.withdrawCommand.DeleteAllWithdrawPermanent(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseWithdrawAll{
		Status:  "success",
		Message: "Successfully delete withdraw permanent",
	}, nil
}

func (w *withdrawCommandHandleGrpc) parseValidationErrors(err error) []errors.ValidationError {
	var validationErrs []errors.ValidationError

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			validationErrs = append(validationErrs, errors.ValidationError{
				Field:   fe.Field(),
				Message: w.getValidationMessage(fe),
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

func (w *withdrawCommandHandleGrpc) getValidationMessage(fe validator.FieldError) string {
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
