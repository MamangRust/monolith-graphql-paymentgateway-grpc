package graphqlerror

import (
	"fmt"

	"github.com/MamangRust/monolith-graphql-payment-gateway-shared/domain/response"
)

func ToGraphqlErrorFromErrorResponse(err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("graphql error: %v", err)
}

func NewGraphqlError(statusText string, message string, code int) error {
	errResp := &response.ErrorResponse{
		Status:  statusText,
		Message: message,
		Code:    code,
	}

	return fmt.Errorf("graphql error: [%d] %s - %s", errResp.Code, errResp.Status, errResp.Message)
}
