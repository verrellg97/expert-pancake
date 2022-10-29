package validator

import "github.com/go-playground/validator/v10"

type Validator struct {
	Validator *validator.Validate
}

func NewValidator() Validator {
	return Validator{
		Validator: validator.New(),
	}
}

func (ths *Validator) Validate(data interface{}) map[string]string {
	if err := ths.Validator.Struct(data); err != nil {
		errorMap := make(map[string]string)

		switch errType := err.(type) {
		case *validator.InvalidValidationError:
			errorMap[errType.Type.Name()] = errType.Error()
		case validator.ValidationErrors:
			for _, value := range errType {
				errorMap[value.Namespace()] = value.Tag() + " " + value.Param()
			}
		default:
			errorMap["message"] = err.Error()
		}

		return errorMap
	}

	return nil
}
