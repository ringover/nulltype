# NullType

## Description

This package provides some types representing a type that may be null.
NullType implements the sql.Scanner, json.Unmarshaler and json.Marshaler interface for :
 * time.Time
 * string
 * int64
 * bool
 * float64
 * uuid.UUID

## NullType and omitempty
By default NullType struct are not empty.
To make `omitempty` tag functional on not valid NullType you have to call `EncoderRegister()` at the beginning of your program.


## Authors
 * Mickael ROUSSE
 * Vincent PEREZ
 * Corentin Gaspart

