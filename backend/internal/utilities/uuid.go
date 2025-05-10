package utilities

import "github.com/google/uuid"

func ValidateUUID(uuidStr string) (*uuid.UUID, error) {
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
