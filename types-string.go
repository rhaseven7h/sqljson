package sqljson

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

// NullString //
type NullString struct {
	sql.NullString
}

// NullStringValidateValuer //
func NullStringValidateValuer(field reflect.Value) interface{} {
	if nullString, ok := field.Interface().(NullString); ok {
		if nullString.Valid {
			return nullString.String
		}
	}
	return nil
}

// StringPtrOrNil //
func (ns NullString) StringPtrOrNil() *string {
	if ns.Valid {
		s := ns.String
		return &s
	}
	return nil
}

// MarshalJSON //
func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON //
func (ns *NullString) UnmarshalJSON(data []byte) error {
	value := new(string)
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}
	if value != nil {
		ns.String = *value
		ns.Valid = true
	} else {
		ns.String = ""
		ns.Valid = false
	}
	return nil
}
