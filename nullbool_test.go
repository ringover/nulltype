package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithBoolNull(t *testing.T) {
	v := Bool{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		B Bool `json:"bool"`
	}
	test.B = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithBoolNotNull(t *testing.T) {
	v := NewBool(true)

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		B Bool `json:"bool"`
	}
	test.B = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithBoolNotNull(t *testing.T) {
	var test struct {
		B Bool `json:"bool"`
	}
	err := json.Unmarshal([]byte(`{"bool":true}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.B.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithBoolNull(t *testing.T) {
	var test struct {
		B Bool `json:"bool"`
	}
	err := json.Unmarshal([]byte(`{"bool": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.B.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanBoolNull(t *testing.T) {
	v := Bool{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanBoolNotNull(t *testing.T) {
	v := Bool{}
	err := v.Scan(0)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueBoolNull(t *testing.T) {
	v := Bool{}
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
func testValueBoolNotNull(t *testing.T) {
	v := NewBool(true)
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueBoolNotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Bool{}).String(), &Bool{})
	type tmp struct {
		Ptr     *Bool `json:"ptr,omitempty"`
		Always  Bool  `json:"always"`
		OkBool  Bool  `json:"ok_bool,omitempty"`
		NokBool Bool  `json:"nok_bool,omitempty"`
	}
	value := tmp{}
	value.Always = NewBool(true)
	value.OkBool = NewBool(true)
	value.NokBool.Valid = false
	value.NokBool.Bool = true
	t.Logf("%+v", value)
	jsoni, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	str := string(jsoni)
	if !strings.Contains(str, "nok_bool") {
		t.Failed()
	}
	t.Log(str)
}

func TestNullBool(t *testing.T) {
	t.Run("testMarshalJSONWithBoolNull", testMarshalJSONWithBoolNull)
	t.Run("testMarshalJSONWithBoolNotNull", testMarshalJSONWithBoolNotNull)
	t.Run("testUnmarshalJSONWithBoolNull", testUnmarshalJSONWithBoolNull)
	t.Run("testUnmarshalJSONWithBoolNotNull", testUnmarshalJSONWithBoolNotNull)
	t.Run("testScanBoolNull", testScanBoolNull)
	t.Run("testScanBoolNotNull", testScanBoolNotNull)
	t.Run("testValueBoolNull", testValueBoolNull)
	t.Run("testValueBoolNotNull", testValueBoolNotNull)
	t.Run("testValueBoolNotNullInStruct", testValueBoolNotNullInStruct)
}
