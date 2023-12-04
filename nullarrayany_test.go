package nulltype

/*
import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithFloat64Null(t *testing.T) {
	v := Float64{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F Float64 `json:"float64"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithFloat64NotNull(t *testing.T) {
	v := NewFloat64(3.1416)

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F Float64 `json:"float64"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithFloat64NotNull(t *testing.T) {
	var test struct {
		F Float64 `json:"float64"`
	}
	err := json.Unmarshal([]byte(`{"float64":3.1416}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithFloat64Null(t *testing.T) {
	var test struct {
		F Float64 `json:"float64"`
	}
	err := json.Unmarshal([]byte(`{"float64": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanFloat64Null(t *testing.T) {
	v := Float64{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanFloat64NotNull(t *testing.T) {
	v := Float64{}
	err := v.Scan(3.1416)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueFloat64Null(t *testing.T) {
	v := Float64{}
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
func testValueFloat64NotNull(t *testing.T) {
	v := NewFloat64(3.1416)
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueFloat64NotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Float64{}).String(), &Float64{})
	type tmp struct {
		Ptr        *Float64 `json:"ptr,omitempty"`
		Always     Float64  `json:"always"`
		OkFloat64  Float64  `json:"ok_Float64,omitempty"`
		NokFloat64 Float64  `json:"nok_Float64,omitempty"`
	}
	value := tmp{}
	value.Always = NewFloat64(1.1)
	value.OkFloat64 = NewFloat64(2.2)
	value.NokFloat64.Valid = false
	value.NokFloat64.Float64 = 0
	t.Logf("%+v", value)
	jsoni, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	str := string(jsoni)
	if !strings.Contains(str, "nok_Float64") {
		t.Failed()
	}
	t.Log(str)
}

func TestNullFloat64(t *testing.T) {
	t.Run("testMarshalJSONWithFloat64Null", testMarshalJSONWithFloat64Null)
	t.Run("testMarshalJSONWithFloat64NotNull", testMarshalJSONWithFloat64NotNull)
	t.Run("testUnmarshalJSONWithFloat64Null", testUnmarshalJSONWithFloat64Null)
	t.Run("testUnmarshalJSONWithFloat64NotNull", testUnmarshalJSONWithFloat64NotNull)
	t.Run("testScanFloat64Null", testScanFloat64Null)
	t.Run("testScanFloat64NotNull", testScanFloat64NotNull)
	t.Run("testValueFloat64Null", testValueFloat64Null)
	t.Run("testValueFloat64NotNull", testValueFloat64NotNull)
	t.Run("testValueFloat64NotNullInStruct", testValueFloat64NotNullInStruct)
}
*/
