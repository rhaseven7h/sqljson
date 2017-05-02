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
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	Convey("Given a sql mock using a struct with sqljson.NullString - with valid value", t, func() {
		mock.
			ExpectQuery("SELECT id, contact_email FROM suppliers").
			WillReturnRows(
				sqlmock.
					NewRows([]string{"id", "contact_email"}).
					AddRow(7, "gmedina@ooyala.com"),
			)
		Convey("When I query a row and scan it", func() {
			supplier := &modelSupplier{}
			row := db.QueryRow("SELECT id, contact_email FROM suppliers")
			dbErr := row.Scan(
				&(*supplier).ID,
				&(*supplier).ContactEmail,
			)
			mockErr := mock.ExpectationsWereMet()
			Convey("Then I should get the row correctly", func() {
				So(dbErr, ShouldBeNil)
				So(mockErr, ShouldBeNil)
				So(supplier.ID, ShouldEqual, 7)
				So(supplier.ContactEmail.Valid, ShouldBeTrue)
				So(supplier.ContactEmail.String, ShouldEqual, "gmedina@ooyala.com")
			})
		})
	})
	Convey("Given a sql mock using a struct with sqljson.NullString - with null value", t, func() {
		mock.
			ExpectQuery("SELECT id, contact_email FROM suppliers").
			WillReturnRows(
				sqlmock.
					NewRows([]string{"id", "contact_email"}).
					AddRow(7, nil),
			)
		Convey("When I query a row and scan it", func() {
			supplier := &modelSupplier{}
			row := db.QueryRow("SELECT id, contact_email FROM suppliers")
			dbErr := row.Scan(
				&(*supplier).ID,
				&(*supplier).ContactEmail,
			)
			mockErr := mock.ExpectationsWereMet()
			Convey("Then I should get the row correctly", func() {
				So(dbErr, ShouldBeNil)
				So(mockErr, ShouldBeNil)
				So(supplier.ID, ShouldEqual, 7)
				So(supplier.ContactEmail.Valid, ShouldBeFalse)
			})
		})
	})
}
