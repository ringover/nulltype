package nulltype

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

/* SQL and JSon null.String */

type String sql.NullString

func NewString(str string) String {
	n := String{}
	n.Valid = true
	n.String = str
	return n
}

func (ns *String) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	val := (*String)(ptr)

	if val.Valid {
		stream.WriteVal(val.String)
	} else {
		stream.WriteVal(nil)
	}
}

// IsEmpty detect whether primitive.ObjectID is empty.
func (ns *String) IsEmpty(ptr unsafe.Pointer) bool {
	val := (*String)(ptr)
	return !val.Valid
}

func (ns *String) UnmarshalCSV(b string) error {
	var err error
	ns.String = b
	return err
}

// MarshalCSV marshals CSV
func (ns String) MarshalCSV() (string, error) {
	if ns.Valid {
		return ns.String, nil
	}
	return "", nil
}

func (ns *String) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if bytes.Equal(b, []byte("null")) {
		ns.Valid = false
		return nil
	}
	ns.String = s
	ns.Valid = true
	return nil
}

func (ns String) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return json.Marshal(nil)
}

func (ns *String) Scan(value interface{}) error {
	if value == nil {
		ns.String, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return convertAssignRows(&ns.String, value)
}

func (ns String) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.String, nil
}
