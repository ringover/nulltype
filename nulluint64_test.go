package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithUint64Null(t *testing.T) {
	v := Uint64{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		I Uint64 `json:"uint64"`
	}
	test.I = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithUint64NotNull(t *testing.T) {
	v := NewUint64(0)

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		I Uint64 `json:"uint64"`
	}
	test.I = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithUint64NotNull(t *testing.T) {
	var test struct {
		I Uint64 `json:"uint64"`
	}
	err := json.Unmarshal([]byte(`{"uint64":0}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.I.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithUint64Null(t *testing.T) {
	var test struct {
		I Uint64 `json:"uint64"`
	}
	err := json.Unmarshal([]byte(`{"uint64": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.I.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanUint64Null(t *testing.T) {
	v := Uint64{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanUint64NotNull(t *testing.T) {
	v := Uint64{}
	err := v.Scan(0)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueUint64Null(t *testing.T) {
	v := Uint64{}
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
func testValueUint64NotNull(t *testing.T) {
	v := NewUint64(0)
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueUint64NotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Uint64{}).String(), &Uint64{})
	type tmp struct {
		Ptr      *Uint64 `json:"ptr,omitempty"`
		Always   Uint64  `json:"always"`
		OkUint64  Uint64  `json:"ok_Uint64,omitempty"`
		NokUint64 Uint64  `json:"nok_Uint64,omitempty"`
	}
	value := tmp{}
	value.Always = NewUint64(1)
	value.OkUint64 = NewUint64(2)
	value.NokUint64.Valid = false
	value.NokUint64.Uint64 = 0
	t.Logf("%+v", value)
	jsoni, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	str := string(jsoni)
	if !strings.Contains(str, "nok_Uint64") {
		t.Failed()
	}
	t.Log(str)
}

func TestNullUint64(t *testing.T) {
	t.Run("testMarshalJSONWithUint64Null", testMarshalJSONWithUint64Null)
	t.Run("testMarshalJSONWithUint64NotNull", testMarshalJSONWithUint64NotNull)
	t.Run("testUnmarshalJSONWithUint64Null", testUnmarshalJSONWithUint64Null)
	t.Run("testUnmarshalJSONWithUint64NotNull", testUnmarshalJSONWithUint64NotNull)
	t.Run("testScanUint64Null", testScanUint64Null)
	t.Run("testScanUint64NotNull", testScanUint64NotNull)
	t.Run("testValueUint64Null", testValueUint64Null)
	t.Run("testValueUint64NotNull", testValueUint64NotNull)
	t.Run("testValueUint64NotNull", testValueUint64NotNullInStruct)
}
