package nulltype

import (
	"bytes"
	"database/sql/driver"
	"strconv"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

/* SQL and JSon null.Uint64 */

type Uint64 struct {
	Uint64 uint64
	Valid  bool // Valid is true if Uint64 is not NULL
}

func NewUint64(i uint64) Uint64 {
	n := Uint64{}
	n.Valid = true
	n.Uint64 = i
	return n
}

func (ni *Uint64) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	val := (*Uint64)(ptr)
	if val.Valid {
		stream.WriteVal(val.Uint64)
	} else {
		stream.WriteVal(nil)
	}

}

// IsEmpty detect whether primitive.ObjectID is empty.
func (ni *Uint64) IsEmpty(ptr unsafe.Pointer) bool {
	val := (*Uint64)(ptr)
	return !val.Valid
}

func (ni *Uint64) UnmarshalCSV(b string) error {
	var err error
	ni.Uint64, err = strconv.ParseUint(b, 10, 64)
	return err
}

// MarshalCSV marshals CSV
func (ni Uint64) MarshalCSV() (string, error) {
	if ni.Valid {
		return strconv.FormatUint(ni.Uint64, 10), nil
	}
	return "", nil
}

func (ni *Uint64) UnmarshalJSON(b []byte) error {
	var i uint64
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	if bytes.Equal(b, []byte("null")) {
		ni.Valid = false
		return nil
	}
	ni.Uint64 = i
	ni.Valid = true
	return nil
}

func (ni Uint64) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Uint64)
	}
	return json.Marshal(nil)
}

func (ni *Uint64) Scan(value any) error {
	if value == nil {
		ni.Uint64, ni.Valid = 0, false
		return nil
	}
	ni.Valid = true
	return convertAssignRows(&ni.Uint64, value)
}

func (ni Uint64) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Uint64, nil
}
