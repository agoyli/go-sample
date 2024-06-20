package utils

import "strings"

type AppError struct {
	code    string
	key     string
	comment string
}

var (
	ErrInvalid  = NewAppError("invalid", "", "please fill with valid options")
	ErrRequired = NewAppError("required", "", "please fill required keys")
	ErrNotfound = NewAppError("not_found", "", "")
)

func (err AppError) Error() string {
	return strings.Trim(err.code+" : "+err.key+" - "+err.comment, " -:")
}

type AppErrorCollection struct {
	Errors []AppError
}

func (err AppErrorCollection) Error() string {
	if err.HasError() {
		suffix := ""
		if len(err.Errors) > 1 {
			suffix = "..."
		}
		return err.Errors[0].Error() + suffix
	}
	return ""
}

func (err AppErrorCollection) HasError() bool {
	return err.Errors != nil && len(err.Errors) > 0
}

func (err *AppErrorCollection) Append(e AppError) {
	err.Errors = append(err.Errors, e)
}
func (err *AppErrorCollection) Merge(e AppErrorCollection) {
	err.Errors = append(err.Errors, e.Errors...)
}

func NewAppError(code, key, comment string) *AppError {
	return &AppError{code, key, comment}
}

func (err AppError) Code() string {
	return err.code
}

func (err AppError) Key() string {
	return err.key
}

func (err AppError) Comment() string {
	return err.comment
}

func (err *AppError) SetKey(key string) *AppError {
	err.key = key
	return err
}

func (err *AppError) SetComment(comment string) *AppError {
	err.comment = comment
	return err
}
