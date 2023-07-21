package nulltype

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"strconv"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

/* SQL and JSon null.Int32 */

type Int32 sql.NullInt32

func NewInt32(i int32) Int32 {
	n := Int32{}
	n.Valid = true
	n.Int32 = i
	return n
}

func (ni *Int32) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	val := (*int32)(ptr)
	stream.WriteVal(val)
}

// IsEmpty detect whether primitive.ObjectID is empty.
func (ni *Int32) IsEmpty(ptr unsafe.Pointer) bool {
	val := (*Int32)(ptr)
	return !val.Valid
}

func (ni *Int32) UnmarshalCSV(b string) error {
	itmp, err := strconv.ParseInt(b, 10, 32)
	ni.Int32 = int32(itmp)
	return err
}

// MarshalCSV marshals CSV
func (ni Int32) MarshalCSV() (string, error) {
	if ni.Valid {
		return strconv.FormatInt(int64(ni.Int32), 10), nil
	}
	return "", nil
}

func (ni *Int32) UnmarshalJSON(b []byte) error {
	var i int32
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	if bytes.Equal(b, []byte("null")) {
		ni.Valid = false
		return nil
	}
	ni.Int32 = i
	ni.Valid = true
	return nil
}

func (ni Int32) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int32)
	}
	return json.Marshal(nil)
}

func (ni *Int32) Scan(value any) error {
	if value == nil {
		ni.Int32, ni.Valid = 0, false
		return nil
	}
	ni.Valid = true
	return convertAssignRows(&ni.Int32, value)
}

func (ni Int32) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Int32, nil
}
