package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithArrayInt64Null(t *testing.T) {
	v := ArrayInt[int64]{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayInt[int64] `json:"array_int64"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithArrayInt64NotNull(t *testing.T) {
	v := NewArrayInt[int64]([]int64{123, -124, 125})

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayInt[int64] `json:"array_int64"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithArrayInt64NotNull(t *testing.T) {
	var test struct {
		F ArrayInt[int64] `json:"array_int64"`
	}
	err := json.Unmarshal([]byte(`{"array_int64": [123, 124, 125]}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithArrayInt64Null(t *testing.T) {
	var test struct {
		F ArrayInt[int64] `json:"array_int64"`
	}
	err := json.Unmarshal([]byte(`{"array_int64": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanArrayInt64Null(t *testing.T) {
	v := ArrayInt[int64]{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanArrayInt64NotNull(t *testing.T) {
	v := ArrayInt[int64]{}
	err := v.Scan([]int64{123, -124, 125})
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueArrayInt64Null(t *testing.T) {
	v := ArrayInt[int64]{}
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
func testValueArrayInt64NotNull(t *testing.T) {
	v := NewArrayInt[int64]([]int64{123, -124, 125})
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueArrayInt64NotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(ArrayInt[int64]{}).String(), &ArrayInt[int64]{})
	type tmp struct {
		Ptr         *ArrayInt[int64] `json:"ptr,omitempty"`
		Always      ArrayInt[int64]  `json:"always"`
		OkArrayInt  ArrayInt[int64]  `json:"ok_ArrayInt,omitempty"`
		NokArrayInt ArrayInt[int64]  `json:"nok_ArrayInt,omitempty"`
	}
	value := tmp{}
	value.Always = NewArrayInt[int64]([]int64{123, -124, 125})
	value.OkArrayInt = NewArrayInt[int64]([]int64{123, -124, 125})
	value.NokArrayInt.Valid = false
	value.NokArrayInt.Array = []int64{}
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

func TestNullArrayInt64(t *testing.T) {
	t.Run("testMarshalJSONWithArrayInt64Null", testMarshalJSONWithArrayInt64Null)
	t.Run("testMarshalJSONWithArrayInt64NotNull", testMarshalJSONWithArrayInt64NotNull)
	t.Run("testUnmarshalJSONWithArrayInt64Null", testUnmarshalJSONWithArrayInt64Null)
	t.Run("testUnmarshalJSONWithArrayInt64NotNull", testUnmarshalJSONWithArrayInt64NotNull)
	t.Run("testScanArrayInt64Null", testScanArrayInt64Null)
	t.Run("testScanArrayInt64NotNull", testScanArrayInt64NotNull)
	t.Run("testValueArrayInt64Null", testValueArrayInt64Null)
	t.Run("testValueArrayInt64NotNull", testValueArrayInt64NotNull)
	t.Run("testValueArrayInt64NotNullInStruct", testValueArrayInt64NotNullInStruct)
}
