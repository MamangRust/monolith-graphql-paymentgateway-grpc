package handler

import (
	"context"
	"fmt"
	"time"

	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/transfer"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/errors"
	"github.com/MamangRust/monolith-graphql-payment-gateway-transfer/service"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"github.com/go-playground/validator/v10"
)

type transferCommandHandleGrpc struct {
	pb.UnimplementedTransferCommandServiceServer

	transferCommandService service.TransferCommandService
}

func NewTransferCommandHandler(service service.TransferCommandService) TransferCommandHandleGrpc {
	return &transferCommandHandleGrpc{
		transferCommandService: service,
	}
}

func (s *transferCommandHandleGrpc) CreateTransfer(ctx context.Context, request *pb.CreateTransferRequest) (*pb.ApiResponseTransfer, error) {
	req := requests.CreateTransferRequest{
		TransferFrom:   request.GetTransferFrom(),
		TransferTo:     request.GetTransferTo(),
		TransferAmount: int(request.GetTransferAmount()),
	}

	if err := req.Validate(); err != nil {
		validations := s.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	transfer, err := s.transferCommandService.CreateTransaction(ctx, &req)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransfer{
		Status:  "success",
		Message: "Successfully created transfer",
		Data: &pb.TransferResponse{
			Id:             int32(transfer.TransferID),
			TransferNo:     transfer.TransferNo.String(),
			TransferFrom:   transfer.TransferFrom,
			TransferTo:     transfer.TransferTo,
			TransferAmount: int32(transfer.TransferAmount),
			TransferTime:   transfer.TransferTime.Format(time.RFC3339),
			CreatedAt:      transfer.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:      transfer.UpdatedAt.Time.Format(time.RFC3339),
		},
	}, nil
}

func (s *transferCommandHandleGrpc) UpdateTransfer(ctx context.Context, request *pb.UpdateTransferRequest) (*pb.ApiResponseTransfer, error) {
	id := int(request.GetTransferId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "transfer_id",
				Message: "Transfer ID is required",
			},
		}))
	}

	req := requests.UpdateTransferRequest{
		TransferID:     &id,
		TransferFrom:   request.GetTransferFrom(),
		TransferTo:     request.GetTransferTo(),
		TransferAmount: int(request.GetTransferAmount()),
	}

	if err := req.Validate(); err != nil {
		validations := s.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	transfer, err := s.transferCommandService.UpdateTransaction(ctx, &req)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransfer{
		Status:  "success",
		Message: "Successfully updated transfer",
		Data: &pb.TransferResponse{
			Id:             int32(transfer.TransferID),
			TransferNo:     transfer.TransferNo.String(),
			TransferFrom:   transfer.TransferFrom,
			TransferTo:     transfer.TransferTo,
			TransferAmount: int32(transfer.TransferAmount),
			TransferTime:   transfer.TransferTime.Format(time.RFC3339),
			CreatedAt:      transfer.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:      transfer.UpdatedAt.Time.Format(time.RFC3339),
		},
	}, nil
}

func (s *transferCommandHandleGrpc) TrashedTransfer(ctx context.Context, request *pb.FindByIdTransferRequest) (*pb.ApIResponseTransferDeleteAt, error) {
	id := int(request.GetTransferId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "transfer_id",
				Message: "Transfer ID is required",
			},
		}))
	}

	transfer, err := s.transferCommandService.TrashedTransfer(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApIResponseTransferDeleteAt{
		Status:  "success",
		Message: "Successfully trashed transfer",
		Data: &pb.TransferResponseDeleteAt{
			Id:             int32(transfer.TransferID),
			TransferNo:     transfer.TransferNo.String(),
			TransferFrom:   transfer.TransferFrom,
			TransferTo:     transfer.TransferTo,
			TransferAmount: int32(transfer.TransferAmount),
			TransferTime:   transfer.TransferTime.Format(time.RFC3339),
			CreatedAt:      transfer.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:      transfer.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:      &wrapperspb.StringValue{Value: transfer.DeletedAt.Time.Format(time.RFC3339)},
		},
	}, nil
}

func (s *transferCommandHandleGrpc) RestoreTransfer(ctx context.Context, request *pb.FindByIdTransferRequest) (*pb.ApIResponseTransferDeleteAt, error) {
	id := int(request.GetTransferId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "transfer_id",
				Message: "Transfer ID is required",
			},
		}))
	}

	transfer, err := s.transferCommandService.RestoreTransfer(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApIResponseTransferDeleteAt{
		Status:  "success",
		Message: "Successfully restored transfer",
		Data: &pb.TransferResponseDeleteAt{
			Id:             int32(transfer.TransferID),
			TransferNo:     transfer.TransferNo.String(),
			TransferFrom:   transfer.TransferFrom,
			TransferTo:     transfer.TransferTo,
			TransferAmount: int32(transfer.TransferAmount),
			TransferTime:   transfer.TransferTime.Format(time.RFC3339),
			CreatedAt:      transfer.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:      transfer.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt:      &wrapperspb.StringValue{Value: transfer.DeletedAt.Time.Format(time.RFC3339)},
		},
	}, nil
}

func (s *transferCommandHandleGrpc) DeleteTransferPermanent(ctx context.Context, request *pb.FindByIdTransferRequest) (*pb.ApiResponseTransferDelete, error) {
	id := int(request.GetTransferId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "transfer_id",
				Message: "Transfer ID is required",
			},
		}))
	}

	_, err := s.transferCommandService.DeleteTransferPermanent(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransferDelete{
		Status:  "success",
		Message: "Successfully deleted transfer permanently",
	}, nil
}

func (s *transferCommandHandleGrpc) RestoreAllTransfer(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTransferAll, error) {
	_, err := s.transferCommandService.RestoreAllTransfer(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransferAll{
		Status:  "success",
		Message: "Successfully restored all transfers",
	}, nil
}

func (s *transferCommandHandleGrpc) DeleteAllTransferPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTransferAll, error) {
	_, err := s.transferCommandService.DeleteAllTransferPermanent(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransferAll{
		Status:  "success",
		Message: "Successfully deleted all transfers permanently",
	}, nil
}

func (s *transferCommandHandleGrpc) parseValidationErrors(err error) []errors.ValidationError {
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

func (s *transferCommandHandleGrpc) getValidationMessage(fe validator.FieldError) string {
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
