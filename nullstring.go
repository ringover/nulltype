package nulltype

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
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
	fmt.Println("DEBUG: custom marshal nullFloat")

	val := (*string)(ptr)
	stream.WriteVal(val)
}

// IsEmpty detect whether primitive.ObjectID is empty.
func (ns *String) IsEmpty(ptr unsafe.Pointer) bool {
	if !ns.Valid {
		return true
	}
	return false
}

func (ns *String) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if bytes.Compare(b, []byte("null")) == 0 {
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
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ns = String{s.String, false}
	} else {
		*ns = String{s.String, true}
	}

	return nil
}

func (ns String) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.String, nil
}
