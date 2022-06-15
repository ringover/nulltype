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

/* SQL and JSon null.Int64 */

type Int64 sql.NullInt64

func NewInt64(i int64) Int64 {
	n := Int64{}
	n.Valid = true
	n.Int64 = i
	return n
}

func (ni *Int64) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	fmt.Println("DEBUG: custom marshal nullFloat")

	val := (*int64)(ptr)
	stream.WriteVal(val)
}

// IsEmpty detect whether primitive.ObjectID is empty.
func (ni *Int64) IsEmpty(ptr unsafe.Pointer) bool {
	if !ni.Valid {
		return true
	}
	return false
}

func (ni *Int64) UnmarshalJSON(b []byte) error {
	var i int64
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	if bytes.Compare(b, []byte("null")) == 0 {
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

func (ni *Int64) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = Int64{i.Int64, false}
	} else {
		*ni = Int64{i.Int64, true}
	}

	return nil
}

func (ni Int64) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Int64, nil
}
