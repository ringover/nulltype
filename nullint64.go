package nulltype

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"strconv"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

/* SQL and JSon null.Int64 */

type Int64 sql.NullInt64

func NewInt64(i int64) Int64 {
	n := Int64{}
	n.Valid = true
	n.Int64 = i
	return n
}

func (ni *Int64) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	val := (*int64)(ptr)
	stream.WriteVal(val)
}

// IsEmpty detect whether primitive.ObjectID is empty.
func (ni *Int64) IsEmpty(ptr unsafe.Pointer) bool {
	val := (*Int64)(ptr)
	return !val.Valid
}

func (ni *Int64) UnmarshalCSV(b string) error {
	var err error
	ni.Int64, err = strconv.ParseInt(b, 10, 64)
	return err
}

// MarshalCSV marshals CSV
func (ni Int64) MarshalCSV() (string, error) {
	if ni.Valid {
		return strconv.FormatInt(ni.Int64, 10), nil
	}
	return "", nil
}

func (ni *Int64) UnmarshalJSON(b []byte) error {
	var i int64
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	if bytes.Equal(b, []byte("null")) {
		ni.Valid = false
		return nil
	}
	ni.Int64 = i
	ni.Valid = true
	return nil
}

func (ni Int64) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int64)
	}
	return json.Marshal(nil)
}

func (ni *Int64) Scan(value any) error {
	if value == nil {
		ni.Int64, ni.Valid = 0, false
		return nil
	}
	ni.Valid = true
	return convertAssignRows(&ni.Int64, value)
}

func (ni Int64) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Int64, nil
}
