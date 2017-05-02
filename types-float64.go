package sqljson

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

// NullFloat64 //
type NullFloat64 struct {
	sql.NullFloat64
}

// NullFloat64ValidateValuer //
func NullFloat64ValidateValuer(field reflect.Value) interface{} {
	if nullFloat64, ok := field.Interface().(NullFloat64); ok {
		if nullFloat64.Valid {
			return nullFloat64.Float64
		}
	}
	return nil
}

// Float64PtrOrNil //
func (ns NullFloat64) Float64PtrOrNil() *float64 {
	if ns.Valid {
		s := ns.Float64
		return &s
	}
	return nil
}

// MarshalJSON //
func (ns NullFloat64) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.Float64)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON //
func (ns *NullFloat64) UnmarshalJSON(data []byte) error {
	value := new(float64)
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}
	if value != nil {
		ns.Float64 = *value
		ns.Valid = true
	} else {
		ns.Float64 = 0.0
		ns.Valid = false
	}
	return nil
}
