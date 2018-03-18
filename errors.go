package api

import (
	"encoding/json"
	"fmt"
)

type Error interface {
	VSZError() bool
}

type VSZError struct {
	Message   json.RawMessage `json:"message"`
	ErrorCode int             `json:"errorCode"`
	ErrorType string          `json:"errorType"`
}

func (p VSZError) Error() string {
	return fmt.Sprintf("Code: %d; Type: %s; Message(s): %s", p.ErrorCode, p.ErrorType, string(p.Message))
}

func (VSZError) AuthenticationError() bool {
	return false
}

func (VSZError) VSZError() bool {
	return true
}
