package nulltype

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"strconv"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

/* SQL and JSon null.Int16 */

type Int16 sql.NullInt16

func NewInt16(i int16) Int16 {
	n := Int16{}
	n.Valid = true
	n.Int16 = i
	return n
}

func (ni *Int16) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	val := (*Int16)(ptr)

	if val.Valid {
		stream.WriteVal(val.Int16)
	} else {
		stream.WriteVal(nil)
	}
}

// IsEmpty detect whether primitive.ObjectID is empty.
func (ni *Int16) IsEmpty(ptr unsafe.Pointer) bool {
	val := (*Int16)(ptr)
	return !val.Valid
}

func (ni *Int16) UnmarshalCSV(b string) error {
	itmp, err := strconv.ParseInt(b, 10, 16)
	ni.Int16 = int16(itmp)
	return err
}

// MarshalCSV marshals CSV
func (ni Int16) MarshalCSV() (string, error) {
	if ni.Valid {
		return strconv.FormatInt(int64(ni.Int16), 10), nil
	}
	return "", nil
}

func (ni *Int16) UnmarshalJSON(b []byte) error {
	var i int16
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	if bytes.Equal(b, []byte("null")) {
		ni.Valid = false
		return nil
	}
	ni.Int16 = i
	ni.Valid = true
	return nil
}

func (ni Int16) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int16)
	}
	return json.Marshal(nil)
}

func (ni *Int16) Scan(value any) error {
	if value == nil {
		ni.Int16, ni.Valid = 0, false
		return nil
	}
	ni.Valid = true
	return convertAssignRows(&ni.Int16, value)
}

func (ni Int16) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Int16, nil
}
