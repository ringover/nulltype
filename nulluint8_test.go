package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithUint8Null(t *testing.T) {
	v := Uint8{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		I Uint8 `json:"uint8"`
	}
	test.I = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithUint8NotNull(t *testing.T) {
	v := NewUint8(0)

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		I Uint8 `json:"uint8"`
	}
	test.I = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithUint8NotNull(t *testing.T) {
	var test struct {
		I Uint8 `json:"uint8"`
	}
	err := json.Unmarshal([]byte(`{"uint8":0}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.I.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithUint8Null(t *testing.T) {
	var test struct {
		I Uint8 `json:"uint8"`
	}
	err := json.Unmarshal([]byte(`{"uint8": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.I.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanUint8Null(t *testing.T) {
	v := Uint8{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanUint8NotNull(t *testing.T) {
	v := Uint8{}
	err := v.Scan(0)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueUint8Null(t *testing.T) {
	v := Uint8{}
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
func testValueUint8NotNull(t *testing.T) {
	v := NewUint8(0)
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueUint8NotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Uint8{}).String(), &Uint8{})
	type tmp struct {
		Ptr      *Uint8 `json:"ptr,omitempty"`
		Always   Uint8  `json:"always"`
		OkUint8  Uint8  `json:"ok_Uint8,omitempty"`
		NokUint8 Uint8  `json:"nok_Uint8,omitempty"`
	}
	value := tmp{}
	value.Always = NewUint8(1)
	value.OkUint8 = NewUint8(2)
	value.NokUint8.Valid = false
	value.NokUint8.Uint8 = 0
	t.Logf("%+v", value)
	jsoni, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	str := string(jsoni)
	if !strings.Contains(str, "nok_Uint8") {
		t.Failed()
	}
	t.Log(str)
}

func TestNullUint8(t *testing.T) {
	t.Run("testMarshalJSONWithUint8Null", testMarshalJSONWithUint8Null)
	t.Run("testMarshalJSONWithUint8NotNull", testMarshalJSONWithUint8NotNull)
	t.Run("testUnmarshalJSONWithUint8Null", testUnmarshalJSONWithUint8Null)
	t.Run("testUnmarshalJSONWithUint8NotNull", testUnmarshalJSONWithUint8NotNull)
	t.Run("testScanUint8Null", testScanUint8Null)
	t.Run("testScanUint8NotNull", testScanUint8NotNull)
	t.Run("testValueUint8Null", testValueUint8Null)
	t.Run("testValueUint8NotNull", testValueUint8NotNull)
	t.Run("testValueUint8NotNull", testValueUint8NotNullInStruct)
}
