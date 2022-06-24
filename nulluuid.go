package nulltype

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"unsafe"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
)

type UUID struct {
	UUID uuid.UUID
	/* Valid is true if UUID is not NULL*/
	Valid bool
}

// NewUUID creates a valid null.UUID using a uuid.UUID
func NewUUID(u uuid.UUID) UUID {
	n := UUID{}
	n.Valid = true
	n.UUID = u
	return n
}

func (n *UUID) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	val := (*uuid.UUID)(ptr)
	stream.WriteVal(val)
}

// IsEmpty detect whether primitive.ObjectID is empty.
func (n *UUID) IsEmpty(ptr unsafe.Pointer) bool {
	val := (*UUID)(ptr)
	if !val.Valid {
		return true
	}
	return false
}

func (n *UUID) UnmarshalCSV(b string) error {
	var err error
	n.UUID, err = uuid.FromBytes([]byte(b))
	return err
}

// MarshalCSV marshals CSV
func (n *UUID) MarshalCSV() (string, error) {
	if n.Valid {
		return n.UUID.String(), nil
	}
	return "", nil
}

func (n *UUID) UnmarshalJSON(b []byte) error {
	var u uuid.UUID
	if bytes.Compare(b, []byte("null")) == 0 {
		n.Valid = false
		return nil
	}
	u, err := uuid.ParseBytes(b)
	if err != nil {
		n.Valid = false
		return err
	}
	n.UUID = u
	n.Valid = true
	return nil
}

func (n UUID) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.UUID)
	}
	return json.Marshal(nil)
}

func (n *UUID) Scan(value interface{}) (err error) {
	if value == nil {
		n.UUID, n.Valid = uuid.UUID{}, false
		return
	}

	switch v := value.(type) {
	case uuid.UUID:
		n.UUID, n.Valid = v, true
		return
	case []byte:
		n.UUID, err = uuid.ParseBytes(v)
		n.Valid = true
		return
	case string:
		n.UUID, err = uuid.Parse(v)
		n.Valid = true
		return
	}

	n.Valid = false
	return fmt.Errorf("Can't convert %T to uuid.UUID", value)
}

func (n UUID) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.UUID, nil
}
