package nulltype

import (
	"reflect"
	"strings"
	"testing"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
)

func testMarshalJSONWithUUIDNull(t *testing.T) {
	v := UUID{}
	v.Valid = false

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		UUID UUID `json:"uuid"`
	}
	test.UUID = v
	jsonstr, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsonstr))
}

func testMarshalJSONWithUUIDNotNull(t *testing.T) {
	v := NewUUID(uuid.New())

	bytes, err := v.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))

	var test struct {
		UUID UUID `json:"uuid"`
	}
	test.UUID = v
	jsonstr, err := json.Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(jsonstr))
}

func testUnmarshalJSONWithUUIDNotNull(t *testing.T) {
	var test struct {
		UUID UUID `json:"uuid"`
	}
	err := json.Unmarshal([]byte(`{"uuid":"a5f40f38-e5bf-486a-b921-113022284177"}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.UUID.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testUnmarshalJSONWithUUIDNull(t *testing.T) {
	var test struct {
		UUID UUID `json:"uuid"`
	}
	err := json.Unmarshal([]byte(`{"uuid": null}`), &test)
	if err != nil {
		t.Fatal(err)
	}
	if test.UUID.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", test)
}

func testScanUUIDNull(t *testing.T) {
	v := UUID{}
	err := v.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == true {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testScanUUIDNotNull(t *testing.T) {
	v := UUID{}
	err := v.Scan("a5f40f38-e5bf-486a-b921-113022284177")
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%+v", v)
}

func testValueUUIDNull(t *testing.T) {
	v := UUID{}
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
func testValueUUIDNotNull(t *testing.T) {
	v := NewUUID(uuid.New())
	value, err := v.Value()
	if err != nil {
		t.Fatal(err)
	}
	if v.Valid == false {
		t.Fatal(err)
	}
	t.Logf("%s", value)
}

func testValueUUIDNotNullInStruct(t *testing.T) {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(UUID{}).String(), &UUID{})
	type tmp struct {
		Ptr     *UUID `json:"ptr,omitempty"`
		Always  UUID  `json:"always"`
		OkUUID  UUID  `json:"ok_UUID,omitempty"`
		NokUUID UUID  `json:"nok_UUID,omitempty"`
	}
	value := tmp{}
	value.Always = NewUUID(uuid.New())
	value.OkUUID = NewUUID(uuid.New())
	value.NokUUID.Valid = false
	value.NokUUID.UUID = uuid.New()
	t.Logf("%+v", value)
	jsoni, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	str := string(jsoni)
	if !strings.Contains(str, "nok_UUID") {
		t.Failed()
	}
	t.Log(str)
}

func TestNullUUID(t *testing.T) {
	t.Run("testMarshalJSONWithUUIDNull", testMarshalJSONWithUUIDNull)
	t.Run("testMarshalJSONWithUUIDNotNull", testMarshalJSONWithUUIDNotNull)
	t.Run("testUnmarshalJSONWithUUIDNull", testUnmarshalJSONWithUUIDNull)
	t.Run("testUnmarshalJSONWithUUIDNotNull", testUnmarshalJSONWithUUIDNotNull)
	t.Run("testScanUUIDNull", testScanUUIDNull)
	t.Run("testScanUUIDNotNull", testScanUUIDNotNull)
	t.Run("testValueUUIDNull", testValueUUIDNull)
	t.Run("testValueUUIDNotNull", testValueUUIDNotNull)
	t.Run("testValueUUIDNotNull", testValueUUIDNotNullInStruct)
}
