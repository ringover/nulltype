package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithArrayUint32Null(t *testing.T) {
	v := ArrayUint[uint32]{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayUint[uint32] `json:"array_uint32"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithArrayUint32NotNull(t *testing.T) {
	v := NewArrayUint[uint32]([]uint32{123, 124, 125})

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayUint[uint32] `json:"array_uint32"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithArrayUint32NotNull(t *testing.T) {
	var test struct {
		F ArrayUint[uint32] `json:"array_uint32"`
	}
	err := json.Unmarshal([]byte(`{"array_uint32": [123, 124, 125]}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithArrayUint32NotNullMustFailed(t *testing.T) {
	var test struct {
		F ArrayUint[uint8] `json:"array_uint32"`
	}
	err := json.Unmarshal([]byte(`{"array_uint32": [123, -124, 125]}`), &test)
	if err != nil {
		t.Logf(err.Error())
	}
	if test.F.Valid == false {
		t.Logf("Test success")
		return
	}
	t.Fatal("Failed")
}

func testUnmarshalJSONWithArrayUint32Null(t *testing.T) {
	var test struct {
		F ArrayUint[uint32] `json:"array_uint32"`
	}
	err := json.Unmarshal([]byte(`{"array_uint32": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanArrayUint32Null(t *testing.T) {
	v := ArrayUint[uint32]{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanArrayUint32NotNull(t *testing.T) {
	v := ArrayUint[uint32]{}
	err := v.Scan([]uint32{123, 124, 125})
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueArrayUint32Null(t *testing.T) {
	v := ArrayUint[uint32]{}
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
func testValueArrayUint32NotNull(t *testing.T) {
	v := NewArrayUint[uint32]([]uint32{123, 124, 125})
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueArrayUint32NotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(ArrayUint[uint32]{}).String(), &ArrayUint[uint32]{})
	type tmp struct {
		Ptr          *ArrayUint[uint32] `json:"ptr,omitempty"`
		Always       ArrayUint[uint32]  `json:"always"`
		OkArrayUint  ArrayUint[uint32]  `json:"ok_ArrayUint,omitempty"`
		NokArrayUint ArrayUint[uint32]  `json:"nok_ArrayUint,omitempty"`
	}
	value := tmp{}
	value.Always = NewArrayUint[uint32]([]uint32{123, 124, 125})
	value.OkArrayUint = NewArrayUint[uint32]([]uint32{123, 124, 125})
	value.NokArrayUint.Valid = false
	value.NokArrayUint.Array = []uint32{}
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

func TestNullArrayUint32(t *testing.T) {
	t.Run("testMarshalJSONWithArrayUint32Null", testMarshalJSONWithArrayUint32Null)
	t.Run("testMarshalJSONWithArrayUint32NotNull", testMarshalJSONWithArrayUint32NotNull)
	t.Run("testUnmarshalJSONWithArrayUint32Null", testUnmarshalJSONWithArrayUint32Null)
	t.Run("testUnmarshalJSONWithArrayUint32NotNull", testUnmarshalJSONWithArrayUint32NotNull)
	t.Run("testScanArrayUint32Null", testScanArrayUint32Null)
	t.Run("testScanArrayUint32NotNull", testScanArrayUint32NotNull)
	t.Run("testValueArrayUint32Null", testValueArrayUint32Null)
	t.Run("testValueArrayUint32NotNull", testValueArrayUint32NotNull)
	t.Run("testValueArrayUint32NotNullInStruct", testValueArrayUint32NotNullInStruct)
	t.Run("testUnmarshalJSONWithArrayUint32NotNullMustFailed", testUnmarshalJSONWithArrayUint32NotNullMustFailed)
}
