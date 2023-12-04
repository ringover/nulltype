package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithArrayFloat32Null(t *testing.T) {
	v := ArrayFloat[float32]{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayFloat[float32] `json:"array_float32"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithArrayFloat32NotNull(t *testing.T) {
	v := NewArrayFloat[float32]([]float32{123.123, -124.124, 125.125})

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayFloat[float32] `json:"array_float32"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithArrayFloat32NotNull(t *testing.T) {
	var test struct {
		F ArrayFloat[float32] `json:"array_float32"`
	}
	err := json.Unmarshal([]byte(`{"array_float32": [123.123, -124.124, 125.125]}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithArrayFloat32Null(t *testing.T) {
	var test struct {
		F ArrayFloat[float32] `json:"array_float32"`
	}
	err := json.Unmarshal([]byte(`{"array_float32": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanArrayFloat32Null(t *testing.T) {
	v := ArrayFloat[float32]{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanArrayFloat32NotNull(t *testing.T) {
	v := ArrayFloat[float32]{}
	err := v.Scan([]float32{123.123, -124.124, 125.125})
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueArrayFloat32Null(t *testing.T) {
	v := ArrayFloat[float32]{}
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
func testValueArrayFloat32NotNull(t *testing.T) {
	v := NewArrayFloat[float32]([]float32{123.123, -124.124, 125.125})
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueArrayFloat32NotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(ArrayFloat[float32]{}).String(), &ArrayFloat[float32]{})
	type tmp struct {
		Ptr           *ArrayFloat[float32] `json:"ptr,omitempty"`
		Always        ArrayFloat[float32]  `json:"always"`
		OkArrayFloat  ArrayFloat[float32]  `json:"ok_ArrayFloat,omitempty"`
		NokArrayFloat ArrayFloat[float32]  `json:"nok_ArrayFloat,omitempty"`
	}
	value := tmp{}
	value.Always = NewArrayFloat[float32]([]float32{123.123, -124.124, 125.125})
	value.OkArrayFloat = NewArrayFloat[float32]([]float32{123.123, -124.124, 125.125})
	value.NokArrayFloat.Valid = false
	value.NokArrayFloat.Array = []float32{}
	t.Logf("%+v", value)
	jsoni, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	str := string(jsoni)
	if !strings.Contains(str, "nok_ArrayFloat") {
		t.Failed()
	}
	t.Log(str)
}

func TestNullArrayFloat32(t *testing.T) {
	t.Run("testMarshalJSONWithArrayFloat32Null", testMarshalJSONWithArrayFloat32Null)
	t.Run("testMarshalJSONWithArrayFloat32NotNull", testMarshalJSONWithArrayFloat32NotNull)
	t.Run("testUnmarshalJSONWithArrayFloat32Null", testUnmarshalJSONWithArrayFloat32Null)
	t.Run("testUnmarshalJSONWithArrayFloat32NotNull", testUnmarshalJSONWithArrayFloat32NotNull)
	t.Run("testScanArrayFloat32Null", testScanArrayFloat32Null)
	t.Run("testScanArrayFloat32NotNull", testScanArrayFloat32NotNull)
	t.Run("testValueArrayFloat32Null", testValueArrayFloat32Null)
	t.Run("testValueArrayFloat32NotNull", testValueArrayFloat32NotNull)
	t.Run("testValueArrayFloat32NotNullInStruct", testValueArrayFloat32NotNullInStruct)
}
