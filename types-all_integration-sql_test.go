package sqljson_test

import (
	"testing"

	"github.com/Rhaseven7h/sqljson"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSQLIntegration(t *testing.T) {
	type modelSupplier struct {
		ID           int
		ContactEmail sqljson.NullString
		IsAdmin      sqljson.NullBool
		Followers    sqljson.NullInt64
		BankBalance  sqljson.NullFloat64
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	Convey("Given a sql mock using a struct with sqljson.Null* - with valid values", t, func() {
		mock.
			ExpectQuery(`
                SELECT
                    id,
                    contact_email,
                    is_admin,
                    followers,
                    bank_balance
                FROM suppliers
            `).
			WillReturnRows(
				sqlmock.
					NewRows([]string{
						"id",
						"contact_email",
						"is_admin",
						"followers",
						"bank_balance",
					}).
					AddRow(
						10,
						"gmedina@ooyala.com",
						true,
						100,
						123.45,
					),
			)
		Convey("When I query a row and scan it", func() {
			supplier := &modelSupplier{}
			row := db.QueryRow(`
                SELECT
                    id,
                    contact_email,
                    is_admin,
                    followers,
                    bank_balance
                FROM suppliers
            `)
			dbErr := row.Scan(
				&(*supplier).ID,
				&(*supplier).ContactEmail,
				&(*supplier).IsAdmin,
				&(*supplier).Followers,
				&(*supplier).BankBalance,
			)
			mockErr := mock.ExpectationsWereMet()
			Convey("Then I should get the row correctly", func() {
				So(dbErr, ShouldBeNil)
				So(mockErr, ShouldBeNil)
				So(supplier.ID, ShouldEqual, 10)
				So(supplier.ContactEmail.Valid, ShouldBeTrue)
				So(supplier.ContactEmail.String, ShouldEqual, "gmedina@ooyala.com")
				So(supplier.IsAdmin.Valid, ShouldBeTrue)
				So(supplier.IsAdmin.Bool, ShouldBeTrue)
				So(supplier.Followers.Valid, ShouldBeTrue)
				So(supplier.Followers.Int64, ShouldEqual, 100)
				So(supplier.BankBalance.Valid, ShouldBeTrue)
				So(supplier.BankBalance.Float64, ShouldEqual, 123.45)
			})
		})
	})
	Convey("Given a sql mock using a struct with sqljson.Null* - with null values", t, func() {
		mock.
			ExpectQuery(`
                SELECT
                    id,
                    contact_email,
                    is_admin,
                    followers,
                    bank_balance
                FROM suppliers
            `).
			WillReturnRows(
				sqlmock.
					NewRows([]string{
						"id",
						"contact_email",
						"is_admin",
						"followers",
						"bank_balance",
					}).
					AddRow(
						10,
						nil,
						nil,
						nil,
						nil,
					),
			)
		Convey("When I query a row and scan it", func() {
			supplier := &modelSupplier{}
			row := db.QueryRow(`
                SELECT
                    id,
                    contact_email,
                    is_admin,
                    followers,
                    bank_balance
                FROM suppliers
            `)
			dbErr := row.Scan(
				&(*supplier).ID,
				&(*supplier).ContactEmail,
				&(*supplier).IsAdmin,
				&(*supplier).Followers,
				&(*supplier).BankBalance,
			)
			mockErr := mock.ExpectationsWereMet()
			Convey("Then I should get the row correctly", func() {
				So(dbErr, ShouldBeNil)
				So(mockErr, ShouldBeNil)
				So(supplier.ID, ShouldEqual, 10)
				So(supplier.ContactEmail.Valid, ShouldBeFalse)
				So(supplier.IsAdmin.Valid, ShouldBeFalse)
				So(supplier.Followers.Valid, ShouldBeFalse)
				So(supplier.BankBalance.Valid, ShouldBeFalse)
			})
		})
	})
}
