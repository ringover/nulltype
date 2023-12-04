package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithArrayUint8Null(t *testing.T) {
	v := ArrayUint[uint8]{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayUint[uint8] `json:"array_uint8"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithArrayUint8NotNull(t *testing.T) {
	v := NewArrayUint[uint8]([]uint8{123, 124, 125})

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayUint[uint8] `json:"array_uint8"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithArrayUint8NotNull(t *testing.T) {
	var test struct {
		F ArrayUint[uint8] `json:"array_uint8"`
	}
	err := json.Unmarshal([]byte(`{"array_uint8": [123, 124, 125]}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithArrayUint8NotNullMustFailed(t *testing.T) {
	var test struct {
		F ArrayUint[uint8] `json:"array_uint8"`
	}
	err := json.Unmarshal([]byte(`{"array_uint8": [123, -124, 125]}`), &test)
	if err != nil {
		t.Logf(err.Error())
	}
	if test.F.Valid == false {
		t.Logf("Test success")
		return
	}
	t.Fatal("Failed")
}

func testUnmarshalJSONWithArrayUint8Null(t *testing.T) {
	var test struct {
		F ArrayUint[uint8] `json:"array_uint8"`
	}
	err := json.Unmarshal([]byte(`{"array_uint8": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanArrayUint8Null(t *testing.T) {
	v := ArrayUint[uint8]{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanArrayUint8NotNull(t *testing.T) {
	v := ArrayUint[uint8]{}
	err := v.Scan([]uint8{123, 124, 125})
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueArrayUint8Null(t *testing.T) {
	v := ArrayUint[uint8]{}
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
func testValueArrayUint8NotNull(t *testing.T) {
	v := NewArrayUint[uint8]([]uint8{123, 124, 125})
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueArrayUint8NotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(ArrayUint[uint8]{}).String(), &ArrayUint[uint8]{})
	type tmp struct {
		Ptr          *ArrayUint[uint8] `json:"ptr,omitempty"`
		Always       ArrayUint[uint8]  `json:"always"`
		OkArrayUint  ArrayUint[uint8]  `json:"ok_ArrayUint,omitempty"`
		NokArrayUint ArrayUint[uint8]  `json:"nok_ArrayUint,omitempty"`
	}
	value := tmp{}
	value.Always = NewArrayUint[uint8]([]uint8{123, 124, 125})
	value.OkArrayUint = NewArrayUint[uint8]([]uint8{123, 124, 125})
	value.NokArrayUint.Valid = false
	value.NokArrayUint.Array = []uint8{}
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

func TestNullArrayUint(t *testing.T) {
	t.Run("testMarshalJSONWithArrayUint8Null", testMarshalJSONWithArrayUint8Null)
	t.Run("testMarshalJSONWithArrayUint8NotNull", testMarshalJSONWithArrayUint8NotNull)
	t.Run("testUnmarshalJSONWithArrayUint8Null", testUnmarshalJSONWithArrayUint8Null)
	t.Run("testUnmarshalJSONWithArrayUint8NotNull", testUnmarshalJSONWithArrayUint8NotNull)
	t.Run("testScanArrayUint8Null", testScanArrayUint8Null)
	t.Run("testScanArrayUint8NotNull", testScanArrayUint8NotNull)
	t.Run("testValueArrayUint8Null", testValueArrayUint8Null)
	t.Run("testValueArrayUint8NotNull", testValueArrayUint8NotNull)
	t.Run("testValueArrayUint8NotNullInStruct", testValueArrayUint8NotNullInStruct)
	t.Run("testUnmarshalJSONWithArrayUint8NotNullMustFailed", testUnmarshalJSONWithArrayUint8NotNullMustFailed)
}
