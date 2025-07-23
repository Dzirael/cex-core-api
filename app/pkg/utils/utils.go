package utils

import (
	"encoding/json"
	"fmt"
	"io"
)

func UnmarhalRawResponse[T any](reader io.Reader) (*T, error) {
	var out T
	bodyBytes, err := io.ReadAll(reader)
	if err != nil {
		return &out, fmt.Errorf("read body: %w", err)
	}

	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		return &out, fmt.Errorf("parse error: %w", err)
	}

	return &out, nil
}

func ToPointer[T any](v T) *T {
	return &v
}
