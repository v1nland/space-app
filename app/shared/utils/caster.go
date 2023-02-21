package utils

import "github.com/google/uuid"

func CastUuidToStringPointer(id *uuid.UUID) *string {
	idAsString := id.String()

	return &idAsString
}
