package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithInt32Null(t *testing.T) {
	v := Int32{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		I Int32 `json:"int32"`
	}
	test.I = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithInt32NotNull(t *testing.T) {
	v := NewInt32(0)

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		I Int32 `json:"int32"`
	}
	test.I = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithInt32NotNull(t *testing.T) {
	var test struct {
		I Int32 `json:"int32"`
	}
	err := json.Unmarshal([]byte(`{"int32":0}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.I.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithInt32Null(t *testing.T) {
	var test struct {
		I Int32 `json:"int32"`
	}
	err := json.Unmarshal([]byte(`{"int32": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.I.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanInt32Null(t *testing.T) {
	v := Int32{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanInt32NotNull(t *testing.T) {
	v := Int32{}
	err := v.Scan(0)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueInt32Null(t *testing.T) {
	v := Int32{}
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
func testValueInt32NotNull(t *testing.T) {
	v := NewInt32(0)
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueInt32NotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Int32{}).String(), &Int32{})
	type tmp struct {
		Ptr      *Int32 `json:"ptr,omitempty"`
		Always   Int32  `json:"always"`
		OkInt32  Int32  `json:"ok_Int32,omitempty"`
		NokInt32 Int32  `json:"nok_Int32,omitempty"`
	}
	value := tmp{}
	value.Always = NewInt32(1)
	value.OkInt32 = NewInt32(2)
	value.NokInt32.Valid = false
	value.NokInt32.Int32 = 0
	t.Logf("%+v", value)
	jsoni, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	str := string(jsoni)
	if !strings.Contains(str, "nok_Int32") {
		t.Failed()
	}
	t.Log(str)
}

func TestNullInt32(t *testing.T) {
	t.Run("testMarshalJSONWithInt32Null", testMarshalJSONWithInt32Null)
	t.Run("testMarshalJSONWithInt32NotNull", testMarshalJSONWithInt32NotNull)
	t.Run("testUnmarshalJSONWithInt32Null", testUnmarshalJSONWithInt32Null)
	t.Run("testUnmarshalJSONWithInt32NotNull", testUnmarshalJSONWithInt32NotNull)
	t.Run("testScanInt32Null", testScanInt32Null)
	t.Run("testScanInt32NotNull", testScanInt32NotNull)
	t.Run("testValueInt32Null", testValueInt32Null)
	t.Run("testValueInt32NotNull", testValueInt32NotNull)
	t.Run("testValueInt32NotNull", testValueInt32NotNullInStruct)
}
