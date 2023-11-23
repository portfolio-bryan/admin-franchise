package valueobjects

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type UID struct {
	value string
}

var ErrInvalidUID = errors.New("invalid UID")

func NewUID(value string) (UID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return UID{}, fmt.Errorf("%w: %s", ErrInvalidUID, value)
	}

	return UID{
		value: v.String(),
	}, nil
}

type Protocol struct {
	value string
}

var ErrInvalidProtocol = errors.New("invalid Protocol")

func NewProtocol(value string) (Protocol, error) {
	// Provisional logic
	if value != "http" && value != "https" {
		return Protocol{}, fmt.Errorf("%w: %s", ErrInvalidProtocol, value)
	}

	return Protocol{
		value: value,
	}, nil
}
