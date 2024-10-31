# NullType

## Description

This package provides some types representing a type that may be null.
NullType implements the sql.Scanner, json.Unmarshaler and json.Marshaler interface for :
 * time.Time
 * string
 * int64
 * int32
 * int16
 * int8
 * uint64
 * uint32
 * uint16
 * uint8
 * bool
 * float64
 * uuid.UUID
 * ArrayInt [int8, int16, int32, int64]
 * ArrayUint [uint8, uint16, uint32, uint64]
 * ArrayFloat [float32, float64]
 * ArrayString
 * ArrayAny

## NullType and omitempty
By default NullType struct are not empty.
To make `omitempty` tag functional on not valid NullType you have to call `EncoderRegister()` at the beginning of your program.


## Authors
 * Mickael ROUSSE
 * Vincent PEREZ
 * Corentin Gaspart

