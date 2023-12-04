package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithArrayUint64Null(t *testing.T) {
	v := ArrayUint[uint64]{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayUint[uint64] `json:"array_uint64"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithArrayUint64NotNull(t *testing.T) {
	v := NewArrayUint[uint64]([]uint64{123, 124, 125})

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayUint[uint64] `json:"array_uint64"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithArrayUint64NotNull(t *testing.T) {
	var test struct {
		F ArrayUint[uint64] `json:"array_uint64"`
	}
	err := json.Unmarshal([]byte(`{"array_uint64": [123, 124, 125]}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithArrayUint64NotNullMustFailed(t *testing.T) {
	var test struct {
		F ArrayUint[uint64] `json:"array_uint64"`
	}
	err := json.Unmarshal([]byte(`{"array_uint64": [123, -124, 125]}`), &test)
	if err != nil {
		t.Logf(err.Error())
	}
	if test.F.Valid == false {
		t.Logf("Test success")
		return
	}
	t.Fatal("Failed")
}

func testUnmarshalJSONWithArrayUint64Null(t *testing.T) {
	var test struct {
		F ArrayUint[uint64] `json:"array_uint64"`
	}
	err := json.Unmarshal([]byte(`{"array_uint64": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanArrayUint64Null(t *testing.T) {
	v := ArrayUint[uint64]{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanArrayUint64NotNull(t *testing.T) {
	v := ArrayUint[uint64]{}
	err := v.Scan([]uint64{123, 124, 125})
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueArrayUint64Null(t *testing.T) {
	v := ArrayUint[uint64]{}
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
func testValueArrayUint64NotNull(t *testing.T) {
	v := NewArrayUint[uint64]([]uint64{123, 124, 125})
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueArrayUint64NotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(ArrayUint[uint64]{}).String(), &ArrayUint[uint64]{})
	type tmp struct {
		Ptr          *ArrayUint[uint64] `json:"ptr,omitempty"`
		Always       ArrayUint[uint64]  `json:"always"`
		OkArrayUint  ArrayUint[uint64]  `json:"ok_ArrayUint,omitempty"`
		NokArrayUint ArrayUint[uint64]  `json:"nok_ArrayUint,omitempty"`
	}
	value := tmp{}
	value.Always = NewArrayUint[uint64]([]uint64{123, 124, 125})
	value.OkArrayUint = NewArrayUint[uint64]([]uint64{123, 124, 125})
	value.NokArrayUint.Valid = false
	value.NokArrayUint.Array = []uint64{}
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

func TestNullArrayUint64(t *testing.T) {
	t.Run("testMarshalJSONWithArrayUint64Null", testMarshalJSONWithArrayUint64Null)
	t.Run("testMarshalJSONWithArrayUint64NotNull", testMarshalJSONWithArrayUint64NotNull)
	t.Run("testUnmarshalJSONWithArrayUint64Null", testUnmarshalJSONWithArrayUint64Null)
	t.Run("testUnmarshalJSONWithArrayUint64NotNull", testUnmarshalJSONWithArrayUint64NotNull)
	t.Run("testScanArrayUint64Null", testScanArrayUint64Null)
	t.Run("testScanArrayUint64NotNull", testScanArrayUint64NotNull)
	t.Run("testValueArrayUint64Null", testValueArrayUint64Null)
	t.Run("testValueArrayUint64NotNull", testValueArrayUint64NotNull)
	t.Run("testValueArrayUint64NotNullInStruct", testValueArrayUint64NotNullInStruct)
	t.Run("testUnmarshalJSONWithArrayUint64NotNullMustFailed", testUnmarshalJSONWithArrayUint64NotNullMustFailed)
}
