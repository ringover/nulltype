package nulltype

import (
	"reflect"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Bool{}).String(), &Bool{})
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Float64{}).String(), &Float64{})
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Int64{}).String(), &Int64{})
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Int32{}).String(), &Int32{})
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Int16{}).String(), &Int16{})
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Int8{}).String(), &Int8{})
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Uint64{}).String(), &Uint64{})
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Uint32{}).String(), &Uint32{})
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Uint16{}).String(), &Uint16{})
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Uint8{}).String(), &Uint8{})
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(String{}).String(), &String{})
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Time{}).String(), &Time{})
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(UUID{}).String(), &UUID{})
}
