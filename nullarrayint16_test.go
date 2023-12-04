package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithArrayInt16Null(t *testing.T) {
	v := ArrayInt[int16]{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayInt[int16] `json:"array_int16"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithArrayInt16NotNull(t *testing.T) {
	v := NewArrayInt[int16]([]int16{123, -124, 125})

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayInt[int16] `json:"array_int16"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithArrayInt16NotNull(t *testing.T) {
	var test struct {
		F ArrayInt[int16] `json:"array_int16"`
	}
	err := json.Unmarshal([]byte(`{"array_int16": [123, 124, 125]}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithArrayInt16Null(t *testing.T) {
	var test struct {
		F ArrayInt[int16] `json:"array_int16"`
	}
	err := json.Unmarshal([]byte(`{"array_int16": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanArrayInt16Null(t *testing.T) {
	v := ArrayInt[int16]{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanArrayPointerInt16NotNull(t *testing.T) {
	v := ArrayInt[int16]{}
	tmp := []int16{123, -124, 125}
	err := v.Scan([]*int16{&tmp[0], &tmp[1], &tmp[2]})
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	if reflect.DeepEqual(v.Array, []int16{0, 0, 0}) == true {
		t.Fatal("Array is not equal")
	}
	t.Logf("%+v", v)
}

func testScanArrayInt16NotNull(t *testing.T) {
	v := ArrayInt[int16]{}
	err := v.Scan([]int16{123, -124, 125})
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	if reflect.DeepEqual(v.Array, []int16{0, 0, 0}) == true {
		t.Fatal("Array is not equal")
	}
	t.Logf("%+v", v)
}

func testValueArrayInt16Null(t *testing.T) {
	v := ArrayInt[int16]{}
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
func testValueArrayInt16NotNull(t *testing.T) {
	v := NewArrayInt[int16]([]int16{123, -124, 125})
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueArrayInt16NotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(ArrayInt[int16]{}).String(), &ArrayInt[int16]{})
	type tmp struct {
		Ptr         *ArrayInt[int16] `json:"ptr,omitempty"`
		Always      ArrayInt[int16]  `json:"always"`
		OkArrayInt  ArrayInt[int16]  `json:"ok_ArrayInt,omitempty"`
		NokArrayInt ArrayInt[int16]  `json:"nok_ArrayInt,omitempty"`
	}
	value := tmp{}
	value.Always = NewArrayInt[int16]([]int16{123, -124, 125})
	value.OkArrayInt = NewArrayInt[int16]([]int16{123, -124, 125})
	value.NokArrayInt.Valid = false
	value.NokArrayInt.Array = []int16{}
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

func TestNullArrayInt16(t *testing.T) {
	t.Run("testMarshalJSONWithArrayInt16Null", testMarshalJSONWithArrayInt16Null)
	t.Run("testMarshalJSONWithArrayInt16NotNull", testMarshalJSONWithArrayInt16NotNull)
	t.Run("testUnmarshalJSONWithArrayInt16Null", testUnmarshalJSONWithArrayInt16Null)
	t.Run("testUnmarshalJSONWithArrayInt16NotNull", testUnmarshalJSONWithArrayInt16NotNull)
	t.Run("testScanArrayInt16Null", testScanArrayInt16Null)
	t.Run("testScanArrayInt16NotNull", testScanArrayInt16NotNull)
	t.Run("testValueArrayInt16Null", testValueArrayInt16Null)
	t.Run("testValueArrayInt16NotNull", testValueArrayInt16NotNull)
	t.Run("testValueArrayInt16NotNullInStruct", testValueArrayInt16NotNullInStruct)
	t.Run("testScanArrayPointerInt16NotNull", testScanArrayPointerInt16NotNull)
}
