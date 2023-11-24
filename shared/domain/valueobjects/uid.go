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

func (u UID) String() string {
	return u.value
}
