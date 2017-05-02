package sqljson_test

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/Rhaseven7h/sqljson"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNullStringValidateValuer(t *testing.T) {
	Convey("Given a Nil sqljson.NullString value", t, func() {
		nullStringIn := sqljson.NullString{
			NullString: sql.NullString{
				String: "",
				Valid:  false,
			},
		}
		Convey("When I get its value using NullStringValidateValuer", func() {
			nullStringOut := sqljson.NullStringValidateValuer(reflect.ValueOf(interface{}(nullStringIn)))
			Convey("Then I should get nil", func() {
				So(nullStringOut, ShouldBeNil)
			})
		})
	})
	Convey("Given an empty string sqljson.NullString value", t, func() {
		nullStringIn := sqljson.NullString{
			NullString: sql.NullString{
				String: "",
				Valid:  true,
			},
		}
		Convey("When I get its value using NullStringValidateValuer", func() {
			nullStringOut := sqljson.NullStringValidateValuer(reflect.ValueOf(interface{}(nullStringIn)))
			Convey("Then I should get an empty string", func() {
				str, ok := nullStringOut.(string)
				So(ok, ShouldBeTrue)
				So(str, ShouldEqual, "")
			})
		})
	})
	Convey("Given a non-empty string sqljson.NullString value", t, func() {
		nullStringIn := sqljson.NullString{
			NullString: sql.NullString{
				String: "dummy",
				Valid:  true,
			},
		}
		Convey("When I get its value using NullStringValidateValuer", func() {
			nullStringOut := sqljson.NullStringValidateValuer(reflect.ValueOf(interface{}(nullStringIn)))
			Convey("Then I should get the non-empty string value", func() {
				str, ok := nullStringOut.(string)
				So(ok, ShouldBeTrue)
				So(str, ShouldEqual, "dummy")
			})
		})
	})
}

func TestStringMarshalJSON(t *testing.T) {
	Convey("Given a non-empty sqljson.NullString string value", t, func() {
		valueIn := sqljson.NullString{
			NullString: sql.NullString{
				String: "dummy",
				Valid:  true,
			},
		}
		Convey("When I marshal it", func() {
			b, err := valueIn.MarshalJSON()
			Convey("Then I should get a string", func() {
				So(err, ShouldBeNil)
				So(string(b), ShouldEqual, `"dummy"`)
			})
		})
	})
	Convey("Given an empty sqljson.NullString string value", t, func() {
		valueIn := sqljson.NullString{
			NullString: sql.NullString{
				String: "",
				Valid:  true,
			},
		}
		Convey("When I marshal it", func() {
			b, err := valueIn.MarshalJSON()
			Convey("Then I should get an empty string", func() {
				So(err, ShouldBeNil)
				So(string(b), ShouldEqual, `""`)
			})
		})
	})
	Convey("Given a nil sqljson.NullString string value", t, func() {
		valueIn := sqljson.NullString{
			NullString: sql.NullString{
				String: "",
				Valid:  false,
			},
		}
		Convey("When I marshal it", func() {
			b, err := valueIn.MarshalJSON()
			Convey("Then I should get null", func() {
				So(err, ShouldBeNil)
				So(string(b), ShouldEqual, `null`)
			})
		})
	})
}

func TestStringUnmarshalJSON(t *testing.T) {
	Convey("Given a sqljson.NullString value pointer, and an empty JSON value (empty, not just null)", t, func() {
		ns := &sqljson.NullString{}
		bJSON := []byte(``)
		So(ns.Valid, ShouldBeFalse)
		Convey("When I Unmarshal it using UnmarshalJSON", func() {
			err := ns.UnmarshalJSON(bJSON)
			Convey("Then I should get an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "unexpected end of JSON input")
			})
		})
	})
	Convey("Given a sqljson.NullString value pointer, and a null JSON value", t, func() {
		ns := &sqljson.NullString{}
		bJSON := []byte(`null`)
		So(ns.Valid, ShouldBeFalse)
		Convey("When I Unmarshal it using UnmarshalJSON", func() {
			err := ns.UnmarshalJSON(bJSON)
			Convey("Then I should get a Null NullString", func() {
				So(err, ShouldBeNil)
				So(ns.Valid, ShouldBeFalse)
			})
		})
	})
	Convey("Given a sqljson.NullString value pointer, and an empty string JSON value", t, func() {
		ns := &sqljson.NullString{}
		bJSON := []byte(`""`)
		So(ns.Valid, ShouldBeFalse)
		Convey("When I Unmarshal it using UnmarshalJSON", func() {
			err := ns.UnmarshalJSON(bJSON)
			Convey("Then I should get a an empty string, not-Null NullString", func() {
				So(err, ShouldBeNil)
				So(ns.Valid, ShouldBeTrue)
				So(ns.String, ShouldEqual, "")
			})
		})
	})
	Convey("Given a sqljson.NullString value pointer, and a non-empty string JSON value", t, func() {
		ns := &sqljson.NullString{}
		bJSON := []byte(`"Hello World!"`)
		So(ns.Valid, ShouldBeFalse)
		Convey("When I Unmarshal it using UnmarshalJSON", func() {
			err := ns.UnmarshalJSON(bJSON)
			Convey("Then I should get a a non-empty string, not-Null NullString", func() {
				So(err, ShouldBeNil)
				So(ns.Valid, ShouldBeTrue)
				So(ns.String, ShouldEqual, "Hello World!")
			})
		})
	})
	Convey("Given a sqljson.NullString value pointer, and an invalid JSON value", t, func() {
		ns := &sqljson.NullString{}
		bJSON := []byte(`xyz`)
		So(ns.Valid, ShouldBeFalse)
		Convey("When I Unmarshal it using UnmarshalJSON", func() {
			err := ns.UnmarshalJSON(bJSON)
			Convey("Then I should get an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "invalid character")
			})
		})
	})
	Convey("Given a sqljson.NullString value pointer, and a boolean true JSON value", t, func() {
		ns := &sqljson.NullString{}
		bJSON := []byte(`true`)
		So(ns.Valid, ShouldBeFalse)
		Convey("When I Unmarshal it using UnmarshalJSON", func() {
			err := ns.UnmarshalJSON(bJSON)
			Convey("Then I should get an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "cannot unmarshal bool")
			})
		})
	})
	Convey("Given a sqljson.NullString value pointer, and a boolean false JSON value", t, func() {
		ns := &sqljson.NullString{}
		bJSON := []byte(`false`)
		So(ns.Valid, ShouldBeFalse)
		Convey("When I Unmarshal it using UnmarshalJSON", func() {
			err := ns.UnmarshalJSON(bJSON)
			Convey("Then I should get an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "cannot unmarshal bool")
			})
		})
	})
	Convey("Given a sqljson.NullString value pointer, and an integer zero JSON value", t, func() {
		ns := &sqljson.NullString{}
		bJSON := []byte(`0`)
		So(ns.Valid, ShouldBeFalse)
		Convey("When I Unmarshal it using UnmarshalJSON", func() {
			err := ns.UnmarshalJSON(bJSON)
			Convey("Then I should get an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "cannot unmarshal number")
			})
		})
	})
	Convey("Given a sqljson.NullString value pointer, and an integer non-zero JSON value", t, func() {
		ns := &sqljson.NullString{}
		bJSON := []byte(`10`)
		So(ns.Valid, ShouldBeFalse)
		Convey("When I Unmarshal it using UnmarshalJSON", func() {
			err := ns.UnmarshalJSON(bJSON)
			Convey("Then I should get an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "cannot unmarshal number")
			})
		})
	})
	Convey("Given a sqljson.NullString value pointer, and a float zero JSON value", t, func() {
		ns := &sqljson.NullString{}
		bJSON := []byte(`0.0`)
		So(ns.Valid, ShouldBeFalse)
		Convey("When I Unmarshal it using UnmarshalJSON", func() {
			err := ns.UnmarshalJSON(bJSON)
			Convey("Then I should get an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "cannot unmarshal number")
			})
		})
	})
	Convey("Given a sqljson.NullString value pointer, and a float non-zero JSON value", t, func() {
		ns := &sqljson.NullString{}
		bJSON := []byte(`123.45`)
		So(ns.Valid, ShouldBeFalse)
		Convey("When I Unmarshal it using UnmarshalJSON", func() {
			err := ns.UnmarshalJSON(bJSON)
			Convey("Then I should get an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "cannot unmarshal number")
			})
		})
	})
}

func TestStringPtrOrNil(t *testing.T) {
	Convey("Given a valid non-Null NullString value", t, func() {
		value := sqljson.NullString{
			NullString: sql.NullString{
				String: "somemail@someserver.net",
				Valid:  true,
			},
		}
		Convey("When we get the StringPtrOrNil", func() {
			res := value.StringPtrOrNil()
			Convey("Then we get the string value", func() {
				So(res, ShouldNotBeNil)
				So(*res, ShouldEqual, "somemail@someserver.net")
			})
		})
	})
	Convey("Given a valid null NullString value", t, func() {
		value := sqljson.NullString{
			NullString: sql.NullString{
				String: "",
				Valid:  false,
			},
		}
		Convey("When we get the StringPtrOrNil", func() {
			res := value.StringPtrOrNil()
			Convey("Then we get the string value", func() {
				So(res, ShouldBeNil)
			})
		})
	})
}
