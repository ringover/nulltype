package nulltype

import (
	"database/sql/driver"
	"strings"

	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

/* SQL and JSon null.ArrayUint */

type ArrayUint[T uint8 | uint16 | uint32 | uint64] struct {
	Array []T
	Valid bool
}

func NewArrayUint[T uint8 | uint16 | uint32 | uint64](i []T) ArrayUint[T] {
	n := ArrayUint[T]{}
	n.Valid = true
	n.Array = i
	return n
}

func (ni *ArrayUint[T]) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	val := (*ArrayUint[T])(ptr)

	if val.Valid {
		stream.WriteVal(val.Array)
	} else {
		stream.WriteVal(nil)
	}
}

// IsEmpty detect whether primitive.ObjectID is empty.
func (ni *ArrayUint[T]) IsEmpty(ptr unsafe.Pointer) bool {
	val := (*ArrayUint[T])(ptr)
	return !val.Valid
}

func (ni *ArrayUint[T]) UnmarshalCSV(b string) error {
	var i []T
	if err := json.Unmarshal([]byte(b), &i); err != nil {
		return err
	}
	if strings.Compare(b, "null") == 0 {
		ni.Valid = false
		return nil
	}
	ni.Array = i
	ni.Valid = true
	return nil
}

// MarshalCSV marshals CSV
func (ni ArrayUint[T]) MarshalCSV() (string, error) {
	if ni.Valid {
		b, err := json.Marshal(ni.Array)
		return string(b), err
	}
	return "", nil
}

func (ni *ArrayUint[T]) UnmarshalJSON(b []byte) error {
	var i []T
	if err := json.Unmarshal([]byte(b), &i); err != nil {
		return err
	}
	if strings.Compare(string(b), "null") == 0 {
		ni.Valid = false
		return nil
	}
	ni.Array = i
	ni.Valid = true

	return nil
}

func (ni ArrayUint[T]) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Array)
	}
	return json.Marshal(nil)
}

func (ni *ArrayUint[T]) Scan(value any) error {
	if value == nil {
		ni.Array, ni.Valid = make([]T, 0), false
		return nil
	}
	ni.Valid = true
	return convertAssignRows(&ni.Array, value)
}

func (ni ArrayUint[T]) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Array, nil
}
