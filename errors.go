package api

import (
	"encoding/json"
	"fmt"
)

// API errors

type Error struct {
	Message   json.RawMessage `json:"message"`
	ErrorCode int             `json:"errorCode"`
	ErrorType string          `json:"errorType"`
}

func (p Error) Error() string {
	return fmt.Sprintf("Code: %d; Type: %s; Message(s): %s", p.ErrorCode, p.ErrorType, string(p.Message))
}
