package nulltype

import (
	"database/sql/driver"
	"strings"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

/* SQL and JSon null.ArrayInt */

type ArrayInt[T int8 | int16 | int32 | int64] struct {
	Array []T
	Valid bool
}

func NewArrayInt[T int8 | int16 | int32 | int64](i []T) ArrayInt[T] {
	n := ArrayInt[T]{}
	n.Valid = true
	n.Array = i
	return n
}

func (ni *ArrayInt[T]) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	val := (*ArrayInt[T])(ptr)

	if val.Valid {
		stream.WriteVal(val.Array)
	} else {
		stream.WriteVal(nil)
	}
}

// IsEmpty detect whether primitive.ObjectID is empty.
func (ni *ArrayInt[T]) IsEmpty(ptr unsafe.Pointer) bool {
	val := (*ArrayInt[T])(ptr)
	return !val.Valid
}

func (ni *ArrayInt[T]) UnmarshalCSV(b string) error {
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
func (ni ArrayInt[T]) MarshalCSV() (string, error) {
	if ni.Valid {
		b, err := json.Marshal(ni.Array)
		return string(b), err
	}
	return "", nil
}

func (ni *ArrayInt[T]) UnmarshalJSON(b []byte) error {
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

func (ni ArrayInt[T]) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Array)
	}
	return json.Marshal(nil)
}

func (ni *ArrayInt[T]) Scan(value any) error {
	if value == nil {
		ni.Array, ni.Valid = make([]T, 0), false
		return nil
	}
	ni.Valid = true
	return convertAssignRows(&ni.Array, value)
}

func (ni ArrayInt[T]) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Array, nil
}
