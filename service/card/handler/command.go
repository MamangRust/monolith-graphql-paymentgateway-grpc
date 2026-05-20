package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/monolith-graphql-payment-gateway-card/service"
	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/card"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"github.com/go-playground/validator/v10"
)

type cardCommandService struct {
	pb.UnimplementedCardCommandServiceServer

	cardCommand service.CardCommandService
}

func NewCardCommandHandleGrpc(cardCommand service.CardCommandService) CardCommandService {
	return &cardCommandService{
		cardCommand: cardCommand,
	}
}

func (s *cardCommandService) CreateCard(ctx context.Context, req *pb.CreateCardRequest) (*pb.ApiResponseCard, error) {
	request := requests.CreateCardRequest{
		UserID:       int(req.UserId),
		CardType:     req.CardType,
		ExpireDate:   req.ExpireDate.AsTime(),
		CVV:          req.Cvv,
		CardProvider: req.CardProvider,
	}

	if err := request.Validate(); err != nil {
		validations := s.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	res, err := s.cardCommand.CreateCard(ctx, &request)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoCard := &pb.CardResponse{
		Id:         int32(res.CardID),
		UserId:     int32(res.UserID),
		CardNumber: res.CardNumber,
		CardType:   res.CardType,
		Cvv:        res.Cvv,
		ExpireDate: res.ExpireDate.Time.Format(time.RFC3339),
		CreatedAt:  res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:  res.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseCard{
		Status:  "success",
		Message: "Successfully created card",
		Data:    protoCard,
	}, nil
}

func (s *cardCommandService) UpdateCard(ctx context.Context, req *pb.UpdateCardRequest) (*pb.ApiResponseCard, error) {
	request := requests.UpdateCardRequest{
		CardID:       int(req.CardId),
		UserID:       int(req.UserId),
		CardType:     req.CardType,
		ExpireDate:   req.ExpireDate.AsTime(),
		CVV:          req.Cvv,
		CardProvider: req.CardProvider,
	}

	if err := request.Validate(); err != nil {
		validations := s.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	res, err := s.cardCommand.UpdateCard(ctx, &request)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoCard := &pb.CardResponse{
		Id:         int32(res.CardID),
		UserId:     int32(res.UserID),
		CardNumber: res.CardNumber,
		CardType:   res.CardType,
		Cvv:        res.Cvv,
		ExpireDate: res.ExpireDate.Time.Format(time.RFC3339),
		CreatedAt:  res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:  res.UpdatedAt.Time.Format(time.RFC3339),
	}

	return &pb.ApiResponseCard{
		Status:  "success",
		Message: "Successfully updated card",
		Data:    protoCard,
	}, nil
}

func (s *cardCommandService) parseValidationErrors(err error) []errors.ValidationError {
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

func (s *cardCommandService) getValidationMessage(fe validator.FieldError) string {
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

func (s *cardCommandService) TrashedCard(ctx context.Context, req *pb.FindByIdCardRequest) (*pb.ApiResponseCardDeleteAt, error) {
	id := int(req.GetCardId())
	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "card_id",
				Message: "Card ID is required",
			},
		}))
	}

	res, err := s.cardCommand.TrashedCard(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoCard := &pb.CardResponseDeleteAt{
		Id:         int32(res.CardID),
		UserId:     int32(res.UserID),
		CardNumber: res.CardNumber,
		CardType:   res.CardType,
		Cvv:        res.Cvv,
		ExpireDate: res.ExpireDate.Time.Format(time.RFC3339),
		CreatedAt:  res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:  res.UpdatedAt.Time.Format(time.RFC3339),
		DeletedAt:  wrapperspb.String(res.DeletedAt.Time.Format(time.RFC3339)),
	}

	return &pb.ApiResponseCardDeleteAt{
		Status:  "success",
		Message: "Successfully trashed card",
		Data:    protoCard,
	}, nil
}

func (s *cardCommandService) RestoreCard(ctx context.Context, req *pb.FindByIdCardRequest) (*pb.ApiResponseCardDeleteAt, error) {
	id := int(req.GetCardId())
	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "card_id",
				Message: "Card ID is required",
			},
		}))
	}

	res, err := s.cardCommand.RestoreCard(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoCard := &pb.CardResponseDeleteAt{
		Id:         int32(res.CardID),
		UserId:     int32(res.UserID),
		CardNumber: res.CardNumber,
		CardType:   res.CardType,
		Cvv:        res.Cvv,
		ExpireDate: res.ExpireDate.Time.Format(time.RFC3339),
		CreatedAt:  res.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:  res.UpdatedAt.Time.Format(time.RFC3339),
		DeletedAt:  wrapperspb.String(res.DeletedAt.Time.Format(time.RFC3339)),
	}

	return &pb.ApiResponseCardDeleteAt{
		Status:  "success",
		Message: "Successfully restored card",
		Data:    protoCard,
	}, nil
}

func (s *cardCommandService) DeleteCardPermanent(ctx context.Context, req *pb.FindByIdCardRequest) (*pb.ApiResponseCardDelete, error) {
	id := int(req.GetCardId())
	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "card_id",
				Message: "Card ID is required",
			},
		}))
	}

	_, err := s.cardCommand.DeleteCardPermanent(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseCardDelete{
		Status:  "success",
		Message: "Successfully deleted card",
	}, nil
}

func (s *cardCommandService) RestoreAllCard(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCardAll, error) {
	_, err := s.cardCommand.RestoreAllCard(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseCardAll{
		Status:  "success",
		Message: "Successfully restore card",
	}, nil
}

func (s *cardCommandService) DeleteAllCardPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCardAll, error) {
	_, err := s.cardCommand.DeleteAllCardPermanent(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseCardAll{
		Status:  "success",
		Message: "Successfully delete card permanent",
	}, nil
}
