package nulltype

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"time"
)

// This code is inspired by database/sql

// RawBytes is a byte slice that holds a reference to memory owned by
// the database itself. After a Scan into a RawBytes, the slice is only
// valid until the next call to Next, Scan, or Close.
type RawBytes []byte

type decimal interface {
	decimalDecompose
	decimalCompose
}

type decimalDecompose interface {
	// Decompose returns the internal decimal state in parts.
	// If the provided buf has sufficient capacity, buf may be returned as the coefficient with
	// the value set and length set as appropriate.
	Decompose(buf []byte) (form byte, negative bool, coefficient []byte, exponent int32)
}

type decimalCompose interface {
	// Compose sets the internal decimal value from parts. If the value cannot be
	// represented then an error should be returned.
	Compose(form byte, negative bool, coefficient []byte, exponent int32) error
}

var errNilPtr = errors.New("destination pointer is nil") // embedded in descriptive error

func asBytes(buf []byte, rv reflect.Value) (b []byte, ok bool) {
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.AppendInt(buf, rv.Int(), 10), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.AppendUint(buf, rv.Uint(), 10), true
	case reflect.Float32:
		return strconv.AppendFloat(buf, rv.Float(), 'g', -1, 32), true
	case reflect.Float64:
		return strconv.AppendFloat(buf, rv.Float(), 'g', -1, 64), true
	case reflect.Bool:
		return strconv.AppendBool(buf, rv.Bool()), true
	case reflect.String:
		s := rv.String()
		return append(buf, s...), true
	}
	return
}

func asString(src any) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	}
	rv := reflect.ValueOf(src)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(rv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(rv.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 32)
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool())
	}
	return fmt.Sprintf("%v", src)
}

func strconvErr(err error) error {
	if ne, ok := err.(*strconv.NumError); ok {
		return ne.Err
	}
	return err
}

func convertAssignRows(dest, src any) error {
	// Common cases, without reflect.
	switch s := src.(type) {
	case string:
		switch d := dest.(type) {
		case *string:
			if d == nil {
				return errNilPtr
			}
			*d = s
			return nil
		case *[]byte:
			if d == nil {
				return errNilPtr
			}
			*d = []byte(s)
			return nil
		case *RawBytes:
			if d == nil {
				return errNilPtr
			}
			*d = append((*d)[:0], s...)
			return nil
		}
	case []byte: // byte = uint8
		p := src.([]byte)
		switch d := dest.(type) {
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
		case *string:
			if d == nil {
				return errNilPtr
			}
			*d = string(s)
			return nil
		case *any:
			if d == nil {
				return errNilPtr
			}
			*d = bytes.Clone(s)
			return nil
		case *[]byte:
			if d == nil {
				return errNilPtr
			}
			*d = bytes.Clone(s)
			return nil
		case *RawBytes:
			if d == nil {
				return errNilPtr
			}
			*d = s
			return nil
		}
	case time.Time:
		switch d := dest.(type) {
		case *time.Time:
			*d = s
			return nil
		case *string:
			*d = s.Format(time.RFC3339Nano)
			return nil
		case *[]byte:
			if d == nil {
				return errNilPtr
			}
			*d = []byte(s.Format(time.RFC3339Nano))
			return nil
		case *RawBytes:
			if d == nil {
				return errNilPtr
			}
			*d = s.AppendFormat((*d)[:0], time.RFC3339Nano)
			return nil
		}
	case decimalDecompose:
		switch d := dest.(type) {
		case decimalCompose:
			return d.Compose(s.Decompose(nil))
		}
	case nil:
		switch d := dest.(type) {
		case *any:
			if d == nil {
				return errNilPtr
			}
			*d = nil
			return nil
		case *[]byte:
			if d == nil {
				return errNilPtr
			}
			*d = nil
			return nil
		case *RawBytes:
			if d == nil {
				return errNilPtr
			}
			*d = nil
			return nil
		}
	case []*uint8: // for uint8 see byte
		p := src.([]*uint8)
		switch d := dest.(type) {
		case *[]uint8:
			*d = make([]uint8, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []uint16:
		p := src.([]uint16)
		switch d := dest.(type) {
		case *[]uint16:
			*d = make([]uint16, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []*uint16:
		p := src.([]*uint16)
		switch d := dest.(type) {
		case *[]uint16:
			fmt.Println("kikooo")
			*d = make([]uint16, len(p))
			for i := range p {
				pi := (p)[i]
				convertAssignRows(&(*d)[i], *pi)
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []uint32:
		p := src.([]uint32)
		switch d := dest.(type) {
		case *[]uint32:
			*d = make([]uint32, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []*uint32:
		p := src.([]*uint32)
		switch d := dest.(type) {
		case *[]uint32:
			*d = make([]uint32, len(p))
			for i := range p {
				pi := (p)[i]
				convertAssignRows(&(*d)[i], *pi)
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []uint64:
		p := src.([]uint64)
		switch d := dest.(type) {
		case *[]uint64:
			*d = make([]uint64, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []*uint64:
		p := src.([]*uint64)
		switch d := dest.(type) {
		case *[]uint64:
			*d = make([]uint64, len(p))
			for i := range p {
				pi := (p)[i]
				convertAssignRows(&(*d)[i], *pi)
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []int8:
		p := src.([]int8)
		switch d := dest.(type) {
		case *[]int8:
			*d = make([]int8, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []*int8:
		p := src.([]*int8)
		switch d := dest.(type) {
		case *[]int8:
			*d = make([]int8, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []int16:
		p := src.([]int16)
		switch d := dest.(type) {
		case *[]int16:
			*d = make([]int16, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []*int16:
		p := src.([]*int16)
		switch d := dest.(type) {
		case *[]int16:
			*d = make([]int16, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []int32:
		p := src.([]int32)
		switch d := dest.(type) {
		case *[]int32:
			*d = make([]int32, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []*int32:
		p := src.([]*int32)
		switch d := dest.(type) {
		case *[]int32:
			*d = make([]int32, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []int64:
		p := src.([]int64)
		switch d := dest.(type) {
		case *[]int64:
			*d = make([]int64, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []*int64:
		p := src.([]*int64)
		switch d := dest.(type) {
		case *[]int64:
			*d = make([]int64, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []float32:
		p := src.([]float32)
		switch d := dest.(type) {
		case *[]float32:
			*d = make([]float32, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []*float32:
		p := src.([]*float32)
		switch d := dest.(type) {
		case *[]float32:
			*d = make([]float32, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []float64:
		p := src.([]float64)
		switch d := dest.(type) {
		case *[]float64:
			*d = make([]float64, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []*float64:
		p := src.([]*float64)
		switch d := dest.(type) {
		case *[]float64:
			*d = make([]float64, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []string:
		p := src.([]string)
		switch d := dest.(type) {
		case *[]string:
			*d = make([]string, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case []*string:
		p := src.([]*string)
		switch d := dest.(type) {
		case *[]string:
			*d = make([]string, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		case *[]any:
			*d = make([]any, len(p))
			for i := range p {
				pi := (p)[i]
				if err := convertAssignRows(&(*d)[i], pi); err != nil {
					return err
				}
			}
			return nil
		}
	case *string:
		p := src.(*string)
		return convertAssignRows(dest, *p)
	case *int8:
		p := src.(*int8)
		return convertAssignRows(dest, *p)
	case *int16:
		p := src.(*int16)
		return convertAssignRows(dest, *p)
	case *int32:
		p := src.(*int32)
		return convertAssignRows(dest, *p)
	case *int64:
		p := src.(*int64)
		return convertAssignRows(dest, *p)
	case *uint8:
		p := src.(*uint8)
		return convertAssignRows(dest, *p)
	case *uint16:
		p := src.(*uint16)
		return convertAssignRows(dest, *p)
	case *uint32:
		p := src.(*uint32)
		return convertAssignRows(dest, *p)
	case *uint64:
		p := src.(*uint64)
		return convertAssignRows(dest, *p)
	case *float32:
		p := src.(*float32)
		return convertAssignRows(dest, *p)
	case *float64:
		p := src.(*float64)
		return convertAssignRows(dest, *p)
	case *time.Time:
		p := src.(*time.Time)
		return convertAssignRows(dest, *p)
	}

	var sv reflect.Value

	switch d := dest.(type) {
	case *string:
		sv = reflect.ValueOf(src)
		switch sv.Kind() {
		case reflect.Bool,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			*d = asString(src)
			return nil
		}
	case *[]byte:
		sv = reflect.ValueOf(src)
		if b, ok := asBytes(nil, sv); ok {
			*d = b
			return nil
		}
	case *RawBytes:
		sv = reflect.ValueOf(src)
		if b, ok := asBytes([]byte(*d)[:0], sv); ok {
			*d = RawBytes(b)
			return nil
		}
	case *bool:
		bv, err := driver.Bool.ConvertValue(src)
		if err == nil {
			*d = bv.(bool)
		}
		return err
	case *any:
		*d = src
		return nil
	}

	if scanner, ok := dest.(sql.Scanner); ok {
		return scanner.Scan(src)
	}

	dpv := reflect.ValueOf(dest)
	if dpv.Kind() != reflect.Pointer {
		return errors.New("destination not a pointer")
	}
	if dpv.IsNil() {
		return errNilPtr
	}

	if !sv.IsValid() {
		sv = reflect.ValueOf(src)
	}

	dv := reflect.Indirect(dpv)
	if sv.IsValid() && sv.Type().AssignableTo(dv.Type()) {
		switch b := src.(type) {
		case []byte:
			dv.Set(reflect.ValueOf(bytes.Clone(b)))
		default:
			dv.Set(sv)
		}
		return nil
	}

	if dv.Kind() == sv.Kind() && sv.Type().ConvertibleTo(dv.Type()) {
		dv.Set(sv.Convert(dv.Type()))
		return nil
	}

	// The following conversions use a string value as an intermediate representation
	// to convert between various numeric types.
	//
	// This also allows scanning into user defined types such as "type Int int64".
	// For symmetry, also check for string destination types.
	switch dv.Kind() {
	case reflect.Pointer:
		if src == nil {
			dv.Set(reflect.Zero(dv.Type()))
			return nil
		}
		dv.Set(reflect.New(dv.Type().Elem()))
		return convertAssignRows(dv.Interface(), src)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if src == nil {
			return fmt.Errorf("converting NULL to %s is unsupported", dv.Kind())
		}
		s := asString(src)
		i64, err := strconv.ParseInt(s, 10, dv.Type().Bits())
		if err != nil {
			err = strconvErr(err)
			return fmt.Errorf("converting driver.Value type %T (%q) to a %s: %v", src, s, dv.Kind(), err)
		}
		dv.SetInt(i64)
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if src == nil {
			return fmt.Errorf("converting NULL to %s is unsupported", dv.Kind())
		}
		s := asString(src)
		u64, err := strconv.ParseUint(s, 10, dv.Type().Bits())
		if err != nil {
			err = strconvErr(err)
			return fmt.Errorf("converting driver.Value type %T (%q) to a %s: %v", src, s, dv.Kind(), err)
		}
		dv.SetUint(u64)
		return nil
	case reflect.Float32, reflect.Float64:
		if src == nil {
			return fmt.Errorf("converting NULL to %s is unsupported", dv.Kind())
		}
		s := asString(src)
		f64, err := strconv.ParseFloat(s, dv.Type().Bits())
		if err != nil {
			err = strconvErr(err)
			return fmt.Errorf("converting driver.Value type %T (%q) to a %s: %v", src, s, dv.Kind(), err)
		}
		dv.SetFloat(f64)
		return nil
	case reflect.String:
		if src == nil {
			return fmt.Errorf("converting NULL to %s is unsupported", dv.Kind())
		}
		switch v := src.(type) {
		case string:
			dv.SetString(v)
			return nil
		case []byte:
			dv.SetString(string(v))
			return nil
		}
	}

	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", src, dest)
}
