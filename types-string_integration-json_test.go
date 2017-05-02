package sqljson_test

import (
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/Rhaseven7h/sqljson"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJSONIntegration(t *testing.T) {
	type jsonIntegrationTestStruct struct {
		ID   int                `json:"id"`
		Name sqljson.NullString `json:"name"`
	}
	Convey("Given a  JSON string using sqljson.NullString - with valid non-empty string", t, func() {
		strJSON := []byte(`{"id": 7, "name": "Gabriel"}`)
		Convey("When I unmarshal it onto a struct using sqljson.NullString", func() {
			jsonStruct := &jsonIntegrationTestStruct{}
			err := json.Unmarshal(strJSON, jsonStruct)
			Convey("Then I should get the correct value back", func() {
				So(err, ShouldBeNil)
				So(jsonStruct.ID, ShouldEqual, 7)
				So(jsonStruct.Name.Valid, ShouldBeTrue)
				So(jsonStruct.Name.String, ShouldEqual, "Gabriel")
			})
		})
	})
	Convey("Given a  JSON string using sqljson.NullString - with valid empty string", t, func() {
		strJSON := []byte(`{"id": 7, "name": ""}`)
		Convey("When I unmarshal it onto a struct using sqljson.NullString", func() {
			jsonStruct := &jsonIntegrationTestStruct{}
			err := json.Unmarshal(strJSON, jsonStruct)
			Convey("Then I should get the correct value back", func() {
				So(err, ShouldBeNil)
				So(jsonStruct.ID, ShouldEqual, 7)
				So(jsonStruct.Name.Valid, ShouldBeTrue)
				So(jsonStruct.Name.String, ShouldEqual, "")
			})
		})
	})
	Convey("Given a  JSON string using sqljson.NullString - with null", t, func() {
		strJSON := []byte(`{"id": 7, "name": null}`)
		Convey("When I unmarshal it onto a struct using sqljson.NullString", func() {
			jsonStruct := &jsonIntegrationTestStruct{}
			err := json.Unmarshal(strJSON, jsonStruct)
			Convey("Then I should get the correct value back", func() {
				So(err, ShouldBeNil)
				So(jsonStruct.ID, ShouldEqual, 7)
				So(jsonStruct.Name.Valid, ShouldBeFalse)
			})
		})
	})
	Convey("Given a  JSON string using sqljson.NullString - with omitted value", t, func() {
		strJSON := []byte(`{"id": 7}`)
		Convey("When I unmarshal it onto a struct using sqljson.NullString", func() {
			jsonStruct := &jsonIntegrationTestStruct{}
			err := json.Unmarshal(strJSON, jsonStruct)
			Convey("Then I should get the correct value back", func() {
				So(err, ShouldBeNil)
				So(jsonStruct.ID, ShouldEqual, 7)
				So(jsonStruct.Name.Valid, ShouldBeFalse)
			})
		})
	})
	Convey("Given a  JSON string using sqljson.NullString - with no (invalid json) value", t, func() {
		strJSON := []byte(`{"id": 7, "name": }`)
		Convey("When I unmarshal it onto a struct using sqljson.NullString", func() {
			jsonStruct := &jsonIntegrationTestStruct{}
			err := json.Unmarshal(strJSON, jsonStruct)
			Convey("Then I should get the correct value back", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "invalid character")
			})
		})
	})
	Convey("Given a  JSON string using sqljson.NullString - with bool true value", t, func() {
		strJSON := []byte(`{"id": 7, "name": true }`)
		Convey("When I unmarshal it onto a struct using sqljson.NullString", func() {
			jsonStruct := &jsonIntegrationTestStruct{}
			err := json.Unmarshal(strJSON, jsonStruct)
			Convey("Then I should get the correct value back", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "cannot unmarshal bool")
			})
		})
	})
	Convey("Given a  JSON string using sqljson.NullString - with bool false value", t, func() {
		strJSON := []byte(`{"id": 7, "name": false }`)
		Convey("When I unmarshal it onto a struct using sqljson.NullString", func() {
			jsonStruct := &jsonIntegrationTestStruct{}
			err := json.Unmarshal(strJSON, jsonStruct)
			Convey("Then I should get the correct value back", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "cannot unmarshal bool")
			})
		})
	})
	Convey("Given a  JSON string using sqljson.NullString - with integer zero value", t, func() {
		strJSON := []byte(`{"id": 7, "name": 0 }`)
		Convey("When I unmarshal it onto a struct using sqljson.NullString", func() {
			jsonStruct := &jsonIntegrationTestStruct{}
			err := json.Unmarshal(strJSON, jsonStruct)
			Convey("Then I should get the correct value back", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "cannot unmarshal number")
			})
		})
	})
	Convey("Given a  JSON string using sqljson.NullString - with integer non-zero value", t, func() {
		strJSON := []byte(`{"id": 7, "name": 5 }`)
		Convey("When I unmarshal it onto a struct using sqljson.NullString", func() {
			jsonStruct := &jsonIntegrationTestStruct{}
			err := json.Unmarshal(strJSON, jsonStruct)
			Convey("Then I should get the correct value back", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "cannot unmarshal number")
			})
		})
	})
	Convey("Given a  JSON string using sqljson.NullString - with float zero value", t, func() {
		strJSON := []byte(`{"id": 7, "name": 0.0 }`)
		Convey("When I unmarshal it onto a struct using sqljson.NullString", func() {
			jsonStruct := &jsonIntegrationTestStruct{}
			err := json.Unmarshal(strJSON, jsonStruct)
			Convey("Then I should get the correct value back", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "cannot unmarshal number")
			})
		})
	})
	Convey("Given a  JSON string using sqljson.NullString - with float non-zero value", t, func() {
		strJSON := []byte(`{"id": 7, "name": 1.23 }`)
		Convey("When I unmarshal it onto a struct using sqljson.NullString", func() {
			jsonStruct := &jsonIntegrationTestStruct{}
			err := json.Unmarshal(strJSON, jsonStruct)
			Convey("Then I should get the correct value back", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "cannot unmarshal number")
			})
		})
	})
	Convey("Given a sqljson.NullString value - with a Null NullString field", t, func() {
		jsonStruct := &jsonIntegrationTestStruct{
			ID: 7,
			Name: sqljson.NullString{
				sql.NullString{
					String: "",
					Valid:  false,
				},
			},
		}
		Convey("When I marshal it onto a JSON string", func() {
			jsonBytes, err := json.Marshal(jsonStruct)
			jsonString := string(jsonBytes)
			Convey("Then I should get a correct answer", func() {
				So(err, ShouldBeNil)
				So(jsonString, ShouldEqual, `{"id":7,"name":null}`)
			})
		})
	})
	Convey("Given a sqljson.NullString value - with a non-Null empty NullString field", t, func() {
		jsonStruct := &jsonIntegrationTestStruct{
			ID: 7,
			Name: sqljson.NullString{
				sql.NullString{
					String: "",
					Valid:  true,
				},
			},
		}
		Convey("When I marshal it onto a JSON string", func() {
			jsonBytes, err := json.Marshal(jsonStruct)
			jsonString := string(jsonBytes)
			Convey("Then I should get a correct answer", func() {
				So(err, ShouldBeNil)
				So(jsonString, ShouldEqual, `{"id":7,"name":""}`)
			})
		})
	})
	Convey("Given a sqljson.NullString value - with a non-Null non-empty NullString field", t, func() {
		jsonStruct := &jsonIntegrationTestStruct{
			ID: 7,
			Name: sqljson.NullString{
				sql.NullString{
					String: "Gabriel",
					Valid:  true,
				},
			},
		}
		Convey("When I marshal it onto a JSON string", func() {
			jsonBytes, err := json.Marshal(jsonStruct)
			jsonString := string(jsonBytes)
			Convey("Then I should get a correct answer", func() {
				So(err, ShouldBeNil)
				So(jsonString, ShouldEqual, `{"id":7,"name":"Gabriel"}`)
			})
		})
	})
}
