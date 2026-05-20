package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/monolith-graphql-payment-gateway-merchant/service"
	pbdocument "github.com/MamangRust/monolith-graphql-payment-gateway-pb/merchant_document"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/requests"
	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"github.com/go-playground/validator/v10"
)

type merchantDocumentCommandHandleGrpc struct {
	pbdocument.UnimplementedMerchantDocumentCommandServiceServer

	merchantDocumentCommand service.MerchantDocumentCommandService
}

func NewMerchantDocumentCommandHandleGrpc(merchantCommand service.MerchantDocumentCommandService) MerchantDocumentCommandHandleGrpc {
	return &merchantDocumentCommandHandleGrpc{
		merchantDocumentCommand: merchantCommand,
	}
}

func (s *merchantDocumentCommandHandleGrpc) Create(ctx context.Context, req *pbdocument.CreateMerchantDocumentRequest) (*pbdocument.ApiResponseMerchantDocument, error) {
	request := requests.CreateMerchantDocumentRequest{
		MerchantID:   int(req.GetMerchantId()),
		DocumentType: req.GetDocumentType(),
		DocumentUrl:  req.GetDocumentUrl(),
	}

	if err := request.Validate(); err != nil {
		validations := s.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	document, err := s.merchantDocumentCommand.CreateMerchantDocument(ctx, &request)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoDocument := &pbdocument.MerchantDocument{
		DocumentId:   int32(document.DocumentID),
		MerchantId:   int32(document.MerchantID),
		DocumentType: document.DocumentType,
		DocumentUrl:  document.DocumentUrl,
		Status:       document.Status,
		Note:         StringValue(document.Note),
		UploadedAt:   document.UploadedAt.Time.Format(time.RFC3339),
		UpdatedAt:    document.UpdatedAt.Time.Format(time.RFC3339),
	}

	response := &pbdocument.ApiResponseMerchantDocument{
		Status:  "success",
		Message: "Successfully created merchant document",
		Data:    protoDocument,
	}

	return response, nil
}

func (s *merchantDocumentCommandHandleGrpc) Update(ctx context.Context, req *pbdocument.UpdateMerchantDocumentRequest) (*pbdocument.ApiResponseMerchantDocument, error) {
	id := int(req.GetDocumentId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "document_id",
				Message: "Document ID is required",
			},
		}))
	}

	request := requests.UpdateMerchantDocumentRequest{
		DocumentID:   &id,
		MerchantID:   int(req.GetMerchantId()),
		DocumentType: req.GetDocumentType(),
		DocumentUrl:  req.GetDocumentUrl(),
		Status:       req.GetStatus(),
		Note:         req.GetNote(),
	}

	if err := request.Validate(); err != nil {
		validations := s.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	document, err := s.merchantDocumentCommand.UpdateMerchantDocument(ctx, &request)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoDocument := &pbdocument.MerchantDocument{
		DocumentId:   int32(document.DocumentID),
		MerchantId:   int32(document.MerchantID),
		DocumentType: document.DocumentType,
		DocumentUrl:  document.DocumentUrl,
		Status:       document.Status,
		Note:         StringValue(document.Note),
		UploadedAt:   document.UploadedAt.Time.Format(time.RFC3339),
		UpdatedAt:    document.UpdatedAt.Time.Format(time.RFC3339),
	}

	response := &pbdocument.ApiResponseMerchantDocument{
		Status:  "success",
		Message: "Successfully updated merchant document",
		Data:    protoDocument,
	}

	return response, nil
}

func (s *merchantDocumentCommandHandleGrpc) UpdateStatus(ctx context.Context, req *pbdocument.UpdateMerchantDocumentStatusRequest) (*pbdocument.ApiResponseMerchantDocument, error) {
	id := int(req.GetDocumentId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "document_id",
				Message: "Document ID is required",
			},
		}))
	}

	request := requests.UpdateMerchantDocumentStatusRequest{
		DocumentID: &id,
		MerchantID: int(req.GetMerchantId()),
		Status:     req.GetStatus(),
		Note:       req.GetNote(),
	}

	if err := request.Validate(); err != nil {
		validations := s.parseValidationErrors(err)
		return nil, errors.ToGrpcError(errors.NewValidationError(validations))
	}

	document, err := s.merchantDocumentCommand.UpdateMerchantDocumentStatus(ctx, &request)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoDocument := &pbdocument.MerchantDocument{
		DocumentId:   int32(document.DocumentID),
		MerchantId:   int32(document.MerchantID),
		DocumentType: document.DocumentType,
		DocumentUrl:  document.DocumentUrl,
		Status:       document.Status,
		Note:         StringValue(document.Note),
		UploadedAt:   document.UploadedAt.Time.Format(time.RFC3339),
		UpdatedAt:    document.UpdatedAt.Time.Format(time.RFC3339),
	}

	response := &pbdocument.ApiResponseMerchantDocument{
		Status:  "success",
		Message: "Successfully updated merchant document status",
		Data:    protoDocument,
	}

	return response, nil
}

func (s *merchantDocumentCommandHandleGrpc) parseValidationErrors(err error) []errors.ValidationError {
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

func (s *merchantDocumentCommandHandleGrpc) getValidationMessage(fe validator.FieldError) string {
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

func (s *merchantDocumentCommandHandleGrpc) Trashed(ctx context.Context, req *pbdocument.FindMerchantDocumentByIdRequest) (*pbdocument.ApiResponseMerchantDocumentDeleteAt, error) {
	id := int(req.GetDocumentId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "document_id",
				Message: "Document ID is required",
			},
		}))
	}

	document, err := s.merchantDocumentCommand.TrashedMerchantDocument(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoDocument := &pbdocument.MerchantDocumentDeleteAt{
		DocumentId:   int32(document.DocumentID),
		MerchantId:   int32(document.MerchantID),
		DocumentType: document.DocumentType,
		DocumentUrl:  document.DocumentUrl,
		Status:       document.Status,
		Note:         StringValue(document.Note),
		UploadedAt:   document.UploadedAt.Time.Format(time.RFC3339),
		UpdatedAt:    document.UpdatedAt.Time.Format(time.RFC3339),
		DeletedAt:    &wrapperspb.StringValue{Value: document.DeletedAt.Time.Format(time.RFC3339)},
	}

	response := &pbdocument.ApiResponseMerchantDocumentDeleteAt{
		Status:  "success",
		Message: "Successfully trashed merchant document",
		Data:    protoDocument,
	}

	return response, nil
}

func (s *merchantDocumentCommandHandleGrpc) Restore(ctx context.Context, req *pbdocument.FindMerchantDocumentByIdRequest) (*pbdocument.ApiResponseMerchantDocumentDeleteAt, error) {
	id := int(req.GetDocumentId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "document_id",
				Message: "Document ID is required",
			},
		}))
	}

	document, err := s.merchantDocumentCommand.RestoreMerchantDocument(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	protoDocument := &pbdocument.MerchantDocumentDeleteAt{
		DocumentId:   int32(document.DocumentID),
		MerchantId:   int32(document.MerchantID),
		DocumentType: document.DocumentType,
		DocumentUrl:  document.DocumentUrl,
		Status:       document.Status,
		Note:         StringValue(document.Note),
		UploadedAt:   document.UploadedAt.Time.Format(time.RFC3339),
		UpdatedAt:    document.UpdatedAt.Time.Format(time.RFC3339),
	}

	response := &pbdocument.ApiResponseMerchantDocumentDeleteAt{
		Status:  "success",
		Message: "Successfully restored merchant document",
		Data:    protoDocument,
	}

	return response, nil
}

func (s *merchantDocumentCommandHandleGrpc) DeletePermanent(ctx context.Context, req *pbdocument.FindMerchantDocumentByIdRequest) (*pbdocument.ApiResponseMerchantDocumentDelete, error) {
	id := int(req.GetDocumentId())

	if id == 0 {
		return nil, errors.ToGrpcError(errors.NewValidationError([]errors.ValidationError{
			{
				Field:   "document_id",
				Message: "Document ID is required",
			},
		}))
	}

	_, err := s.merchantDocumentCommand.DeleteMerchantDocumentPermanent(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	response := &pbdocument.ApiResponseMerchantDocumentDelete{
		Status:  "success",
		Message: "Successfully permanently deleted merchant document",
	}

	return response, nil
}

func (s *merchantDocumentCommandHandleGrpc) RestoreAll(ctx context.Context, _ *emptypb.Empty) (*pbdocument.ApiResponseMerchantDocumentAll, error) {
	_, err := s.merchantDocumentCommand.RestoreAllMerchantDocument(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	response := &pbdocument.ApiResponseMerchantDocumentAll{
		Status:  "success",
		Message: "Successfully restored all merchant documents",
	}

	return response, nil
}

func (s *merchantDocumentCommandHandleGrpc) DeleteAllPermanent(ctx context.Context, _ *emptypb.Empty) (*pbdocument.ApiResponseMerchantDocumentAll, error) {
	_, err := s.merchantDocumentCommand.DeleteAllMerchantDocumentPermanent(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	response := &pbdocument.ApiResponseMerchantDocumentAll{
		Status:  "success",
		Message: "Successfully permanently deleted all merchant documents",
	}

	return response, nil
}

func StringValue(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func Int32Value(v *int32) int32 {
	if v == nil {
		return 0
	}

	return *v
}
