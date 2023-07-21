package nulltype

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"time"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

/* SQL and JSon null.Time
 * https://github.com/go-sql-driver/mysql/blob/master/utils.go
 * Patch for JSON Marshal & Unmarshal
 */

var timeFormat string = "2006-01-02 15:04:05.999999"

type Time struct {
	Time     time.Time
	Timezone string
	Valid    bool
	/* Valid is true if Time is not NULL */
}

func NewTime(t time.Time) Time {
	n := Time{}
	n.Valid = true
	n.Time = t
	n.Timezone = "UTC"
	return n
}

func (n *Time) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	val := (*time.Time)(ptr)
	stream.WriteVal(val)
}

// IsEmpty detect whether primitive.ObjectID is empty.
func (n *Time) IsEmpty(ptr unsafe.Pointer) bool {
	val := (*Time)(ptr)
	return !val.Valid
}

func (n Time) String() string {
	if n.Valid {
		b, _ := n.Time.MarshalText()
		return string(b)
	}
	return ""
}

func (n *Time) UnmarshalCSV(b string) error {
	var t time.Time
	/* When we received an empty timestamp */
	if len(b) <= 2 {
		n.Time = time.Time{}
		n.Valid = true
		return nil
	}
	n.Timezone = "UTC"
	if bytes.Equal([]byte(b), []byte("null")) {
		n.Valid = false
		return nil
	}
	if err := json.Unmarshal([]byte(b), &t); err != nil {
		return err
	}
	n.Time = t
	n.Valid = true
	return nil
}

// MarshalCSV marshals CSV
func (n Time) MarshalCSV() (string, error) {
	if n.Valid {
		location, err := time.LoadLocation(n.Timezone)
		if err != nil {
			location = time.UTC
		}
		s := n.Time.In(location).Format("2006-01-02 15:04:05")
		return s, nil
	}
	return "", nil
}

func (n *Time) UnmarshalJSON(b []byte) error {
	var t time.Time
	/* When we received an empty timestamp */
	if len(b) <= 2 {
		n.Time = time.Time{}
		n.Valid = true
		return nil
	}
	if bytes.Equal(b, []byte("null")) {
		n.Valid = false
		return nil
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	n.Timezone = "UTC"
	n.Time = t
	n.Valid = true
	return nil
}

func (n Time) MarshalText() ([]byte, error) {
	if n.Valid {
		location, err := time.LoadLocation(n.Timezone)
		if err != nil {
			location = time.UTC
		}
		return json.Marshal(n.Time.In(location))
	}
	return json.Marshal(nil)
}

func (n Time) MarshalJSON() ([]byte, error) {
	if n.Valid {
		location, err := time.LoadLocation(n.Timezone)
		if err != nil {
			location = time.UTC
		}
		return json.Marshal(n.Time.In(location))
	}
	return json.Marshal(nil)
}

func parseDateTime(str string, loc *time.Location) (t time.Time, err error) {
	base := "0000-00-00 00:00:00.0000000"
	switch len(str) {
	case 10, 19, 21, 22, 23, 24, 25, 26: // up to "YYYY-MM-DD HH:MM:SS.MMMMMM"
		if str == base[:len(str)] {
			return
		}
		t, err = time.Parse(timeFormat[:len(str)], str)
	default:
		err = fmt.Errorf("invalid time string: %s", str)
		return
	}

	// Adjust location
	if err == nil && loc != time.UTC {
		y, mo, d := t.Date()
		h, mi, s := t.Clock()
		t, err = time.Date(y, mo, d, h, mi, s, t.Nanosecond(), loc), nil
	}

	return
}

func (nt *Time) Scan(value interface{}) (err error) {
	if value == nil {
		nt.Time, nt.Valid = time.Time{}, false
		return
	}

	switch v := value.(type) {
	case *time.Time:
		nt.Time, nt.Valid = *v, true
		return
	case time.Time:
		nt.Time, nt.Valid = v, true
		return
	case []byte:
		nt.Time, err = parseDateTime(string(v), time.UTC)
		nt.Valid = (err == nil)
		return
	case string:
		nt.Time, err = parseDateTime(v, time.UTC)
		nt.Valid = (err == nil)
		return
	}

	nt.Valid = false
	return fmt.Errorf("can't convert %T to time.Time", value)
	//nt.Valid = true
	//return convertAssignRows(&nt.Time, value)
}

func (n Time) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}
