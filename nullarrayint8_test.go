package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithArrayInt8Null(t *testing.T) {
	v := ArrayInt[int8]{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayInt[int8] `json:"array_int8"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithArrayInt8NotNull(t *testing.T) {
	v := NewArrayInt[int8]([]int8{123, -124, 125})

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayInt[int8] `json:"array_int8"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithArrayInt8NotNull(t *testing.T) {
	var test struct {
		F ArrayInt[int8] `json:"array_int8"`
	}
	err := json.Unmarshal([]byte(`{"array_int8": [123, 124, 125]}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithArrayInt8Null(t *testing.T) {
	var test struct {
		F ArrayInt[int8] `json:"array_int8"`
	}
	err := json.Unmarshal([]byte(`{"array_int8": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanArrayInt8Null(t *testing.T) {
	v := ArrayInt[int8]{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanArrayInt8NotNull(t *testing.T) {
	v := ArrayInt[int8]{}
	err := v.Scan([]int8{123, -124, 125})
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueArrayInt8Null(t *testing.T) {
	v := ArrayInt[int8]{}
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
func testValueArrayInt8NotNull(t *testing.T) {
	v := NewArrayInt[int8]([]int8{123, -124, 125})
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueArrayInt8NotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(ArrayInt[int8]{}).String(), &ArrayInt[int8]{})
	type tmp struct {
		Ptr         *ArrayInt[int8] `json:"ptr,omitempty"`
		Always      ArrayInt[int8]  `json:"always"`
		OkArrayInt  ArrayInt[int8]  `json:"ok_ArrayInt,omitempty"`
		NokArrayInt ArrayInt[int8]  `json:"nok_ArrayInt,omitempty"`
	}
	value := tmp{}
	value.Always = NewArrayInt[int8]([]int8{123, -124, 125})
	value.OkArrayInt = NewArrayInt[int8]([]int8{123, -124, 125})
	value.NokArrayInt.Valid = false
	value.NokArrayInt.Array = []int8{}
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

func TestNullArrayInt(t *testing.T) {
	t.Run("testMarshalJSONWithArrayInt8Null", testMarshalJSONWithArrayInt8Null)
	t.Run("testMarshalJSONWithArrayInt8NotNull", testMarshalJSONWithArrayInt8NotNull)
	t.Run("testUnmarshalJSONWithArrayInt8Null", testUnmarshalJSONWithArrayInt8Null)
	t.Run("testUnmarshalJSONWithArrayInt8NotNull", testUnmarshalJSONWithArrayInt8NotNull)
	t.Run("testScanArrayInt8Null", testScanArrayInt8Null)
	t.Run("testScanArrayInt8NotNull", testScanArrayInt8NotNull)
	t.Run("testValueArrayInt8Null", testValueArrayInt8Null)
	t.Run("testValueArrayInt8NotNull", testValueArrayInt8NotNull)
	t.Run("testValueArrayInt8NotNullInStruct", testValueArrayInt8NotNullInStruct)
}
