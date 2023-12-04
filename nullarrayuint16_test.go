package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithArrayUint16Null(t *testing.T) {
	v := ArrayUint[uint16]{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayUint[uint16] `json:"array_uint16"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithArrayUint16NotNull(t *testing.T) {
	v := NewArrayUint[uint16]([]uint16{123, 124, 125})

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayUint[uint16] `json:"array_uint16"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithArrayUint16NotNull(t *testing.T) {
	var test struct {
		F ArrayUint[uint16] `json:"array_uint16"`
	}
	err := json.Unmarshal([]byte(`{"array_uint16": [123, 124, 125]}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithArrayUint16NotNullMustFailed(t *testing.T) {
	var test struct {
		F ArrayUint[uint16] `json:"array_uint16"`
	}
	err := json.Unmarshal([]byte(`{"array_uint16": [123, -124, 125]}`), &test)
	if err != nil {
		t.Logf(err.Error())
	}
	if test.F.Valid == false {
		t.Logf("Test success")
		return
	}
	t.Fatal("Failed")
}

func testUnmarshalJSONWithArrayUint16Null(t *testing.T) {
	var test struct {
		F ArrayUint[uint16] `json:"array_uint16"`
	}
	err := json.Unmarshal([]byte(`{"array_uint16": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanArrayUint16Null(t *testing.T) {
	v := ArrayUint[uint16]{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanArrayUint16NotNull(t *testing.T) {
	v := ArrayUint[uint16]{}
	err := v.Scan([]uint16{123, 124, 125})
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueArrayUint16Null(t *testing.T) {
	v := ArrayUint[uint16]{}
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
func testValueArrayUint16NotNull(t *testing.T) {
	v := NewArrayUint[uint16]([]uint16{123, 124, 125})
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueArrayUint16NotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(ArrayUint[uint16]{}).String(), &ArrayUint[uint16]{})
	type tmp struct {
		Ptr          *ArrayUint[uint16] `json:"ptr,omitempty"`
		Always       ArrayUint[uint16]  `json:"always"`
		OkArrayUint  ArrayUint[uint16]  `json:"ok_ArrayUint,omitempty"`
		NokArrayUint ArrayUint[uint16]  `json:"nok_ArrayUint,omitempty"`
	}
	value := tmp{}
	value.Always = NewArrayUint[uint16]([]uint16{123, 124, 125})
	value.OkArrayUint = NewArrayUint[uint16]([]uint16{123, 124, 125})
	value.NokArrayUint.Valid = false
	value.NokArrayUint.Array = []uint16{}
	t.Logf("%+v", value)
	jsoni, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	str := string(jsoni)
	if !strings.Contains(str, "nok_ArrayUint") {
		t.Failed()
	}
	t.Log(str)
}

func TestNullArrayUint16(t *testing.T) {
	t.Run("testMarshalJSONWithArrayUint16Null", testMarshalJSONWithArrayUint16Null)
	t.Run("testMarshalJSONWithArrayUint16NotNull", testMarshalJSONWithArrayUint16NotNull)
	t.Run("testUnmarshalJSONWithArrayUint16Null", testUnmarshalJSONWithArrayUint16Null)
	t.Run("testUnmarshalJSONWithArrayUint16NotNull", testUnmarshalJSONWithArrayUint16NotNull)
	t.Run("testScanArrayUint16Null", testScanArrayUint16Null)
	t.Run("testScanArrayUint16NotNull", testScanArrayUint16NotNull)
	t.Run("testValueArrayUint16Null", testValueArrayUint16Null)
	t.Run("testValueArrayUint16NotNull", testValueArrayUint16NotNull)
	t.Run("testValueArrayUint16NotNullInStruct", testValueArrayUint16NotNullInStruct)
	t.Run("testUnmarshalJSONWithArrayUint16NotNullMustFailed", testUnmarshalJSONWithArrayUint16NotNullMustFailed)
}
