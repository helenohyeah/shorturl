package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// Custom null handling

// NullInt64 is an nullable int64.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type NullInt64 struct {
	sql.NullInt64
}

// NewNullInt64 creates a new NullInt64
func NewNullInt64(i int64, valid bool) NullInt64 {
	return NullInt64{
		NullInt64: sql.NullInt64{
			Int64: i,
			Valid: valid,
		},
	}
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this NullInt64 is null.
func (i NullInt64) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatInt(i.Int64, 10)), nil
}

// String function will return an empty string if invalid, or the
// string value of the Int64 (useful for text output, such as csv writing)
func (i NullInt64) String() string {
	if !i.Valid {
		return ""
	}
	return strconv.FormatInt(i.Int64, 10)
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports number and null input.
// 0 will not be considered a null Int.
// It also supports unmarshalling a sql.NullInt64.
func (i *NullInt64) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v.(type) {
	case float64:
		// Unmarshal again, directly to int64, to avoid intermediate float64
		err = json.Unmarshal(data, &i.Int64)
	case nil:
		i.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type NullInt64", reflect.TypeOf(v).Name())
	}
	i.Valid = (err == nil)
	return err
}

// NullBool is a nullable bool.
// It does not consider false values to be null.
// It will decode to null, not false, if null.
type NullBool struct {
	sql.NullBool
}

// NewNullBool creates a new NullBool
func NewNullBool(b bool, valid bool) NullBool {
	return NullBool{
		NullBool: sql.NullBool{
			Bool:  b,
			Valid: valid,
		},
	}
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this NullBool is null.
func (b NullBool) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return []byte("null"), nil
	}
	if !b.Bool {
		return []byte("false"), nil
	}
	return []byte("true"), nil
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports number and null input.
// 0 will not be considered a null Bool.
// It also supports unmarshalling a sql.NullBool.
func (b *NullBool) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case bool:
		b.Bool = x
	case nil:
		b.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type NullBool64", reflect.TypeOf(v).Name())
	}
	b.Valid = (err == nil)
	return err
}

// NullString is a nullable string. It supports SQL and JSON serialization.
// It will marshal to null if null. Blank string input will be considered null.
type NullString struct {
	sql.NullString
}

// NewNullString creates a new NullString
func NewNullString(s string, valid bool) NullString {
	return NullString{
		NullString: sql.NullString{
			String: s,
			Valid:  valid,
		},
	}
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this String is null.
func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports string and null input. Blank string input produces a null String.
// It also supports unmarshalling a sql.NullString.
func (s *NullString) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case string:
		s.String = x
	case nil:
		s.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type NullString", reflect.TypeOf(v).Name())
	}
	s.Valid = (err == nil)
	return err
}
