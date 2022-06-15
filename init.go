package nulltype

import (
	"reflect"

	jsoniter "github.com/json-iterator/go"
)

func EncoderRegister() {
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Bool{}).String(), &Bool{})
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Float64{}).String(), &Float64{})
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Int64{}).String(), &Int64{})
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(String{}).String(), &String{})
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(Time{}).String(), &Time{})
	jsoniter.RegisterTypeEncoder(reflect.TypeOf(UUID{}).String(), &UUID{})
}
