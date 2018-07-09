package nulltype

import (
	"encoding/json"
	"testing"
)

func testMarshalJSONWithStringNull(t *testing.T) {
	v := String{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		Str String `json:"str"`
	}
	test.Str = v
	jsonstr, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsonstr))
}

func testMarshalJSONWithStringNotNull(t *testing.T) {
	v := NewString("testMarshalJSONWithStringNotNull")

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		Str String `json:"str"`
	}
	test.Str = v
	jsonstr, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsonstr))
}

func testUnmarshalJSONWithStringNotNull(t *testing.T) {
	var test struct {
		Str String `json:"str"`
	}
	err := json.Unmarshal([]byte(`{"str":"testUnmarshalJSONWithStringNotNull"}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithStringNull(t *testing.T) {
	var test struct {
		Str String `json:"str"`
	}
	err := json.Unmarshal([]byte(`{"str": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanStringNull(t *testing.T) {
	v := String{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanStringNotNull(t *testing.T) {
	v := String{}
	err := v.Scan("testScanStringNotNull")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueStringNull(t *testing.T) {
	v := String{}
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
func testValueStringNotNull(t *testing.T) {
	v := NewString("testValueStringNotNull")
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", value)
}

func TestNullString(t *testing.T) {
	t.Run("testMarshalJSONWithStringNull", testMarshalJSONWithStringNull)
	t.Run("testMarshalJSONWithStringNotNull", testMarshalJSONWithStringNotNull)
	t.Run("testUnmarshalJSONWithStringNull", testUnmarshalJSONWithStringNull)
	t.Run("testUnmarshalJSONWithStringNotNull", testUnmarshalJSONWithStringNotNull)
	t.Run("testScanStringNull", testScanStringNull)
	t.Run("testScanStringNotNull", testScanStringNotNull)
	t.Run("testValueStringNull", testValueStringNull)
	t.Run("testValueStringNotNull", testValueStringNotNull)
}
