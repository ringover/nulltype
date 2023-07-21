package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithInt8Null(t *testing.T) {
	v := Int8{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		I Int8 `json:"int8"`
	}
	test.I = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithInt8NotNull(t *testing.T) {
	v := NewInt8(0)

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		I Int8 `json:"int8"`
	}
	test.I = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithInt8NotNull(t *testing.T) {
	var test struct {
		I Int8 `json:"int8"`
	}
	err := json.Unmarshal([]byte(`{"int8":0}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.I.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithInt8Null(t *testing.T) {
	var test struct {
		I Int8 `json:"int8"`
	}
	err := json.Unmarshal([]byte(`{"int8": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.I.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanInt8Null(t *testing.T) {
	v := Int8{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanInt8NotNull(t *testing.T) {
	v := Int8{}
	err := v.Scan(0)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueInt8Null(t *testing.T) {
	v := Int8{}
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
func testValueInt8NotNull(t *testing.T) {
	v := NewInt8(0)
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueInt8NotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Int8{}).String(), &Int8{})
	type tmp struct {
		Ptr     *Int8 `json:"ptr,omitempty"`
		Always  Int8  `json:"always"`
		OkInt8  Int8  `json:"ok_Int8,omitempty"`
		NokInt8 Int8  `json:"nok_Int8,omitempty"`
	}
	value := tmp{}
	value.Always = NewInt8(1)
	value.OkInt8 = NewInt8(2)
	value.NokInt8.Valid = false
	value.NokInt8.Int8 = 0
	t.Logf("%+v", value)
	jsoni, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	str := string(jsoni)
	if !strings.Contains(str, "nok_Int8") {
		t.Failed()
	}
	t.Log(str)
}

func TestNullInt8(t *testing.T) {
	t.Run("testMarshalJSONWithInt8Null", testMarshalJSONWithInt8Null)
	t.Run("testMarshalJSONWithInt8NotNull", testMarshalJSONWithInt8NotNull)
	t.Run("testUnmarshalJSONWithInt8Null", testUnmarshalJSONWithInt8Null)
	t.Run("testUnmarshalJSONWithInt8NotNull", testUnmarshalJSONWithInt8NotNull)
	t.Run("testScanInt8Null", testScanInt8Null)
	t.Run("testScanInt8NotNull", testScanInt8NotNull)
	t.Run("testValueInt8Null", testValueInt8Null)
	t.Run("testValueInt8NotNull", testValueInt8NotNull)
	t.Run("testValueInt8NotNull", testValueInt8NotNullInStruct)
}
