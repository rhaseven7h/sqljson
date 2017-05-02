package sqljson_test

import (
	"encoding/json"
	"testing"

	"github.com/Rhaseven7h/sqljson"
	. "github.com/smartystreets/goconvey/convey"
)

func TestJSONIntegration(t *testing.T) {
	type jsonIntegrationTestStruct struct {
		ID          int                 `json:"id"`
		Name        sqljson.NullString  `json:"name"`
		IsAdmin     sqljson.NullBool    `json:"is_admin"`
		Followers   sqljson.NullInt64   `json:"followers"`
		BankBalance sqljson.NullFloat64 `json:"bank_balance"`
	}
	Convey("Given a  JSON string using valid non-null non-zero values", t, func() {
		strJSON := []byte(`
            {
                "id": 10,
                "name": "Gabriel",
                "is_admin": true,
                "followers": 1000,
                "bank_balance": 123.45
            }
        `)
		Convey("When I unmarshal it onto structs using sqljson.Null*", func() {
			jsonStruct := &jsonIntegrationTestStruct{}
			err := json.Unmarshal(strJSON, jsonStruct)
			Convey("Then I should get the correct values back", func() {
				So(err, ShouldBeNil)
				So(jsonStruct.ID, ShouldEqual, 10)
				So(jsonStruct.Name.Valid, ShouldBeTrue)
				So(jsonStruct.Name.String, ShouldEqual, "Gabriel")
				So(jsonStruct.IsAdmin.Valid, ShouldBeTrue)
				So(jsonStruct.IsAdmin.Bool, ShouldBeTrue)
				So(jsonStruct.Followers.Valid, ShouldBeTrue)
				So(jsonStruct.Followers.Int64, ShouldEqual, 1000)
				So(jsonStruct.BankBalance.Valid, ShouldBeTrue)
				So(jsonStruct.BankBalance.Float64, ShouldEqual, 123.45)
			})
		})
	})
	Convey("Given a  JSON string using valid non-null zero values", t, func() {
		strJSON := []byte(`
            {
                "id": 0,
                "name": "",
                "is_admin": false,
                "followers": 0,
                "bank_balance": 0.0
            }
        `)
		Convey("When I unmarshal it onto structs using sqljson.Null*", func() {
			jsonStruct := &jsonIntegrationTestStruct{}
			err := json.Unmarshal(strJSON, jsonStruct)
			Convey("Then I should get the correct values back", func() {
				So(err, ShouldBeNil)
				So(jsonStruct.ID, ShouldEqual, 0)
				So(jsonStruct.Name.Valid, ShouldBeTrue)
				So(jsonStruct.Name.String, ShouldEqual, "")
				So(jsonStruct.IsAdmin.Valid, ShouldBeTrue)
				So(jsonStruct.IsAdmin.Bool, ShouldBeFalse)
				So(jsonStruct.Followers.Valid, ShouldBeTrue)
				So(jsonStruct.Followers.Int64, ShouldEqual, 0)
				So(jsonStruct.BankBalance.Valid, ShouldBeTrue)
				So(jsonStruct.BankBalance.Float64, ShouldEqual, 0.0)
			})
		})
	})
	Convey("Given a  JSON string using valid null values", t, func() {
		strJSON := []byte(`
            {
                "id": 10,
                "name": null,
                "is_admin": null,
                "followers": null,
                "bank_balance": null
            }
        `)
		Convey("When I unmarshal it onto structs using sqljson.Null*", func() {
			jsonStruct := &jsonIntegrationTestStruct{}
			err := json.Unmarshal(strJSON, jsonStruct)
			Convey("Then I should get the correct values back", func() {
				So(err, ShouldBeNil)
				So(jsonStruct.ID, ShouldEqual, 10)
				So(jsonStruct.Name.Valid, ShouldBeFalse)
				So(jsonStruct.IsAdmin.Valid, ShouldBeFalse)
				So(jsonStruct.Followers.Valid, ShouldBeFalse)
				So(jsonStruct.BankBalance.Valid, ShouldBeFalse)
			})
		})
	})
}
