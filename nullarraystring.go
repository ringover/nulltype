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

/* SQL and JSon null.ArrString */

type ArrString struct {
	Array []string
	Valid bool
}

func NewArrString(i []string) ArrString {
	n := ArrString{}
	n.Valid = true
	n.Array = i
	return n
}

func (ni *ArrString) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	val := (*ArrString)(ptr)

	if val.Valid {
		stream.WriteVal(val.Array)
	} else {
		stream.WriteVal(nil)
	}
}

// IsEmpty detect whether primitive.ObjectID is empty.
func (ni *ArrString) IsEmpty(ptr unsafe.Pointer) bool {
	val := (*ArrString)(ptr)
	return !val.Valid
}

func (ni *ArrString) UnmarshalCSV(b string) error {
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
func (ni ArrString) MarshalCSV() (string, error) {
	if ni.Valid {
		b, err := json.Marshal(ni.Array)
		return string(b), err
	}
	return "", nil
}

func (ni *ArrString) UnmarshalJSON(b []byte) error {
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

func (ni ArrString) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Array)
	}
	return json.Marshal(nil)
}

func (ni *ArrString) Scan(value any) error {
	if value == nil {
		ni.Array, ni.Valid = make([]string, 0), false
		return nil
	}
	ni.Valid = true
	return convertAssignRows(&ni.Array, value)
}

func (ni ArrString) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Array, nil
}
