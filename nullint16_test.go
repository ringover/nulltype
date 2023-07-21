package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithInt16Null(t *testing.T) {
	v := Int16{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		I Int16 `json:"int16"`
	}
	test.I = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithInt16NotNull(t *testing.T) {
	v := NewInt16(0)

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		I Int16 `json:"int16"`
	}
	test.I = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithInt16NotNull(t *testing.T) {
	var test struct {
		I Int16 `json:"int16"`
	}
	err := json.Unmarshal([]byte(`{"int16":0}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.I.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithInt16Null(t *testing.T) {
	var test struct {
		I Int16 `json:"int16"`
	}
	err := json.Unmarshal([]byte(`{"int16": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.I.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanInt16Null(t *testing.T) {
	v := Int16{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanInt16NotNull(t *testing.T) {
	v := Int16{}
	err := v.Scan(0)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueInt16Null(t *testing.T) {
	v := Int16{}
	v.Valid = false

	err, value := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if value == nil {
		t.Log(value)
	} else {
		t.Fatal(value)
	}
}
func testValueInt16NotNull(t *testing.T) {
	v := NewInt16(0)
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueInt16NotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Int16{}).String(), &Int16{})
	type tmp struct {
		Ptr      *Int16 `json:"ptr,omitempty"`
		Always   Int16  `json:"always"`
		OkInt16  Int16  `json:"ok_Int16,omitempty"`
		NokInt16 Int16  `json:"nok_Int16,omitempty"`
	}
	value := tmp{}
	value.Always = NewInt16(1)
	value.OkInt16 = NewInt16(2)
	value.NokInt16.Valid = false
	value.NokInt16.Int16 = 0
	t.Logf("%+v", value)
	jsoni, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	str := string(jsoni)
	if !strings.Contains(str, "nok_Int16") {
		t.Failed()
	}
	t.Log(str)
}

func TestNullInt16(t *testing.T) {
	t.Run("testMarshalJSONWithInt16Null", testMarshalJSONWithInt16Null)
	t.Run("testMarshalJSONWithInt16NotNull", testMarshalJSONWithInt16NotNull)
	t.Run("testUnmarshalJSONWithInt16Null", testUnmarshalJSONWithInt16Null)
	t.Run("testUnmarshalJSONWithInt16NotNull", testUnmarshalJSONWithInt16NotNull)
	t.Run("testScanInt16Null", testScanInt16Null)
	t.Run("testScanInt16NotNull", testScanInt16NotNull)
	t.Run("testValueInt16Null", testValueInt16Null)
	t.Run("testValueInt16NotNull", testValueInt16NotNull)
	t.Run("testValueInt16NotNull", testValueInt16NotNullInStruct)
}
