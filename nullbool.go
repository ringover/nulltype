package nulltype

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"reflect"
	"strconv"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

/* SQL and JSon null.Bool */

type Bool sql.NullBool

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func NewBool(b bool) Bool {
	nb := Bool{}
	nb.Valid = true
	nb.Bool = b
	return nb
}

func (nb *Bool) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	val := (*Bool)(ptr)

	if val.Valid {
		stream.WriteVal(val.Bool)
	} else {
		stream.WriteVal(nil)
	}

}

// IsEmpty detect whether primitive.ObjectID is empty.
func (nb *Bool) IsEmpty(ptr unsafe.Pointer) bool {
	val := (*Bool)(ptr)
	if !val.Valid {
		return true
	}
	return false
}

func (nb *Bool) UnmarshalCSV(b string) error {
	var err error
	nb.Bool, err = strconv.ParseBool(b)
	return err
}

// MarshalCSV marshals CSV
func (nb Bool) MarshalCSV() (string, error) {
	if nb.Valid {
		return strconv.FormatBool(nb.Bool), nil
	}
	return "", nil
}

func (nb *Bool) UnmarshalJSON(b []byte) error {
	var i bool
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	if bytes.Compare(b, []byte("null")) == 0 {
		nb.Valid = false
		return nil
	}
	nb.Bool = i
	nb.Valid = true
	return nil
}

func (nb Bool) MarshalJSON() ([]byte, error) {
	if nb.Valid {
		return json.Marshal(nb.Bool)
	}
	return json.Marshal(nil)
}

func (nb *Bool) Scan(value interface{}) error {
	var b sql.NullBool
	if err := b.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nb = Bool{b.Bool, false}
	} else {
		*nb = Bool{b.Bool, true}
	}

	return nil
}

func (nb Bool) Value() (driver.Value, error) {
	if !nb.Valid {
		return nil, nil
	}
	return nb.Bool, nil
}
