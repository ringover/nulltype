package nulltype

import (
	"testing"
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

func TestNullFloat64(t *testing.T) {
	t.Run("testMarshalJSONWithFloat64Null", testMarshalJSONWithFloat64Null)
	t.Run("testMarshalJSONWithFloat64NotNull", testMarshalJSONWithFloat64NotNull)
	t.Run("testUnmarshalJSONWithFloat64Null", testUnmarshalJSONWithFloat64Null)
	t.Run("testUnmarshalJSONWithFloat64NotNull", testUnmarshalJSONWithFloat64NotNull)
	t.Run("testScanFloat64Null", testScanFloat64Null)
	t.Run("testScanFloat64NotNull", testScanFloat64NotNull)
	t.Run("testValueFloat64Null", testValueFloat64Null)
	t.Run("testValueFloat64NotNull", testValueFloat64NotNull)
}
