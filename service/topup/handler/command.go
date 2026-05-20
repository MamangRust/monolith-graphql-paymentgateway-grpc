package handler

import (
	"context"
	"fmt"
	"time"

	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/topup"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/errors"
	"github.com/MamangRust/monolith-graphql-payment-gateway-topup/service"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"github.com/go-playground/validator/v10"
)

type topupCommandHandleGrpc struct {
	pb.UnimplementedTopupCommandServiceServer

	service service.TopupCommandService
}

func NewTopupCommandHandleGrpc(service service.TopupCommandService) TopupCommandHandleGrpc {
	return &topupCommandHandleGrpc{
		service: service,
	}
}

func (s *topupCommandHandleGrpc) CreateTopup(ctx context.Context, req *pb.CreateTopupRequest) (*pb.ApiResponseTopup, error) {
	request := requests.CreateTopupRequest{
		CardNumber:  req.GetCardNumber(),
		TopupAmount: int(req.GetTopupAmount()),
		TopupMethod: req.GetTopupMethod(),
	}

	if err := request.Validate(); err != nil {
		validations := s.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	res, err := s.service.CreateTopup(ctx, &request)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoTopup := &pb.TopupResponse{
		Id:          int32(res.TopupID),
		CardNumber:  res.CardNumber,
		TopupNo:     res.TopupNo.String(),
		TopupAmount: int32(res.TopupAmount),
		TopupMethod: res.TopupMethod,
		TopupTime:   res.TopupTime.Format(time.RFC3339),
		CreatedAt:   res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:   res.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseTopup{
		Status:  "success",
		Message: "Successfully created topup",
		Data:    protoTopup,
	}, nil
}

func (s *topupCommandHandleGrpc) UpdateTopup(ctx context.Context, req *pb.UpdateTopupRequest) (*pb.ApiResponseTopup, error) {
	id := int(req.GetTopupId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "topup_id",
				Message: "Topup ID is required",
			},
		}))
	}

	request := requests.UpdateTopupRequest{
		TopupID:     &id,
		CardNumber:  req.GetCardNumber(),
		TopupAmount: int(req.GetTopupAmount()),
		TopupMethod: req.GetTopupMethod(),
	}

	if err := request.Validate(); err != nil {
		validations := s.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	res, err := s.service.UpdateTopup(ctx, &request)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoTopup := &pb.TopupResponse{
		Id:          int32(res.TopupID),
		CardNumber:  res.CardNumber,
		TopupNo:     res.TopupNo.String(),
		TopupAmount: int32(res.TopupAmount),
		TopupMethod: res.TopupMethod,
		TopupTime:   res.TopupTime.Format(time.RFC3339),
		CreatedAt:   res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:   res.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseTopup{
		Status:  "success",
		Message: "Successfully updated topup",
		Data:    protoTopup,
	}, nil
}

func (s *topupCommandHandleGrpc) parseValidationErrors(err error) []errors.ValidationError {
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

func (s *topupCommandHandleGrpc) getValidationMessage(fe validator.FieldError) string {
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

func (s *topupCommandHandleGrpc) TrashedTopup(ctx context.Context, req *pb.FindByIdTopupRequest) (*pb.ApiResponseTopupDeleteAt, error) {
	id := int(req.GetTopupId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "topup_id",
				Message: "Topup ID is required",
			},
		}))
	}

	res, err := s.service.TrashedTopup(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoTopup := &pb.TopupResponseDeleteAt{
		Id:          int32(res.TopupID),
		CardNumber:  res.CardNumber,
		TopupNo:     res.TopupNo.String(),
		TopupAmount: int32(res.TopupAmount),
		TopupMethod: res.TopupMethod,
		TopupTime:   res.TopupTime.Format(time.RFC3339),
		CreatedAt:   res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:   res.UpdatedAt.Time.Format(time.RFC3339),
		DeletedAt:   wrapperspb.String(res.DeletedAt.Time.Format(time.RFC3339)),
	}

	return &pb.ApiResponseTopupDeleteAt{
		Status:  "success",
		Message: "Successfully trashed topup",
		Data:    protoTopup,
	}, nil
}

func (s *topupCommandHandleGrpc) RestoreTopup(ctx context.Context, req *pb.FindByIdTopupRequest) (*pb.ApiResponseTopupDeleteAt, error) {
	id := int(req.GetTopupId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "topup_id",
				Message: "Topup ID is required",
			},
		}))
	}

	res, err := s.service.RestoreTopup(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoTopup := &pb.TopupResponseDeleteAt{
		Id:          int32(res.TopupID),
		CardNumber:  res.CardNumber,
		TopupNo:     res.TopupNo.String(),
		TopupAmount: int32(res.TopupAmount),
		TopupMethod: res.TopupMethod,
		TopupTime:   res.TopupTime.Format(time.RFC3339),
		CreatedAt:   res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:   res.UpdatedAt.Time.Format(time.RFC3339),
		DeletedAt:   wrapperspb.String(res.DeletedAt.Time.Format(time.RFC3339)),
	}

	return &pb.ApiResponseTopupDeleteAt{
		Status:  "success",
		Message: "Successfully restored topup",
		Data:    protoTopup,
	}, nil
}

func (s *topupCommandHandleGrpc) DeleteTopupPermanent(ctx context.Context, req *pb.FindByIdTopupRequest) (*pb.ApiResponseTopupDelete, error) {
	id := int(req.GetTopupId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "topup_id",
				Message: "Topup ID is required",
			},
		}))
	}

	_, err := s.service.DeleteTopupPermanent(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTopupDelete{
		Status:  "success",
		Message: "Successfully deleted topup permanently",
	}, nil
}

func (s *topupCommandHandleGrpc) RestoreAllTopup(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTopupAll, error) {
	_, err := s.service.RestoreAllTopup(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTopupAll{
		Status:  "success",
		Message: "Successfully restore all topup",
	}, nil
}

func (s *topupCommandHandleGrpc) DeleteAllTopupPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTopupAll, error) {
	_, err := s.service.DeleteAllTopupPermanent(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTopupAll{
		Status:  "success",
		Message: "Successfully delete topup permanent",
	}, nil
}
