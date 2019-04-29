// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-04-28 21:37:43.700319 -0500 CDT m=+0.001899921
package slice

import "fmt"


// Bool
func BoolForEach(input []bool, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(bool):
		fn := function.(func(bool))
		for _, element := range input {
			fn(element)
		}

	case func(bool, int):
		fn := function.(func(bool, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(bool)", "func(bool, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

// Byte
func ByteForEach(input []byte, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(byte):
		fn := function.(func(byte))
		for _, element := range input {
			fn(element)
		}

	case func(byte, int):
		fn := function.(func(byte, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(byte)", "func(byte, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

// Complex128
func Complex128ForEach(input []complex128, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(complex128):
		fn := function.(func(complex128))
		for _, element := range input {
			fn(element)
		}

	case func(complex128, int):
		fn := function.(func(complex128, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(complex128)", "func(complex128, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

// Complex64
func Complex64ForEach(input []complex64, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(complex64):
		fn := function.(func(complex64))
		for _, element := range input {
			fn(element)
		}

	case func(complex64, int):
		fn := function.(func(complex64, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(complex64)", "func(complex64, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

// Error
func ErrorForEach(input []error, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(error):
		fn := function.(func(error))
		for _, element := range input {
			fn(element)
		}

	case func(error, int):
		fn := function.(func(error, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(error)", "func(error, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

// Float32
func Float32ForEach(input []float32, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(float32):
		fn := function.(func(float32))
		for _, element := range input {
			fn(element)
		}

	case func(float32, int):
		fn := function.(func(float32, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(float32)", "func(float32, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

// Float64
func Float64ForEach(input []float64, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(float64):
		fn := function.(func(float64))
		for _, element := range input {
			fn(element)
		}

	case func(float64, int):
		fn := function.(func(float64, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(float64)", "func(float64, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

// Int
func IntForEach(input []int, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(int):
		fn := function.(func(int))
		for _, element := range input {
			fn(element)
		}

	case func(int, int):
		fn := function.(func(int, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(int)", "func(int, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

// Int16
func Int16ForEach(input []int16, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(int16):
		fn := function.(func(int16))
		for _, element := range input {
			fn(element)
		}

	case func(int16, int):
		fn := function.(func(int16, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(int16)", "func(int16, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

// Int32
func Int32ForEach(input []int32, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(int32):
		fn := function.(func(int32))
		for _, element := range input {
			fn(element)
		}

	case func(int32, int):
		fn := function.(func(int32, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(int32)", "func(int32, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

// Int64
func Int64ForEach(input []int64, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(int64):
		fn := function.(func(int64))
		for _, element := range input {
			fn(element)
		}

	case func(int64, int):
		fn := function.(func(int64, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(int64)", "func(int64, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

// Int8
func Int8ForEach(input []int8, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(int8):
		fn := function.(func(int8))
		for _, element := range input {
			fn(element)
		}

	case func(int8, int):
		fn := function.(func(int8, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(int8)", "func(int8, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

// Rune
func RuneForEach(input []rune, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(rune):
		fn := function.(func(rune))
		for _, element := range input {
			fn(element)
		}

	case func(rune, int):
		fn := function.(func(rune, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(rune)", "func(rune, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

// String
func StringForEach(input []string, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(string):
		fn := function.(func(string))
		for _, element := range input {
			fn(element)
		}

	case func(string, int):
		fn := function.(func(string, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(string)", "func(string, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

// Uint
func UintForEach(input []uint, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(uint):
		fn := function.(func(uint))
		for _, element := range input {
			fn(element)
		}

	case func(uint, int):
		fn := function.(func(uint, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(uint)", "func(uint, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

// Uint16
func Uint16ForEach(input []uint16, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(uint16):
		fn := function.(func(uint16))
		for _, element := range input {
			fn(element)
		}

	case func(uint16, int):
		fn := function.(func(uint16, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(uint16)", "func(uint16, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

// Uint32
func Uint32ForEach(input []uint32, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(uint32):
		fn := function.(func(uint32))
		for _, element := range input {
			fn(element)
		}

	case func(uint32, int):
		fn := function.(func(uint32, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(uint32)", "func(uint32, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

// Uint64
func Uint64ForEach(input []uint64, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(uint64):
		fn := function.(func(uint64))
		for _, element := range input {
			fn(element)
		}

	case func(uint64, int):
		fn := function.(func(uint64, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(uint64)", "func(uint64, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

// Uint8
func Uint8ForEach(input []uint8, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(uint8):
		fn := function.(func(uint8))
		for _, element := range input {
			fn(element)
		}

	case func(uint8, int):
		fn := function.(func(uint8, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(uint8)", "func(uint8, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

// Uintptr
func UintptrForEach(input []uintptr, function interface{}) error {
	if !(isFunc(function)) {
		e := fmt.Sprintf("parameter (%s) is not a function", function)
		return newError(e)
	}

	switch function.(type) {
	case func(uintptr):
		fn := function.(func(uintptr))
		for _, element := range input {
			fn(element)
		}

	case func(uintptr, int):
		fn := function.(func(uintptr, int))
		for i, element := range input {
			fn(element, i)
		}

	default:
		e := fmt.Sprintf("function (%s) must have a firm: '%s' or '%s' and got: '%s'",
			getFuncName(function), "func(uintptr)", "func(uintptr, int)", getFuncFirm(function))

		return newError(e)
	}

	return nil
}

