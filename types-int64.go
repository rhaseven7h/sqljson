package sqljson

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

// NullInt64 //
type NullInt64 struct {
	sql.NullInt64
}

// NullInt64ValidateValuer //
func NullInt64ValidateValuer(field reflect.Value) interface{} {
	if nullInt64, ok := field.Interface().(NullInt64); ok {
		if nullInt64.Valid {
			return nullInt64.Int64
		}
	}
	return nil
}

// Int64PtrOrNil //
func (ns NullInt64) Int64PtrOrNil() *int64 {
	if ns.Valid {
		s := ns.Int64
		return &s
	}
	return nil
}

// MarshalJSON //
func (ns NullInt64) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.Int64)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON //
func (ns *NullInt64) UnmarshalJSON(data []byte) error {
	value := new(int64)
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}
	if value != nil {
		ns.Int64 = *value
		ns.Valid = true
	} else {
		ns.Int64 = 0
		ns.Valid = false
	}
	return nil
}
