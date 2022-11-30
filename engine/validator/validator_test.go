package validator

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestValidate(t *testing.T) {

	Convey("Given validate instance", t, func() {

		validator := NewValidator()

		type customType struct {
			Key    string `validate:"required"`
			Value  string `validate:"required"`
			Number int    `validate:"required,min=10"`
		}

		Convey("error", func() {
			errMap := validator.Validate(customType{})
			So(errMap, ShouldResemble, map[string]string{
				"customType.Key":    "required ",
				"customType.Value":  "required ",
				"customType.Number": "required ",
			})
		})

		Convey("error for minimal value", func() {
			errMap := validator.Validate(customType{
				Key:    "test",
				Value:  "something here",
				Number: 5,
			})
			So(errMap, ShouldResemble, map[string]string{
				"customType.Number": "min 10",
			})
		})

		Convey("success", func() {
			errMap := validator.Validate(customType{
				Key:    "test",
				Value:  "something here",
				Number: 10,
			})
			So(errMap, ShouldBeNil)
		})
	})

}
