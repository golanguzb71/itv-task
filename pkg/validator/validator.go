package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	if len(v) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString("Validation errors:\n")
	for _, err := range v {
		sb.WriteString(fmt.Sprintf("- %s: %s\n", err.Field, err.Message))
	}
	return sb.String()
}

func ValidateStruct(s interface{}) (ValidationErrors, error) {
	validate := validator.New()
	err := validate.Struct(s)

	if err == nil {
		return nil, nil
	}

	var validationErrors ValidationErrors
	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		for _, err := range validationErrs {
			field := ToSnakeCase(err.Field())
			message := getErrorMessage(err)
			validationErrors = append(validationErrors, ValidationError{
				Field:   field,
				Message: message,
			})
		}
		return validationErrors, nil
	}

	return nil, err
}

// Helper function to convert error tag to a human-readable message
func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Must be a valid email address"
	case "min":
		return fmt.Sprintf("Must be at least %s", err.Param())
	case "max":
		return fmt.Sprintf("Must not be greater than %s", err.Param())
	case "gt":
		return fmt.Sprintf("Must be greater than %s", err.Param())
	case "lt":
		return fmt.Sprintf("Must be less than %s", err.Param())
	default:
		return fmt.Sprintf("Failed validation on %s", err.Tag())
	}
}

// ToSnakeCase converts a camel case string to snake case
func ToSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && 'A' <= r && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}
