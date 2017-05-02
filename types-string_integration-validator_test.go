package sqljson_test

import (
	"database/sql"
	"testing"

	"github.com/Rhaseven7h/sqljson"

	validator "gopkg.in/go-playground/validator.v9"

	. "github.com/smartystreets/goconvey/convey"
)

func TestValidatorIntegration(t *testing.T) {
	type validatorStruct struct {
		ID           int
		ContactEmail sqljson.NullString `validate:"required,email"`
	}
	validate := validator.New()
	validate.RegisterCustomTypeFunc(sqljson.NullStringValidateValuer, sqljson.NullString{})
	Convey("Given a struct to validate with: valid email values", t, func() {
		vs := &validatorStruct{
			ID: 7,
			ContactEmail: sqljson.NullString{
				sql.NullString{
					String: "user.subaddr@server.tld",
					Valid:  true,
				},
			},
		}
		Convey("When I validate it", func() {
			err := validate.Struct(vs)
			Convey("Then I should get an appropriate result", func() {
				So(err, ShouldBeNil)
			})
		})
	})
	Convey("Given a struct to validate with: not valid email value", t, func() {
		vs := &validatorStruct{
			ID: 7,
			ContactEmail: sqljson.NullString{
				sql.NullString{
					String: "not empty, but not an email",
					Valid:  true,
				},
			},
		}
		Convey("When I validate it", func() {
			err := validate.Struct(vs)
			validationErrors := err.(validator.ValidationErrors)
			Convey("Then I should get an appropriate result", func() {
				So(err, ShouldNotBeNil)
				So(len(validationErrors), ShouldEqual, 1)
				So(validationErrors[0].Tag(), ShouldEqual, "email")
				So(validationErrors[0].StructField(), ShouldEqual, "ContactEmail")
			})
		})
	})
	Convey("Given a struct to validate with: empty value", t, func() {
		vs := &validatorStruct{
			ID: 7,
			ContactEmail: sqljson.NullString{
				sql.NullString{
					String: "",
					Valid:  true,
				},
			},
		}
		Convey("When I validate it", func() {
			err := validate.Struct(vs)
			validationErrors := err.(validator.ValidationErrors)
			Convey("Then I should get an appropriate result", func() {
				So(err, ShouldNotBeNil)
				So(len(validationErrors), ShouldEqual, 1)
				So(validationErrors[0].Tag(), ShouldEqual, "required")
				So(validationErrors[0].StructField(), ShouldEqual, "ContactEmail")
			})
		})
	})
	Convey("Given a struct to validate with: null value", t, func() {
		vs := &validatorStruct{
			ID: 7,
			ContactEmail: sqljson.NullString{
				sql.NullString{
					String: "",
					Valid:  false,
				},
			},
		}
		Convey("When I validate it", func() {
			err := validate.Struct(vs)
			validationErrors := err.(validator.ValidationErrors)
			Convey("Then I should get an appropriate result", func() {
				So(err, ShouldNotBeNil)
				So(len(validationErrors), ShouldEqual, 1)
				So(validationErrors[0].Tag(), ShouldEqual, "required")
				So(validationErrors[0].StructField(), ShouldEqual, "ContactEmail")
			})
		})
	})
}
