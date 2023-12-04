package nulltype

import (
	"reflect"
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithArrayStringNull(t *testing.T) {
	v := ArrayString{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayString `json:"array_string"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testMarshalJSONWithArrayStringNotNull(t *testing.T) {
	v := NewArrayString([]string{"pomme", "banane", "poire"})

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		F ArrayString `json:"array_string"`
	}
	test.F = v
	jsoni, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsoni))
}

func testUnmarshalJSONWithArrayStringNotNull(t *testing.T) {
	var test struct {
		F ArrayString `json:"array_string"`
	}
	err := json.Unmarshal([]byte(`{"array_string": ["pomme", "banane", "poire"]}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithArrayStringNull(t *testing.T) {
	var test struct {
		F ArrayString `json:"array_string"`
	}
	err := json.Unmarshal([]byte(`{"array_string": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.F.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanArrayStringNull(t *testing.T) {
	v := ArrayString{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanArrayStringNotNull(t *testing.T) {
	v := ArrayString{}
	err := v.Scan([]string{"pomme", "banane", "poire"})
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueArrayStringNull(t *testing.T) {
	v := ArrayString{}
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
func testValueArrayStringNotNull(t *testing.T) {
	v := NewArrayString([]string{"pomme", "banane", "poire"})
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%d", value)
}

func testValueArrayStringNotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(ArrayString{}).String(), &ArrayString{})
	type tmp struct {
		Ptr            *ArrayString `json:"ptr,omitempty"`
		Always         ArrayString  `json:"always"`
		OkArrayString  ArrayString  `json:"ok_ArrayString,omitempty"`
		NokArrayString ArrayString  `json:"nok_ArrayString,omitempty"`
	}
	value := tmp{}
	value.Always = NewArrayString([]string{"pomme", "banane", "poire"})
	value.OkArrayString = NewArrayString([]string{"pomme", "banane", "poire"})
	value.NokArrayString.Valid = false
	value.NokArrayString.Array = []string{}
	t.Logf("%+v", value)
	jsoni, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	str := string(jsoni)
	if !strings.Contains(str, "nok_ArrayString") {
		t.Failed()
	}
	t.Log(str)
}

func TestNullArrayString(t *testing.T) {
	t.Run("testMarshalJSONWithArrayStringNull", testMarshalJSONWithArrayStringNull)
	t.Run("testMarshalJSONWithArrayStringNotNull", testMarshalJSONWithArrayStringNotNull)
	t.Run("testUnmarshalJSONWithArrayStringNull", testUnmarshalJSONWithArrayStringNull)
	t.Run("testUnmarshalJSONWithArrayStringNotNull", testUnmarshalJSONWithArrayStringNotNull)
	t.Run("testScanArrayStringNull", testScanArrayStringNull)
	t.Run("testScanArrayStringNotNull", testScanArrayStringNotNull)
	t.Run("testValueArrayStringNull", testValueArrayStringNull)
	t.Run("testValueArrayStringNotNull", testValueArrayStringNotNull)
	t.Run("testValueArrayStringNotNullInStruct", testValueArrayStringNotNullInStruct)
}
