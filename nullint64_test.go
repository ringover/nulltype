package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithInt64Null(t *testing.T) {
	v := Int64{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		I Int64 `json:"int64"`
	}
	test.I = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithInt64NotNull(t *testing.T) {
	v := NewInt64(0)

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		I Int64 `json:"int64"`
	}
	test.I = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithInt64NotNull(t *testing.T) {
	var test struct {
		I Int64 `json:"int64"`
	}
	err := json.Unmarshal([]byte(`{"int64":0}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.I.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithInt64Null(t *testing.T) {
	var test struct {
		I Int64 `json:"int64"`
	}
	err := json.Unmarshal([]byte(`{"int64": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.I.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanInt64Null(t *testing.T) {
	v := Int64{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanInt64NotNull(t *testing.T) {
	v := Int64{}
	err := v.Scan(0)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueInt64Null(t *testing.T) {
	v := Int64{}
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
func testValueInt64NotNull(t *testing.T) {
	v := NewInt64(0)
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueInt64NotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Int64{}).String(), &Int64{})
	type tmp struct {
		Ptr      *Int64 `json:"ptr,omitempty"`
		Always   Int64  `json:"always"`
		OkInt64  Int64  `json:"ok_Int64,omitempty"`
		NokInt64 Int64  `json:"nok_Int64,omitempty"`
	}
	value := tmp{}
	value.Always = NewInt64(1)
	value.OkInt64 = NewInt64(2)
	value.NokInt64.Valid = false
	value.NokInt64.Int64 = 0
	t.Logf("%+v", value)
	jsoni, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	str := string(jsoni)
	if !strings.Contains(str, "nok_Int64") {
		t.Failed()
	}
	t.Log(str)
}

func TestNullInt64(t *testing.T) {
	t.Run("testMarshalJSONWithInt64Null", testMarshalJSONWithInt64Null)
	t.Run("testMarshalJSONWithInt64NotNull", testMarshalJSONWithInt64NotNull)
	t.Run("testUnmarshalJSONWithInt64Null", testUnmarshalJSONWithInt64Null)
	t.Run("testUnmarshalJSONWithInt64NotNull", testUnmarshalJSONWithInt64NotNull)
	t.Run("testScanInt64Null", testScanInt64Null)
	t.Run("testScanInt64NotNull", testScanInt64NotNull)
	t.Run("testValueInt64Null", testValueInt64Null)
	t.Run("testValueInt64NotNull", testValueInt64NotNull)
	t.Run("testValueInt64NotNull", testValueInt64NotNullInStruct)
}
