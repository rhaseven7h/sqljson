package sqljson_test

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/rhaseven7h/sqljson"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNullBoolValidateValuer(t *testing.T) {
	Convey("Given a non-null sqljson.NullBool true value", t, func() {
		ov := sqljson.NullBool{
			NullBool: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
		}
		ivIn := reflect.ValueOf(ov)
		Convey("When evaluated", func() {
			ivOut := sqljson.NullBoolValidateValuer(ivIn)
			Convey("Then we should get the value", func() {
				nv, ok := ivOut.(bool)
				So(ok, ShouldBeTrue)
				So(nv, ShouldBeTrue)
			})
		})
	})
	Convey("Given a non-null sqljson.NullBool false value", t, func() {
		ov := sqljson.NullBool{
			NullBool: sql.NullBool{
				Bool:  false,
				Valid: true,
			},
		}
		ivIn := reflect.ValueOf(ov)
		Convey("When evaluated", func() {
			ivOut := sqljson.NullBoolValidateValuer(ivIn)
			Convey("Then we should get the value", func() {
				nv, ok := ivOut.(bool)
				So(ok, ShouldBeTrue)
				So(nv, ShouldBeFalse)
			})
		})
	})
	Convey("Given a null sqljson.NullBool value", t, func() {
		ov := sqljson.NullBool{
			NullBool: sql.NullBool{
				Bool:  false,
				Valid: false,
			},
		}
		ivIn := reflect.ValueOf(ov)
		Convey("When evaluated", func() {
			ivOut := sqljson.NullBoolValidateValuer(ivIn)
			Convey("Then we should get the value", func() {
				So(ivOut, ShouldBeNil)
			})
		})
	})
}

func TestBoolPtrOrNil(t *testing.T) {
	Convey("Given a sqljson.NullBool non-null true value", t, func() {
		in := sqljson.NullBool{
			NullBool: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
		}
		Convey("When I call BoolPtrOrNil", func() {
			bptr := in.BoolPtrOrNil()
			Convey("Then I should get a true bool value pointer", func() {
				So(bptr, ShouldNotBeNil)
				So(*bptr, ShouldBeTrue)
			})
		})
	})
	Convey("Given a sqljson.NullBool non-null false value", t, func() {
		in := sqljson.NullBool{
			NullBool: sql.NullBool{
				Bool:  false,
				Valid: true,
			},
		}
		Convey("When I call BoolPtrOrNil", func() {
			bptr := in.BoolPtrOrNil()
			Convey("Then I should get a false bool value pointer", func() {
				So(bptr, ShouldNotBeNil)
				So(*bptr, ShouldBeFalse)
			})
		})
	})
	Convey("Given a sqljson.NullBool null value", t, func() {
		in := sqljson.NullBool{
			NullBool: sql.NullBool{
				Bool:  false,
				Valid: false,
			},
		}
		Convey("When I call BoolPtrOrNil", func() {
			bptr := in.BoolPtrOrNil()
			Convey("Then I should get a nil pointer", func() {
				So(bptr, ShouldBeNil)
			})
		})
	})
}
func TestBoolMarshalJSON(t *testing.T) {
	Convey("Given a non-null sqljson.NullBool", t, func() {
		in := sqljson.NullBool{
			NullBool: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
		}
		Convey("When marshalled to JSON", func() {
			b, err := in.MarshalJSON()
			Convey("Then I should get a valid value", func() {
				So(err, ShouldBeNil)
				So(string(b), ShouldEqual, "true")
			})
		})
	})
	Convey("Given a null sqljson.NullBool", t, func() {
		in := sqljson.NullBool{
			NullBool: sql.NullBool{
				Bool:  true,
				Valid: false,
			},
		}
		Convey("When marshalled to JSON", func() {
			b, err := in.MarshalJSON()
			Convey("Then I should get a null value", func() {
				So(err, ShouldBeNil)
				So(string(b), ShouldEqual, "null")
			})
		})
	})
}

func TestBoolUnmarshalJSON(t *testing.T) {
	Convey("Given a true JSON string", t, func() {
		obj := sqljson.NullBool{}
		jsonIn := `true`
		Convey("When unmarshalled", func() {
			err := obj.UnmarshalJSON([]byte(jsonIn))
			Convey("Then I should get the corresponding value back", func() {
				So(err, ShouldBeNil)
				So(obj.Valid, ShouldBeTrue)
				So(obj.Bool, ShouldBeTrue)
			})
		})
	})
	Convey("Given a false JSON string", t, func() {
		obj := sqljson.NullBool{}
		jsonIn := `false`
		Convey("When unmarshalled", func() {
			err := obj.UnmarshalJSON([]byte(jsonIn))
			Convey("Then I should get the corresponding value back", func() {
				So(err, ShouldBeNil)
				So(obj.Valid, ShouldBeTrue)
				So(obj.Bool, ShouldBeFalse)
			})
		})
	})
	Convey("Given a null JSON string", t, func() {
		obj := sqljson.NullBool{}
		jsonIn := `null`
		Convey("When unmarshalled", func() {
			err := obj.UnmarshalJSON([]byte(jsonIn))
			Convey("Then I should get the corresponding value back", func() {
				So(err, ShouldBeNil)
				So(obj.Valid, ShouldBeFalse)
			})
		})
	})
	type errorCaseDef struct {
		Label      string
		InputJSON  string
		TestString string
	}
	errorCases := []errorCaseDef{
		errorCaseDef{"an invalid", `invalid value`, "invalid character"},
		errorCaseDef{"an empty", ``, "unexpected end of JSON input"},
		errorCaseDef{"an integer", `7`, "cannot unmarshal number into Go value of type bool"},
		errorCaseDef{"a float", `123.45`, "cannot unmarshal number into Go value of type bool"},
		errorCaseDef{"a string", `"Hello!"`, "cannot unmarshal string into Go value of type bool"},
		errorCaseDef{"an object", `{"status":"ok"}`, "cannot unmarshal object into Go value of type bool"},
		errorCaseDef{"an array", `[1,2,3]`, "cannot unmarshal array into Go value of type bool"},
	}
	for _, errorCase := range errorCases {
		Convey("Given "+errorCase.Label+" JSON string", t, func() {
			obj := sqljson.NullBool{}
			jsonIn := errorCase.InputJSON
			Convey("When unmarshalled", func() {
				err := obj.UnmarshalJSON([]byte(jsonIn))
				Convey("Then I should get an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldContainSubstring, errorCase.TestString)
				})
			})
		})
	}
}
