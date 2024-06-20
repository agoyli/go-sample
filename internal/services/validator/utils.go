package validator

import (
	"main/internal/utils"
	"strings"

	"github.com/go-playground/validator"
)

func ValidateStruct(r interface{}) utils.AppErrorCollection {
	v := validator.New()
	errs := utils.AppErrorCollection{}
	if err := v.Struct(r); err != nil {
		for _, valErr := range err.(validator.ValidationErrors) {
			errs.Append(*utils.NewAppError(valErr.Tag(), pascalCaseToSnakeCase(valErr.Field()), err.Error()))
		}
	}
	return errs
}

func pascalCaseToSnakeCase(str string) string {
	snake := []rune{}
	isUpper := false
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			if !isUpper {
				snake = append(snake, '_')
			}
			isUpper = true
			snake = append(snake, rune(char+32)) // Convert to lowercase
		} else {
			snake = append(snake, char)
			isUpper = false
		}
	}
	return strings.Trim(string(snake), " _")
}
