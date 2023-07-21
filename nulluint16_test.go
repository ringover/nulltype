package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithUint16Null(t *testing.T) {
	v := Uint16{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		I Uint16 `json:"uint16"`
	}
	test.I = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithUint16NotNull(t *testing.T) {
	v := NewUint16(0)

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		I Uint16 `json:"uint16"`
	}
	test.I = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithUint16NotNull(t *testing.T) {
	var test struct {
		I Uint16 `json:"uint16"`
	}
	err := json.Unmarshal([]byte(`{"uint16":0}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.I.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithUint16Null(t *testing.T) {
	var test struct {
		I Uint16 `json:"uint16"`
	}
	err := json.Unmarshal([]byte(`{"uint16": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.I.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanUint16Null(t *testing.T) {
	v := Uint16{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanUint16NotNull(t *testing.T) {
	v := Uint16{}
	err := v.Scan(0)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueUint16Null(t *testing.T) {
	v := Uint16{}
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
func testValueUint16NotNull(t *testing.T) {
	v := NewUint16(0)
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueUint16NotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Uint16{}).String(), &Uint16{})
	type tmp struct {
		Ptr       *Uint16 `json:"ptr,omitempty"`
		Always    Uint16  `json:"always"`
		OkUint16  Uint16  `json:"ok_Uint16,omitempty"`
		NokUint16 Uint16  `json:"nok_Uint16,omitempty"`
	}
	value := tmp{}
	value.Always = NewUint16(1)
	value.OkUint16 = NewUint16(2)
	value.NokUint16.Valid = false
	value.NokUint16.Uint16 = 0
	t.Logf("%+v", value)
	jsoni, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	str := string(jsoni)
	if !strings.Contains(str, "nok_Uint16") {
		t.Failed()
	}
	t.Log(str)
}

func TestNullUint16(t *testing.T) {
	t.Run("testMarshalJSONWithUint16Null", testMarshalJSONWithUint16Null)
	t.Run("testMarshalJSONWithUint16NotNull", testMarshalJSONWithUint16NotNull)
	t.Run("testUnmarshalJSONWithUint16Null", testUnmarshalJSONWithUint16Null)
	t.Run("testUnmarshalJSONWithUint16NotNull", testUnmarshalJSONWithUint16NotNull)
	t.Run("testScanUint16Null", testScanUint16Null)
	t.Run("testScanUint16NotNull", testScanUint16NotNull)
	t.Run("testValueUint16Null", testValueUint16Null)
	t.Run("testValueUint16NotNull", testValueUint16NotNull)
	t.Run("testValueUint16NotNull", testValueUint16NotNullInStruct)
}
