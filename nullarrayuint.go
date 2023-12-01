package nulltype

import (
	//	"bytes"
	//	"database/sql"
	"database/sql/driver"
	"strings"

	//	"strconv"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

/* SQL and JSon null.ArrUint */

type ArrUint[T uint8 | uint16 | uint32 | uint64] struct {
	Array []T
	Valid bool
}

func NewArrUint[T uint8 | uint16 | uint32 | uint64](i []T) ArrUint[T] {
	n := ArrUint[T]{}
	n.Valid = true
	n.Array = i
	return n
}

func (ni *ArrUint[T]) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	val := (*ArrUint[T])(ptr)

	if val.Valid {
		stream.WriteVal(val.Array)
	} else {
		stream.WriteVal(nil)
	}
}

// IsEmpty detect whether primitive.ObjectID is empty.
func (ni *ArrUint[T]) IsEmpty(ptr unsafe.Pointer) bool {
	val := (*ArrUint[T])(ptr)
	return !val.Valid
}

func (ni *ArrUint[T]) UnmarshalCSV(b string) error {
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
func (ni ArrUint[T]) MarshalCSV() (string, error) {
	if ni.Valid {
		b, err := json.Marshal(ni.Array)
		return string(b), err
	}
	return "", nil
}

func (ni *ArrUint[T]) UnmarshalJSON(b []byte) error {
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

func (ni ArrUint[T]) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Array)
	}
	return json.Marshal(nil)
}

func (ni *ArrUint[T]) Scan(value any) error {
	if value == nil {
		ni.Array, ni.Valid = make([]T, 0), false
		return nil
	}
	ni.Valid = true
	return convertAssignRows(&ni.Array, value)
}

func (ni ArrUint[T]) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Array, nil
}
