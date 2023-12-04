package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithArrayFloat64Null(t *testing.T) {
	v := ArrayFloat[float64]{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayFloat[float64] `json:"array_float64"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithArrayFloat64NotNull(t *testing.T) {
	v := NewArrayFloat[float64]([]float64{123.123, -124.124, 125.125})

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayFloat[float64] `json:"array_float64"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithArrayFloat64NotNull(t *testing.T) {
	var test struct {
		F ArrayFloat[float64] `json:"array_float64"`
	}
	err := json.Unmarshal([]byte(`{"array_float64": [123.123, -124.124, 125.125]}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithArrayFloat64Null(t *testing.T) {
	var test struct {
		F ArrayFloat[float64] `json:"array_float64"`
	}
	err := json.Unmarshal([]byte(`{"array_float64": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanArrayFloat64Null(t *testing.T) {
	v := ArrayFloat[float64]{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanArrayFloat64NotNull(t *testing.T) {
	v := ArrayFloat[float64]{}
	err := v.Scan([]float64{123.123, -124.124, 125.125})
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	if reflect.DeepEqual(v.Array, []float64{0.0, 0.0, 0.0}) == true {
		t.Fatal("Array is not equal")
	}
	t.Logf("%+v", v)
}

func testScanArrayPointerFloat64NotNull(t *testing.T) {
	v := ArrayFloat[float64]{}
	tmp := []float64{123.123, -124.124, 125.125}
	err := v.Scan([]*float64{&tmp[0], &tmp[1], &tmp[2]})
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	if reflect.DeepEqual(v.Array, []float64{0.0, 0.0, 0.0}) == true {
		t.Fatal("Array is not equal")
	}
	t.Logf("%+v", v)
}

func testValueArrayFloat64Null(t *testing.T) {
	v := ArrayFloat[float64]{}
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
func testValueArrayFloat64NotNull(t *testing.T) {
	v := NewArrayFloat[float64]([]float64{123.123, -124.124, 125.125})
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueArrayFloat64NotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(ArrayFloat[float64]{}).String(), &ArrayFloat[float64]{})
	type tmp struct {
		Ptr           *ArrayFloat[float64] `json:"ptr,omitempty"`
		Always        ArrayFloat[float64]  `json:"always"`
		OkArrayFloat  ArrayFloat[float64]  `json:"ok_ArrayFloat,omitempty"`
		NokArrayFloat ArrayFloat[float64]  `json:"nok_ArrayFloat,omitempty"`
	}
	value := tmp{}
	value.Always = NewArrayFloat[float64]([]float64{123.123, -124.124, 125.125})
	value.OkArrayFloat = NewArrayFloat[float64]([]float64{123.123, -124.124, 125.125})
	value.NokArrayFloat.Valid = false
	value.NokArrayFloat.Array = []float64{}
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

func TestNullArrayFloat64(t *testing.T) {
	t.Run("testMarshalJSONWithArrayFloat64Null", testMarshalJSONWithArrayFloat64Null)
	t.Run("testMarshalJSONWithArrayFloat64NotNull", testMarshalJSONWithArrayFloat64NotNull)
	t.Run("testUnmarshalJSONWithArrayFloat64Null", testUnmarshalJSONWithArrayFloat64Null)
	t.Run("testUnmarshalJSONWithArrayFloat64NotNull", testUnmarshalJSONWithArrayFloat64NotNull)
	t.Run("testScanArrayFloat64Null", testScanArrayFloat64Null)
	t.Run("testScanArrayFloat64NotNull", testScanArrayFloat64NotNull)
	t.Run("testValueArrayFloat64Null", testValueArrayFloat64Null)
	t.Run("testValueArrayFloat64NotNull", testValueArrayFloat64NotNull)
	t.Run("testValueArrayFloat64NotNullInStruct", testValueArrayFloat64NotNullInStruct)
	t.Run("testScanArrayPointerFloat64NotNull", testScanArrayPointerFloat64NotNull)
}
