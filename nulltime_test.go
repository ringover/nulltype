package nulltype

import (
	"encoding/json"
	"testing"
	"time"
)

func testMarshalJSONWithTimeNull(t *testing.T) {
	v := Time{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		Time Time `json:"time"`
	}
	test.Time = v
	jsonstr, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsonstr))
}

func testMarshalJSONWithTimeNotNull(t *testing.T) {
	v := NewTime(time.Now())

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		Time Time `json:"time"`
	}
	test.Time = v
	jsonstr, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsonstr))
}

func testUnmarshalJSONWithTimeNotNull(t *testing.T) {
	var test struct {
		Time Time `json:"time"`
	}
	err := json.Unmarshal([]byte(`{"time":"2018-06-05T22:26:45.782195524Z"}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.Time.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithTimeNull(t *testing.T) {
	var test struct {
		Time Time `json:"time"`
	}
	err := json.Unmarshal([]byte(`{"time": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.Time.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanTimeNull(t *testing.T) {
	v := Time{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanTimeNotNull(t *testing.T) {
	v := Time{}
	err := v.Scan("2018-06-05 22:26:45")
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueTimeNull(t *testing.T) {
	v := Time{}
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
func testValueTimeNotNull(t *testing.T) {
	v := NewTime(time.Now())
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%s", value)
}

func TestNullTime(t *testing.T) {
	t.Run("testMarshalJSONWithTimeNull", testMarshalJSONWithTimeNull)
	t.Run("testMarshalJSONWithTimeNotNull", testMarshalJSONWithTimeNotNull)
	t.Run("testUnmarshalJSONWithTimeNull", testUnmarshalJSONWithTimeNull)
	t.Run("testUnmarshalJSONWithTimeNotNull", testUnmarshalJSONWithTimeNotNull)
	t.Run("testScanTimeNull", testScanTimeNull)
	t.Run("testScanTimeNotNull", testScanTimeNotNull)
	t.Run("testValueTimeNull", testValueTimeNull)
	t.Run("testValueTimeNotNull", testValueTimeNotNull)
}
