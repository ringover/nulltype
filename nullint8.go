package nulltype

import (
	"bytes"
	"database/sql/driver"
	"strconv"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

/* SQL and JSon null.Int8 */

type Int8 struct {
	Int8  int8
	Valid bool // Valid is true if Uint8 is not NULL
}

func NewInt8(i int8) Int8 {
	n := Int8{}
	n.Valid = true
	n.Int8 = i
	return n
}

func (ni *Int8) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	val := (*int8)(ptr)
	stream.WriteVal(val)
}

// IsEmpty detect whether primitive.ObjectID is empty.
func (ni *Int8) IsEmpty(ptr unsafe.Pointer) bool {
	val := (*Int8)(ptr)
	return !val.Valid
}

func (ni *Int8) UnmarshalCSV(b string) error {
	itmp, err := strconv.ParseInt(b, 10, 16)
	ni.Int8 = int8(itmp)
	return err
}

// MarshalCSV marshals CSV
func (ni Int8) MarshalCSV() (string, error) {
	if ni.Valid {
		return strconv.FormatInt(int64(ni.Int8), 10), nil
	}
	return "", nil
}

func (ni *Int8) UnmarshalJSON(b []byte) error {
	var i int8
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	if bytes.Equal(b, []byte("null")) {
		ni.Valid = false
		return nil
	}
	ni.Int8 = i
	ni.Valid = true
	return nil
}

func (ni Int8) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int8)
	}
	return json.Marshal(nil)
}

func (ni *Int8) Scan(value any) error {
	if value == nil {
		ni.Int8, ni.Valid = 0, false
		return nil
	}
	ni.Valid = true
	return convertAssignRows(&ni.Int8, value)
}

func (ni Int8) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Int8, nil
}
