package nulltype

import (
	"bytes"
	"database/sql/driver"
	"strconv"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

/* SQL and JSon null.Uint16 */

type Uint16 struct {
	Uint16 uint16
	Valid  bool // Valid is true if Uint16 is not NULL
}

func NewUint16(i uint16) Uint16 {
	n := Uint16{}
	n.Valid = true
	n.Uint16 = i
	return n
}

func (ni *Uint16) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	val := (*uint16)(ptr)
	stream.WriteVal(val)
}

// IsEmpty detect whether primitive.ObjectID is empty.
func (ni *Uint16) IsEmpty(ptr unsafe.Pointer) bool {
	val := (*Uint16)(ptr)
	return !val.Valid
}

func (ni *Uint16) UnmarshalCSV(b string) error {
	uitmp, err := strconv.ParseUint(b, 10, 16)
	ni.Uint16 = uint16(uitmp)
	return err
}

// MarshalCSV marshals CSV
func (ni Uint16) MarshalCSV() (string, error) {
	if ni.Valid {
		return strconv.FormatUint(uint64(ni.Uint16), 10), nil
	}
	return "", nil
}

func (ni *Uint16) UnmarshalJSON(b []byte) error {
	var i uint16
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	if bytes.Equal(b, []byte("null")) {
		ni.Valid = false
		return nil
	}
	ni.Uint16 = i
	ni.Valid = true
	return nil
}

func (ni Uint16) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Uint16)
	}
	return json.Marshal(nil)
}

func (ni *Uint16) Scan(value any) error {
	if value == nil {
		ni.Uint16, ni.Valid = 0, false
		return nil
	}
	ni.Valid = true
	return convertAssignRows(&ni.Uint16, value)
}

func (ni Uint16) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Uint16, nil
}
