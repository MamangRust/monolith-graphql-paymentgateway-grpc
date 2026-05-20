package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/monolith-graphql-payment-gateway-merchant/service"
	pbmerchant "github.com/MamangRust/monolith-graphql-payment-gateway-pb/merchant"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"github.com/go-playground/validator/v10"
)

type merchantCommandHandleGrpc struct {
	pbmerchant.UnimplementedMerchantCommandServiceServer

	merchantCommand service.MerchantCommandService
}

func NewMerchantCommandHandleGrpc(merchantCommand service.MerchantCommandService) MerchantCommandHandleGrpc {
	return &merchantCommandHandleGrpc{
		merchantCommand: merchantCommand,
	}
}

func (s *merchantCommandHandleGrpc) CreateMerchant(ctx context.Context, req *pbmerchant.CreateMerchantRequest) (*pbmerchant.ApiResponseMerchant, error) {
	request := requests.CreateMerchantRequest{
		Name:   req.GetName(),
		UserID: int(req.GetUserId()),
	}

	if err := request.Validate(); err != nil {
		validations := s.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	merchant, err := s.merchantCommand.CreateMerchant(ctx, &request)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoMerchant := &pbmerchant.MerchantResponse{
		Id:        int32(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		Status:    merchant.Status,
		UserId:    int32(merchant.UserID),
		CreatedAt: merchant.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: merchant.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pbmerchant.ApiResponseMerchant{
		Status:  "success",
		Message: "Successfully created merchant",
		Data:    protoMerchant,
	}, nil
}

func (s *merchantCommandHandleGrpc) UpdateMerchant(ctx context.Context, req *pbmerchant.UpdateMerchantRequest) (*pbmerchant.ApiResponseMerchant, error) {
	id := int(req.GetMerchantId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "merchant_id",
				Message: "Merchant ID is required",
			},
		}))
	}

	request := requests.UpdateMerchantRequest{
		MerchantID: &id,
		Name:       req.GetName(),
		UserID:     int(req.GetUserId()),
		Status:     req.GetStatus(),
	}

	if err := request.Validate(); err != nil {
		validations := s.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	merchant, err := s.merchantCommand.UpdateMerchant(ctx, &request)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoMerchant := &pbmerchant.MerchantResponse{
		Id:        int32(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		Status:    merchant.Status,
		UserId:    int32(merchant.UserID),
		CreatedAt: merchant.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: merchant.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pbmerchant.ApiResponseMerchant{
		Status:  "success",
		Message: "Successfully updated merchant",
		Data:    protoMerchant,
	}, nil
}

func (s *merchantCommandHandleGrpc) parseValidationErrors(err error) []errors.ValidationError {
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

func (s *merchantCommandHandleGrpc) getValidationMessage(fe validator.FieldError) string {
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

func (s *merchantCommandHandleGrpc) TrashedMerchant(ctx context.Context, req *pbmerchant.FindByIdMerchantRequest) (*pbmerchant.ApiResponseMerchantDeleteAt, error) {
	id := int(req.GetMerchantId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "merchant_id",
				Message: "Merchant ID is required",
			},
		}))
	}

	merchant, err := s.merchantCommand.TrashedMerchant(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoMerchant := &pbmerchant.MerchantResponseDeleteAt{
		Id:        int32(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		Status:    merchant.Status,
		UserId:    int32(merchant.UserID),
		CreatedAt: merchant.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: merchant.UpdatedAt.Time.Format(time.RFC3339),
		DeletedAt: wrapperspb.String(merchant.DeletedAt.Time.Format(time.RFC3339)),
	}

	return &pbmerchant.ApiResponseMerchantDeleteAt{
		Status:  "success",
		Message: "Successfully trashed merchant",
		Data:    protoMerchant,
	}, nil
}

func (s *merchantCommandHandleGrpc) RestoreMerchant(ctx context.Context, req *pbmerchant.FindByIdMerchantRequest) (*pbmerchant.ApiResponseMerchantDeleteAt, error) {
	id := int(req.GetMerchantId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "merchant_id",
				Message: "Merchant ID is required",
			},
		}))
	}

	merchant, err := s.merchantCommand.RestoreMerchant(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoMerchant := &pbmerchant.MerchantResponseDeleteAt{
		Id:        int32(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		Status:    merchant.Status,
		UserId:    int32(merchant.UserID),
		CreatedAt: merchant.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: merchant.UpdatedAt.Time.Format(time.RFC3339),
		DeletedAt: wrapperspb.String(merchant.DeletedAt.Time.Format(time.RFC3339)),
	}

	return &pbmerchant.ApiResponseMerchantDeleteAt{
		Status:  "success",
		Message: "Successfully restored merchant",
		Data:    protoMerchant,
	}, nil
}

func (s *merchantCommandHandleGrpc) DeleteMerchantPermanent(ctx context.Context, req *pbmerchant.FindByIdMerchantRequest) (*pbmerchant.ApiResponseMerchantDelete, error) {
	id := int(req.GetMerchantId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "merchant_id",
				Message: "Merchant ID is required",
			},
		}))
	}

	_, err := s.merchantCommand.DeleteMerchantPermanent(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pbmerchant.ApiResponseMerchantDelete{
		Status:  "success",
		Message: "Successfully deleted merchant",
	}, nil
}

func (s *merchantCommandHandleGrpc) RestoreAllMerchant(ctx context.Context, _ *emptypb.Empty) (*pbmerchant.ApiResponseMerchantAll, error) {
	_, err := s.merchantCommand.RestoreAllMerchant(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pbmerchant.ApiResponseMerchantAll{
		Status:  "success",
		Message: "Successfully restore all merchant",
	}, nil
}

func (s *merchantCommandHandleGrpc) DeleteAllMerchantPermanent(ctx context.Context, _ *emptypb.Empty) (*pbmerchant.ApiResponseMerchantAll, error) {
	_, err := s.merchantCommand.DeleteAllMerchantPermanent(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pbmerchant.ApiResponseMerchantAll{
		Status:  "success",
		Message: "Successfully delete all merchant",
	}, nil
}
