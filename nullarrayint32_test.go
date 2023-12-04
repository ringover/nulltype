package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithArrayInt32Null(t *testing.T) {
	v := ArrayInt[int32]{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayInt[int32] `json:"array_int32"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithArrayInt32NotNull(t *testing.T) {
	v := NewArrayInt[int32]([]int32{123, -124, 125})

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayInt[int32] `json:"array_int32"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithArrayInt32NotNull(t *testing.T) {
	var test struct {
		F ArrayInt[int32] `json:"array_int32"`
	}
	err := json.Unmarshal([]byte(`{"array_int32": [123, 124, 125]}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithArrayInt32Null(t *testing.T) {
	var test struct {
		F ArrayInt[int32] `json:"array_int32"`
	}
	err := json.Unmarshal([]byte(`{"array_int32": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanArrayInt32Null(t *testing.T) {
	v := ArrayInt[int32]{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanArrayInt32NotNull(t *testing.T) {
	v := ArrayInt[int32]{}
	err := v.Scan([]int32{123, -124, 125})
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueArrayInt32Null(t *testing.T) {
	v := ArrayInt[int32]{}
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
func testValueArrayInt32NotNull(t *testing.T) {
	v := NewArrayInt[int32]([]int32{123, -124, 125})
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueArrayInt32NotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(ArrayInt[int32]{}).String(), &ArrayInt[int32]{})
	type tmp struct {
		Ptr         *ArrayInt[int32] `json:"ptr,omitempty"`
		Always      ArrayInt[int32]  `json:"always"`
		OkArrayInt  ArrayInt[int32]  `json:"ok_ArrayInt,omitempty"`
		NokArrayInt ArrayInt[int32]  `json:"nok_ArrayInt,omitempty"`
	}
	value := tmp{}
	value.Always = NewArrayInt[int32]([]int32{123, -124, 125})
	value.OkArrayInt = NewArrayInt[int32]([]int32{123, -124, 125})
	value.NokArrayInt.Valid = false
	value.NokArrayInt.Array = []int32{}
	t.Logf("%+v", value)
	jsoni, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	str := string(jsoni)
	if !strings.Contains(str, "nok_ArrayInt") {
		t.Failed()
	}
	t.Log(str)
}

func TestNullArrayInt32(t *testing.T) {
	t.Run("testMarshalJSONWithArrayInt32Null", testMarshalJSONWithArrayInt32Null)
	t.Run("testMarshalJSONWithArrayInt32NotNull", testMarshalJSONWithArrayInt32NotNull)
	t.Run("testUnmarshalJSONWithArrayInt32Null", testUnmarshalJSONWithArrayInt32Null)
	t.Run("testUnmarshalJSONWithArrayInt32NotNull", testUnmarshalJSONWithArrayInt32NotNull)
	t.Run("testScanArrayInt32Null", testScanArrayInt32Null)
	t.Run("testScanArrayInt32NotNull", testScanArrayInt32NotNull)
	t.Run("testValueArrayInt32Null", testValueArrayInt32Null)
	t.Run("testValueArrayInt32NotNull", testValueArrayInt32NotNull)
	t.Run("testValueArrayInt32NotNullInStruct", testValueArrayInt32NotNullInStruct)
}
