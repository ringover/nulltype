package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithUint32Null(t *testing.T) {
	v := Uint32{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		I Uint32 `json:"uint32"`
	}
	test.I = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithUint32NotNull(t *testing.T) {
	v := NewUint32(0)

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		I Uint32 `json:"uint32"`
	}
	test.I = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithUint32NotNull(t *testing.T) {
	var test struct {
		I Uint32 `json:"uint32"`
	}
	err := json.Unmarshal([]byte(`{"uint32":0}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.I.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithUint32Null(t *testing.T) {
	var test struct {
		I Uint32 `json:"uint32"`
	}
	err := json.Unmarshal([]byte(`{"uint32": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.I.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanUint32Null(t *testing.T) {
	v := Uint32{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanUint32NotNull(t *testing.T) {
	v := Uint32{}
	err := v.Scan(0)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueUint32Null(t *testing.T) {
	v := Uint32{}
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
func testValueUint32NotNull(t *testing.T) {
	v := NewUint32(0)
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueUint32NotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Uint32{}).String(), &Uint32{})
	type tmp struct {
		Ptr       *Uint32 `json:"ptr,omitempty"`
		Always    Uint32  `json:"always"`
		OkUint32  Uint32  `json:"ok_Uint32,omitempty"`
		NokUint32 Uint32  `json:"nok_Uint32,omitempty"`
	}
	value := tmp{}
	value.Always = NewUint32(1)
	value.OkUint32 = NewUint32(2)
	value.NokUint32.Valid = false
	value.NokUint32.Uint32 = 0
	t.Logf("%+v", value)
	jsoni, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	str := string(jsoni)
	if !strings.Contains(str, "nok_Uint32") {
		t.Failed()
	}
	t.Log(str)
}

func TestNullUint32(t *testing.T) {
	t.Run("testMarshalJSONWithUint32Null", testMarshalJSONWithUint32Null)
	t.Run("testMarshalJSONWithUint32NotNull", testMarshalJSONWithUint32NotNull)
	t.Run("testUnmarshalJSONWithUint32Null", testUnmarshalJSONWithUint32Null)
	t.Run("testUnmarshalJSONWithUint32NotNull", testUnmarshalJSONWithUint32NotNull)
	t.Run("testScanUint32Null", testScanUint32Null)
	t.Run("testScanUint32NotNull", testScanUint32NotNull)
	t.Run("testValueUint32Null", testValueUint32Null)
	t.Run("testValueUint32NotNull", testValueUint32NotNull)
	t.Run("testValueUint32NotNull", testValueUint32NotNullInStruct)
}
