package handler

import (
	"context"
	"fmt"

	pb "github.com/MamangRust/monolith-graphql-payment-gateway-pb/role"
	"github.com/MamangRust/monolith-graphql-payment-gateway-role/service"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"github.com/go-playground/validator/v10"
)

type roleCommandHandleGrpc struct {
	pb.UnimplementedRoleCommandServiceServer
	roleCommand service.RoleCommandService
}

func NewRoleCommandHandleGrpc(roleCommand service.RoleCommandService) RoleCommandHandlerGrpc {
	return &roleCommandHandleGrpc{
		roleCommand: roleCommand,
	}
}

func (s *roleCommandHandleGrpc) CreateRole(ctx context.Context, reqPb *pb.CreateRoleRequest) (*pb.ApiResponseRole, error) {
	req := &requests.CreateRoleRequest{
		Name: reqPb.Name,
	}

	if err := req.Validate(); err != nil {
		validations := s.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	role, err := s.roleCommand.CreateRole(ctx, req)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoRole := &pb.RoleResponse{
		Id:        int32(role.RoleID),
		Name:      role.RoleName,
		CreatedAt: role.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: role.UpdatedAt.Time.Format("2006-01-02"),
	}

	return &pb.ApiResponseRole{
		Status:  "success",
		Message: "Successfully created role",
		Data:    protoRole,
	}, nil
}

func (s *roleCommandHandleGrpc) UpdateRole(ctx context.Context, reqPb *pb.UpdateRoleRequest) (*pb.ApiResponseRole, error) {
	roleID := int(reqPb.GetId())

	if roleID == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "id",
				Message: "Role ID is required",
			},
		}))
	}

	name := reqPb.GetName()

	req := &requests.UpdateRoleRequest{
		ID:   &roleID,
		Name: name,
	}

	if err := req.Validate(); err != nil {
		validations := s.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	role, err := s.roleCommand.UpdateRole(ctx, req)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoRole := &pb.RoleResponse{
		Id:        int32(role.RoleID),
		Name:      role.RoleName,
		CreatedAt: role.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: role.UpdatedAt.Time.Format("2006-01-02"),
	}

	return &pb.ApiResponseRole{
		Status:  "success",
		Message: "Successfully updated role",
		Data:    protoRole,
	}, nil
}

func (s *roleCommandHandleGrpc) parseValidationErrors(err error) []errors.ValidationError {
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

func (s *roleCommandHandleGrpc) getValidationMessage(fe validator.FieldError) string {
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

func (s *roleCommandHandleGrpc) TrashedRole(ctx context.Context, req *pb.FindByIdRoleRequest) (*pb.ApiResponseRoleDeleteAt, error) {
	roleID := int(req.GetRoleId())

	if roleID == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "role_id",
				Message: "Role ID is required",
			},
		}))
	}

	role, err := s.roleCommand.TrashedRole(ctx, roleID)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoRole := &pb.RoleResponseDeleteAt{
		Id:        int32(role.RoleID),
		Name:      role.RoleName,
		CreatedAt: role.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: role.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt: wrapperspb.String(role.DeletedAt.Time.Format("2006-01-02")),
	}

	return &pb.ApiResponseRoleDeleteAt{
		Status:  "success",
		Message: "Successfully trashed role",
		Data:    protoRole,
	}, nil
}

func (s *roleCommandHandleGrpc) RestoreRole(ctx context.Context, req *pb.FindByIdRoleRequest) (*pb.ApiResponseRoleDeleteAt, error) {
	roleID := int(req.GetRoleId())

	if roleID == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "role_id",
				Message: "Role ID is required",
			},
		}))
	}

	role, err := s.roleCommand.RestoreRole(ctx, roleID)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoRole := &pb.RoleResponseDeleteAt{
		Id:        int32(role.RoleID),
		Name:      role.RoleName,
		CreatedAt: role.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: role.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt: wrapperspb.String(role.DeletedAt.Time.Format("2006-01-02")),
	}

	return &pb.ApiResponseRoleDeleteAt{
		Status:  "success",
		Message: "Successfully restored role",
		Data:    protoRole,
	}, nil
}

func (s *roleCommandHandleGrpc) DeleteRolePermanent(ctx context.Context, req *pb.FindByIdRoleRequest) (*pb.ApiResponseRoleDelete, error) {
	roleID := int(req.GetRoleId())

	if roleID == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "role_id",
				Message: "Role ID is required",
			},
		}))
	}

	_, err := s.roleCommand.DeleteRolePermanent(ctx, roleID)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseRoleDelete{
		Status:  "success",
		Message: "Successfully deleted role permanently",
	}, nil
}

func (s *roleCommandHandleGrpc) RestoreAllRole(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseRoleAll, error) {
	_, err := s.roleCommand.RestoreAllRole(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseRoleAll{
		Status:  "success",
		Message: "Successfully restore all roles",
	}, nil
}

func (s *roleCommandHandleGrpc) DeleteAllRolePermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseRoleAll, error) {
	_, err := s.roleCommand.DeleteAllRolePermanent(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseRoleAll{
		Status:  "success",
		Message: "delete all roles permanent",
	}, nil
}
