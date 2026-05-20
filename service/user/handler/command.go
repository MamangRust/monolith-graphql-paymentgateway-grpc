package handler

import (
	"context"
	"fmt"
	"time"

	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/user"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/errors"
	"github.com/MamangRust/monolith-graphql-payment-gateway-user/service"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"github.com/go-playground/validator/v10"
)

type userCommandHandleGrpc struct {
	pb.UnimplementedUserCommandServiceServer

	userCommandService service.UserCommandService
}

func NewUserCommandHandleGrpc(query service.UserCommandService) UserCommandHandleGrpc {
	return &userCommandHandleGrpc{
		userCommandService: query,
	}
}

func (s *userCommandHandleGrpc) Create(ctx context.Context, request *pb.CreateUserRequest) (*pb.ApiResponseUser, error) {
	req := &requests.CreateUserRequest{
		FirstName:       request.GetFirstname(),
		LastName:        request.GetLastname(),
		Email:           request.GetEmail(),
		Password:        request.GetPassword(),
		ConfirmPassword: request.GetConfirmPassword(),
	}

	if err := req.Validate(); err != nil {
		validations := s.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	user, err := s.userCommandService.CreateUser(ctx, req)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseUser{
		Status:  "success",
		Message: "Successfully created user",
		Data: &pb.UserResponse{
			Id:        int32(user.UserID),
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Time.Format(time.RFC3339),
		},
	}, nil
}

func (s *userCommandHandleGrpc) Update(ctx context.Context, request *pb.UpdateUserRequest) (*pb.ApiResponseUser, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "id",
				Message: "User ID is required",
			},
		}))
	}

	req := &requests.UpdateUserRequest{
		UserID:          &id,
		FirstName:       request.GetFirstname(),
		LastName:        request.GetLastname(),
		Email:           request.GetEmail(),
		Password:        request.GetPassword(),
		ConfirmPassword: request.GetConfirmPassword(),
	}

	if err := req.Validate(); err != nil {
		validations := s.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	user, err := s.userCommandService.UpdateUser(ctx, req)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseUser{
		Status:  "success",
		Message: "Successfully updated user",
		Data: &pb.UserResponse{
			Id:        int32(user.UserID),
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Time.Format(time.RFC3339),
		},
	}, nil
}

func (s *userCommandHandleGrpc) TrashedUser(ctx context.Context, request *pb.FindByIdUserRequest) (*pb.ApiResponseUserDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "user_id",
				Message: "User ID is required",
			},
		}))
	}

	user, err := s.userCommandService.TrashedUser(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseUserDeleteAt{
		Status:  "success",
		Message: "Successfully trashed user",
		Data: &pb.UserResponseDeleteAt{
			Id:        int32(user.UserID),
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt: &wrapperspb.StringValue{Value: user.DeletedAt.Time.Format(time.RFC3339)},
		},
	}, nil
}

func (s *userCommandHandleGrpc) RestoreUser(ctx context.Context, request *pb.FindByIdUserRequest) (*pb.ApiResponseUserDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "user_id",
				Message: "User ID is required",
			},
		}))
	}

	user, err := s.userCommandService.RestoreUser(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseUserDeleteAt{
		Status:  "success",
		Message: "Successfully restored user",
		Data: &pb.UserResponseDeleteAt{
			Id:        int32(user.UserID),
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Time.Format(time.RFC3339),
			DeletedAt: &wrapperspb.StringValue{Value: user.DeletedAt.Time.Format(time.RFC3339)},
		},
	}, nil
}

func (s *userCommandHandleGrpc) DeleteUserPermanent(ctx context.Context, request *pb.FindByIdUserRequest) (*pb.ApiResponseUserDelete, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "id",
				Message: "User ID is required",
			},
		}))
	}

	_, err := s.userCommandService.DeleteUserPermanent(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseUserDelete{
		Status:  "success",
		Message: "Successfully deleted user permanently",
	}, nil
}

func (s *userCommandHandleGrpc) RestoreAllUser(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseUserAll, error) {
	_, err := s.userCommandService.RestoreAllUser(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseUserAll{
		Status:  "success",
		Message: "Successfully restored all users",
	}, nil
}

func (s *userCommandHandleGrpc) DeleteAllUserPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseUserAll, error) {
	_, err := s.userCommandService.DeleteAllUserPermanent(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseUserAll{
		Status:  "success",
		Message: "Successfully deleted all users permanently",
	}, nil
}

func (s *userCommandHandleGrpc) parseValidationErrors(err error) []errors.ValidationError {
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

func (s *userCommandHandleGrpc) getValidationMessage(fe validator.FieldError) string {
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
