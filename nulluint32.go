package nulltype

import (
	"bytes"
	"database/sql/driver"
	"strconv"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

/* SQL and JSon null.Uint32 */

type Uint32 struct {
	Uint32 uint32
	Valid  bool // Valid is true if Uint32 is not NULL
}

func NewUint32(i uint32) Uint32 {
	n := Uint32{}
	n.Valid = true
	n.Uint32 = i
	return n
}

func (ni *Uint32) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	val := (*uint32)(ptr)
	stream.WriteVal(val)
}

// IsEmpty detect whether primitive.ObjectID is empty.
func (ni *Uint32) IsEmpty(ptr unsafe.Pointer) bool {
	val := (*Uint32)(ptr)
	return !val.Valid
}

func (ni *Uint32) UnmarshalCSV(b string) error {
	uitmp, err := strconv.ParseUint(b, 10, 32)
	ni.Uint32 = uint32(uitmp)
	return err
}

// MarshalCSV marshals CSV
func (ni Uint32) MarshalCSV() (string, error) {
	if ni.Valid {
		return strconv.FormatUint(uint64(ni.Uint32), 10), nil
	}
	return "", nil
}

func (ni *Uint32) UnmarshalJSON(b []byte) error {
	var i uint32
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	if bytes.Equal(b, []byte("null")) {
		ni.Valid = false
		return nil
	}
	ni.Uint32 = i
	ni.Valid = true
	return nil
}

func (ni Uint32) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Uint32)
	}
	return json.Marshal(nil)
}

func (ni *Uint32) Scan(value any) error {
	if value == nil {
		ni.Uint32, ni.Valid = 0, false
		return nil
	}
	ni.Valid = true
	return convertAssignRows(&ni.Uint32, value)
}

func (ni Uint32) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Uint32, nil
}
