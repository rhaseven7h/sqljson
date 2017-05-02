package sqljson_test

import (
	"database/sql"
	"testing"

	"github.com/rhaseven7h/sqljson"

	validator "gopkg.in/go-playground/validator.v9"

	. "github.com/smartystreets/goconvey/convey"
)

func TestValidatorIntegration(t *testing.T) {
	type validatorStruct struct {
		ID           int
		ContactEmail sqljson.NullString  `json:"contact_email" validate:"required,email"`
		IsAdmin      sqljson.NullBool    `json:"is_admin" validate:"required"`
		Followers    sqljson.NullInt64   `json:"followers" validate:"required,min=1000"`
		BankBalance  sqljson.NullFloat64 `json:"bank_balance" validate:"required,max=555.55"`
	}

	validate := validator.New()
	validate.RegisterCustomTypeFunc(sqljson.NullStringValidateValuer, sqljson.NullString{})
	validate.RegisterCustomTypeFunc(sqljson.NullBoolValidateValuer, sqljson.NullBool{})
	validate.RegisterCustomTypeFunc(sqljson.NullInt64ValidateValuer, sqljson.NullInt64{})
	validate.RegisterCustomTypeFunc(sqljson.NullFloat64ValidateValuer, sqljson.NullFloat64{})

	Convey("Given a struct to validate with valid values", t, func() {
		vs := &validatorStruct{
			ID: 10,
			ContactEmail: sqljson.NullString{
				NullString: sql.NullString{
					String: "user.subaddr@server.tld",
					Valid:  true,
				},
			},
			IsAdmin: sqljson.NullBool{
				NullBool: sql.NullBool{
					Bool:  true,
					Valid: true,
				},
			},
			Followers: sqljson.NullInt64{
				NullInt64: sql.NullInt64{
					Int64: 1001,
					Valid: true,
				},
			},
			BankBalance: sqljson.NullFloat64{
				NullFloat64: sql.NullFloat64{
					Float64: 555.55,
					Valid:   true,
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
	Convey("Given a struct to validate with invalid values", t, func() {
		vs := &validatorStruct{
			ID: 10,
			ContactEmail: sqljson.NullString{
				NullString: sql.NullString{
					String: "not an email",
					Valid:  true,
				},
			},
			IsAdmin: sqljson.NullBool{
				NullBool: sql.NullBool{
					Bool:  false,
					Valid: true,
				},
			},
			Followers: sqljson.NullInt64{
				NullInt64: sql.NullInt64{
					Int64: 999,
					Valid: true,
				},
			},
			BankBalance: sqljson.NullFloat64{
				NullFloat64: sql.NullFloat64{
					Float64: 555.56,
					Valid:   true,
				},
			},
		}
		Convey("When I validate it", func() {
			err := validate.Struct(vs)
			Convey("Then I should get appropriate errors", func() {
				validationErrors := err.(validator.ValidationErrors)
				So(err, ShouldNotBeNil)
				So(len(validationErrors), ShouldEqual, 4)
				So(validationErrors[0].Field(), ShouldEqual, "ContactEmail")
				So(validationErrors[0].Tag(), ShouldEqual, "email")
				So(validationErrors[1].Field(), ShouldEqual, "IsAdmin")
				So(validationErrors[1].Tag(), ShouldEqual, "required")
				So(validationErrors[2].Field(), ShouldEqual, "Followers")
				So(validationErrors[2].Tag(), ShouldEqual, "min")
				So(validationErrors[3].Field(), ShouldEqual, "BankBalance")
				So(validationErrors[3].Tag(), ShouldEqual, "max")
			})
		})
	})
}
