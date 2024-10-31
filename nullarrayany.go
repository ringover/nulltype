package nulltype

import (
	"database/sql/driver"
	"strings"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

/* SQL and JSon null.ArrayAny */

type ArrayAny struct {
	Array []any
	Valid bool
}

func NewArrayAny(i []any) ArrayAny {
	n := ArrayAny{}
	n.Valid = true
	n.Array = i
	return n
}

func (ni *ArrayAny) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	val := (*ArrayAny)(ptr)

	if val.Valid {
		stream.WriteVal(val.Array)
	} else {
		stream.WriteVal(nil)
	}
}

// IsEmpty detect whether primitive.ObjectID is empty.
func (ni *ArrayAny) IsEmpty(ptr unsafe.Pointer) bool {
	val := (*ArrayAny)(ptr)
	return !val.Valid
}

func (ni *ArrayAny) UnmarshalCSV(b string) error {
	var i []any
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
func (ni ArrayAny) MarshalCSV() (string, error) {
	if ni.Valid {
		b, err := json.Marshal(ni.Array)
		return string(b), err
	}
	return "", nil
}

func (ni *ArrayAny) UnmarshalJSON(b []byte) error {
	var i []any
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

func (ni ArrayAny) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Array)
	}
	return json.Marshal(nil)
}

func (ni *ArrayAny) Scan(value any) error {
	if value == nil {
		ni.Array, ni.Valid = make([]any, 0), false
		return nil
	}
	ni.Valid = true
	return convertAssignRows(&ni.Array, value)
}

func (ni ArrayAny) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Array, nil
}
