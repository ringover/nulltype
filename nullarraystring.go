package nulltype

import (
	"database/sql/driver"
	"strings"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

/* SQL and JSon null.ArrayString */

type ArrayString struct {
	Array []string
	Valid bool
}

func NewArrayString(i []string) ArrayString {
	n := ArrayString{}
	n.Valid = true
	n.Array = i
	return n
}

func (ni *ArrayString) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	val := (*ArrayString)(ptr)

	if val.Valid {
		stream.WriteVal(val.Array)
	} else {
		stream.WriteVal(nil)
	}
}

// IsEmpty detect whether primitive.ObjectID is empty.
func (ni *ArrayString) IsEmpty(ptr unsafe.Pointer) bool {
	val := (*ArrayString)(ptr)
	return !val.Valid
}

func (ni *ArrayString) UnmarshalCSV(b string) error {
	var i []string
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
func (ni ArrayString) MarshalCSV() (string, error) {
	if ni.Valid {
		b, err := json.Marshal(ni.Array)
		return string(b), err
	}
	return "", nil
}

func (ni *ArrayString) UnmarshalJSON(b []byte) error {
	var i []string
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

func (ni ArrayString) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Array)
	}
	return json.Marshal(nil)
}

func (ni *ArrayString) Scan(value any) error {
	if value == nil {
		ni.Array, ni.Valid = make([]string, 0), false
		return nil
	}
	ni.Valid = true
	return convertAssignRows(&ni.Array, value)
}

func (ni ArrayString) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Array, nil
}
