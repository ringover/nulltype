package nulltype

import (
	"bytes"
	"database/sql/driver"
	"strconv"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

/* SQL and JSon null.Uint8 */

type Uint8 struct {
	Uint8 uint8
	Valid bool // Valid is true if Uint8 is not NULL
}

func NewUint8(i uint8) Uint8 {
	n := Uint8{}
	n.Valid = true
	n.Uint8 = i
	return n
}

func (ni *Uint8) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	val := (*uint8)(ptr)
	stream.WriteVal(val)
}

// IsEmpty detect whether primitive.ObjectID is empty.
func (ni *Uint8) IsEmpty(ptr unsafe.Pointer) bool {
	val := (*Uint8)(ptr)
	return !val.Valid
}

func (ni *Uint8) UnmarshalCSV(b string) error {
	uitmp, err := strconv.ParseUint(b, 10, 8)
	ni.Uint8 = uint8(uitmp)
	return err
}

// MarshalCSV marshals CSV
func (ni Uint8) MarshalCSV() (string, error) {
	if ni.Valid {
		return strconv.FormatUint(uint64(ni.Uint8), 10), nil
	}
	return "", nil
}

func (ni *Uint8) UnmarshalJSON(b []byte) error {
	var i uint8
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	if bytes.Equal(b, []byte("null")) {
		ni.Valid = false
		return nil
	}
	ni.Uint8 = i
	ni.Valid = true
	return nil
}

func (ni Uint8) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Uint8)
	}
	return json.Marshal(nil)
}

func (ni *Uint8) Scan(value any) error {
	if value == nil {
		ni.Uint8, ni.Valid = 0, false
		return nil
	}
	ni.Valid = true
	return convertAssignRows(&ni.Uint8, value)
}

func (ni Uint8) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Uint8, nil
}
