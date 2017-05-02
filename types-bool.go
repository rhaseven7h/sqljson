package sqljson

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

// NullBool //
type NullBool struct {
	sql.NullBool
}

// NullBoolValidateValuer //
func NullBoolValidateValuer(field reflect.Value) interface{} {
	if nullBool, ok := field.Interface().(NullBool); ok {
		if nullBool.Valid {
			return nullBool.Bool
		}
	}
	return nil
}

// BoolPtrOrNil //
func (ns NullBool) BoolPtrOrNil() *bool {
	if ns.Valid {
		s := ns.Bool
		return &s
	}
	return nil
}

// MarshalJSON //
func (ns NullBool) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.Bool)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON //
func (ns *NullBool) UnmarshalJSON(data []byte) error {
	value := new(bool)
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}
	if value != nil {
		ns.Bool = *value
		ns.Valid = true
	} else {
		ns.Bool = false
		ns.Valid = false
	}
	return nil
}
