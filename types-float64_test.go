package sqljson_test

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/Rhaseven7h/sqljson"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNullFloat64ValidateValuer(t *testing.T) {
	Convey("Given a non-null sqljson.NullFloat64 non-zero value", t, func() {
		ov := sqljson.NullFloat64{
			NullFloat64: sql.NullFloat64{
				Float64: 100.5,
				Valid:   true,
			},
		}
		ivIn := reflect.ValueOf(ov)
		Convey("When evaluated", func() {
			ivOut := sqljson.NullFloat64ValidateValuer(ivIn)
			Convey("Then we should get the value", func() {
				nv, ok := ivOut.(float64)
				So(ok, ShouldBeTrue)
				So(nv, ShouldEqual, 100.5)
			})
		})
	})
	Convey("Given a non-null sqljson.NullFloat64 zero value", t, func() {
		ov := sqljson.NullFloat64{
			NullFloat64: sql.NullFloat64{
				Float64: 0.0,
				Valid:   true,
			},
		}
		ivIn := reflect.ValueOf(ov)
		Convey("When evaluated", func() {
			ivOut := sqljson.NullFloat64ValidateValuer(ivIn)
			Convey("Then we should get the value", func() {
				nv, ok := ivOut.(float64)
				So(ok, ShouldBeTrue)
				So(nv, ShouldEqual, 0.0)
			})
		})
	})
	Convey("Given a null sqljson.NullFloat64 value", t, func() {
		ov := sqljson.NullFloat64{
			NullFloat64: sql.NullFloat64{
				Float64: 0.0,
				Valid:   false,
			},
		}
		ivIn := reflect.ValueOf(ov)
		Convey("When evaluated", func() {
			ivOut := sqljson.NullFloat64ValidateValuer(ivIn)
			Convey("Then we should get the value", func() {
				So(ivOut, ShouldBeNil)
			})
		})
	})
}

func TestFloat64PtrOrNil(t *testing.T) {
	Convey("Given a sqljson.NullFloat64 non-null non-zero value", t, func() {
		in := sqljson.NullFloat64{
			NullFloat64: sql.NullFloat64{
				Float64: 100.5,
				Valid:   true,
			},
		}
		Convey("When I call Float64PtrOrNil", func() {
			bptr := in.Float64PtrOrNil()
			Convey("Then I should get a true float64 value pointer", func() {
				So(bptr, ShouldNotBeNil)
				So(*bptr, ShouldEqual, 100.5)
			})
		})
	})
	Convey("Given a sqljson.NullFloat64 non-null non-zero value", t, func() {
		in := sqljson.NullFloat64{
			NullFloat64: sql.NullFloat64{
				Float64: 100.5,
				Valid:   true,
			},
		}
		Convey("When I call Float64PtrOrNil", func() {
			bptr := in.Float64PtrOrNil()
			Convey("Then I should get a false float64 value pointer", func() {
				So(bptr, ShouldNotBeNil)
				So(*bptr, ShouldEqual, 100.5)
			})
		})
	})
	Convey("Given a sqljson.NullFloat64 null value", t, func() {
		in := sqljson.NullFloat64{
			NullFloat64: sql.NullFloat64{
				Float64: 0.0,
				Valid:   false,
			},
		}
		Convey("When I call Float64PtrOrNil", func() {
			bptr := in.Float64PtrOrNil()
			Convey("Then I should get a nil pointer", func() {
				So(bptr, ShouldBeNil)
			})
		})
	})
}
func TestFloat64MarshalJSON(t *testing.T) {
	Convey("Given a non-null sqljson.NullFloat64", t, func() {
		in := sqljson.NullFloat64{
			NullFloat64: sql.NullFloat64{
				Float64: 100.5,
				Valid:   true,
			},
		}
		Convey("When marshalled to JSON", func() {
			b, err := in.MarshalJSON()
			Convey("Then I should get a valid value", func() {
				So(err, ShouldBeNil)
				So(string(b), ShouldEqual, "100.5")
			})
		})
	})
	Convey("Given a null sqljson.NullFloat64", t, func() {
		in := sqljson.NullFloat64{
			NullFloat64: sql.NullFloat64{
				Float64: 0.0,
				Valid:   false,
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

func TestFloat64UnmarshalJSON(t *testing.T) {
	Convey("Given a non-zero JSON string", t, func() {
		obj := sqljson.NullFloat64{}
		jsonIn := `123.5`
		Convey("When unmarshalled", func() {
			err := obj.UnmarshalJSON([]byte(jsonIn))
			Convey("Then I should get the corresponding value back", func() {
				So(err, ShouldBeNil)
				So(obj.Valid, ShouldBeTrue)
				So(obj.Float64, ShouldEqual, 123.5)
			})
		})
	})
	Convey("Given a zero JSON string", t, func() {
		obj := sqljson.NullFloat64{}
		jsonIn := `0.0`
		Convey("When unmarshalled", func() {
			err := obj.UnmarshalJSON([]byte(jsonIn))
			Convey("Then I should get the corresponding value back", func() {
				So(err, ShouldBeNil)
				So(obj.Valid, ShouldBeTrue)
				So(obj.Float64, ShouldEqual, 0.0)
			})
		})
	})
	Convey("Given a non-zero integer JSON string", t, func() {
		obj := sqljson.NullFloat64{}
		jsonIn := `123`
		Convey("When unmarshalled", func() {
			err := obj.UnmarshalJSON([]byte(jsonIn))
			Convey("Then I should get the corresponding value back", func() {
				So(err, ShouldBeNil)
				So(obj.Valid, ShouldBeTrue)
				So(obj.Float64, ShouldEqual, 123)
			})
		})
	})
	Convey("Given a zero integer JSON string", t, func() {
		obj := sqljson.NullFloat64{}
		jsonIn := `0`
		Convey("When unmarshalled", func() {
			err := obj.UnmarshalJSON([]byte(jsonIn))
			Convey("Then I should get the corresponding value back", func() {
				So(err, ShouldBeNil)
				So(obj.Valid, ShouldBeTrue)
				So(obj.Float64, ShouldEqual, 0)
			})
		})
	})
	Convey("Given a null JSON string", t, func() {
		obj := sqljson.NullFloat64{}
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
		errorCaseDef{"a bool true", `true`, "cannot unmarshal bool into Go value of type float64"},
		errorCaseDef{"a bool false", `false`, "cannot unmarshal bool into Go value of type float64"},
		errorCaseDef{"a string", `"Hello!"`, "cannot unmarshal string into Go value of type float64"},
		errorCaseDef{"an object", `{"status":"ok"}`, "cannot unmarshal object into Go value of type float64"},
		errorCaseDef{"an array", `[1,2,3]`, "cannot unmarshal array into Go value of type float64"},
	}
	for _, errorCase := range errorCases {
		Convey("Given "+errorCase.Label+" JSON string", t, func() {
			obj := sqljson.NullFloat64{}
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
