package validator

import (
	"main/internal/services"
)

func ValidateUserCreate(dto services.UserCreateDto) error {
	err := ValidateStruct(dto)
	// other checks: err.Append(*utils.ErrInvalid)
	if err.HasError() {
		return err
	}
	return nil
}

func ValidateUserQuery(dto services.UserQueryDto) error {
	err := ValidateStruct(dto)

	if err.HasError() {
		return err
	}
	return nil
}
