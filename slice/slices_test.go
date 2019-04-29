// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-04-28 21:37:43.703222 -0500 CDT m=+0.004802983
package slice

import (
	"reflect"
	"testing"
)

func TestBoolForEach(t *testing.T) {
	inputInterface := getInput("bool")
	input := inputInterface.([]bool)

	expectInterface := getExpect("bool")
	expect := expectInterface.([]bool)

	current := make([]bool, len(expect))

	if err := BoolForEach(input, func(e bool, i int) {
		current[i] = boolFunction(e)
	}); err != nil {
		t.Errorf("running BoolForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("BoolForEach returns %v; expected %v", current, expect)
	}
}

func TestByteForEach(t *testing.T) {
	inputInterface := getInput("byte")
	input := inputInterface.([]byte)

	expectInterface := getExpect("byte")
	expect := expectInterface.([]byte)

	current := make([]byte, len(expect))

	if err := ByteForEach(input, func(e byte, i int) {
		current[i] = byteFunction(e)
	}); err != nil {
		t.Errorf("running ByteForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("ByteForEach returns %v; expected %v", current, expect)
	}
}

func TestComplex128ForEach(t *testing.T) {
	inputInterface := getInput("complex128")
	input := inputInterface.([]complex128)

	expectInterface := getExpect("complex128")
	expect := expectInterface.([]complex128)

	current := make([]complex128, len(expect))

	if err := Complex128ForEach(input, func(e complex128, i int) {
		current[i] = complex128Function(e)
	}); err != nil {
		t.Errorf("running Complex128ForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("Complex128ForEach returns %v; expected %v", current, expect)
	}
}

func TestComplex64ForEach(t *testing.T) {
	inputInterface := getInput("complex64")
	input := inputInterface.([]complex64)

	expectInterface := getExpect("complex64")
	expect := expectInterface.([]complex64)

	current := make([]complex64, len(expect))

	if err := Complex64ForEach(input, func(e complex64, i int) {
		current[i] = complex64Function(e)
	}); err != nil {
		t.Errorf("running Complex64ForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("Complex64ForEach returns %v; expected %v", current, expect)
	}
}

func TestErrorForEach(t *testing.T) {
	inputInterface := getInput("error")
	input := inputInterface.([]error)

	expectInterface := getExpect("error")
	expect := expectInterface.([]error)

	current := make([]error, len(expect))

	if err := ErrorForEach(input, func(e error, i int) {
		current[i] = errorFunction(e)
	}); err != nil {
		t.Errorf("running ErrorForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("ErrorForEach returns %v; expected %v", current, expect)
	}
}

func TestFloat32ForEach(t *testing.T) {
	inputInterface := getInput("float32")
	input := inputInterface.([]float32)

	expectInterface := getExpect("float32")
	expect := expectInterface.([]float32)

	current := make([]float32, len(expect))

	if err := Float32ForEach(input, func(e float32, i int) {
		current[i] = float32Function(e)
	}); err != nil {
		t.Errorf("running Float32ForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("Float32ForEach returns %v; expected %v", current, expect)
	}
}

func TestFloat64ForEach(t *testing.T) {
	inputInterface := getInput("float64")
	input := inputInterface.([]float64)

	expectInterface := getExpect("float64")
	expect := expectInterface.([]float64)

	current := make([]float64, len(expect))

	if err := Float64ForEach(input, func(e float64, i int) {
		current[i] = float64Function(e)
	}); err != nil {
		t.Errorf("running Float64ForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("Float64ForEach returns %v; expected %v", current, expect)
	}
}

func TestIntForEach(t *testing.T) {
	inputInterface := getInput("int")
	input := inputInterface.([]int)

	expectInterface := getExpect("int")
	expect := expectInterface.([]int)

	current := make([]int, len(expect))

	if err := IntForEach(input, func(e int, i int) {
		current[i] = intFunction(e)
	}); err != nil {
		t.Errorf("running IntForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("IntForEach returns %v; expected %v", current, expect)
	}
}

func TestInt16ForEach(t *testing.T) {
	inputInterface := getInput("int16")
	input := inputInterface.([]int16)

	expectInterface := getExpect("int16")
	expect := expectInterface.([]int16)

	current := make([]int16, len(expect))

	if err := Int16ForEach(input, func(e int16, i int) {
		current[i] = int16Function(e)
	}); err != nil {
		t.Errorf("running Int16ForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("Int16ForEach returns %v; expected %v", current, expect)
	}
}

func TestInt32ForEach(t *testing.T) {
	inputInterface := getInput("int32")
	input := inputInterface.([]int32)

	expectInterface := getExpect("int32")
	expect := expectInterface.([]int32)

	current := make([]int32, len(expect))

	if err := Int32ForEach(input, func(e int32, i int) {
		current[i] = int32Function(e)
	}); err != nil {
		t.Errorf("running Int32ForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("Int32ForEach returns %v; expected %v", current, expect)
	}
}

func TestInt64ForEach(t *testing.T) {
	inputInterface := getInput("int64")
	input := inputInterface.([]int64)

	expectInterface := getExpect("int64")
	expect := expectInterface.([]int64)

	current := make([]int64, len(expect))

	if err := Int64ForEach(input, func(e int64, i int) {
		current[i] = int64Function(e)
	}); err != nil {
		t.Errorf("running Int64ForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("Int64ForEach returns %v; expected %v", current, expect)
	}
}

func TestInt8ForEach(t *testing.T) {
	inputInterface := getInput("int8")
	input := inputInterface.([]int8)

	expectInterface := getExpect("int8")
	expect := expectInterface.([]int8)

	current := make([]int8, len(expect))

	if err := Int8ForEach(input, func(e int8, i int) {
		current[i] = int8Function(e)
	}); err != nil {
		t.Errorf("running Int8ForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("Int8ForEach returns %v; expected %v", current, expect)
	}
}

func TestRuneForEach(t *testing.T) {
	inputInterface := getInput("rune")
	input := inputInterface.([]rune)

	expectInterface := getExpect("rune")
	expect := expectInterface.([]rune)

	current := make([]rune, len(expect))

	if err := RuneForEach(input, func(e rune, i int) {
		current[i] = runeFunction(e)
	}); err != nil {
		t.Errorf("running RuneForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("RuneForEach returns %v; expected %v", current, expect)
	}
}

func TestStringForEach(t *testing.T) {
	inputInterface := getInput("string")
	input := inputInterface.([]string)

	expectInterface := getExpect("string")
	expect := expectInterface.([]string)

	current := make([]string, len(expect))

	if err := StringForEach(input, func(e string, i int) {
		current[i] = stringFunction(e)
	}); err != nil {
		t.Errorf("running StringForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("StringForEach returns %v; expected %v", current, expect)
	}
}

func TestUintForEach(t *testing.T) {
	inputInterface := getInput("uint")
	input := inputInterface.([]uint)

	expectInterface := getExpect("uint")
	expect := expectInterface.([]uint)

	current := make([]uint, len(expect))

	if err := UintForEach(input, func(e uint, i int) {
		current[i] = uintFunction(e)
	}); err != nil {
		t.Errorf("running UintForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("UintForEach returns %v; expected %v", current, expect)
	}
}

func TestUint16ForEach(t *testing.T) {
	inputInterface := getInput("uint16")
	input := inputInterface.([]uint16)

	expectInterface := getExpect("uint16")
	expect := expectInterface.([]uint16)

	current := make([]uint16, len(expect))

	if err := Uint16ForEach(input, func(e uint16, i int) {
		current[i] = uint16Function(e)
	}); err != nil {
		t.Errorf("running Uint16ForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("Uint16ForEach returns %v; expected %v", current, expect)
	}
}

func TestUint32ForEach(t *testing.T) {
	inputInterface := getInput("uint32")
	input := inputInterface.([]uint32)

	expectInterface := getExpect("uint32")
	expect := expectInterface.([]uint32)

	current := make([]uint32, len(expect))

	if err := Uint32ForEach(input, func(e uint32, i int) {
		current[i] = uint32Function(e)
	}); err != nil {
		t.Errorf("running Uint32ForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("Uint32ForEach returns %v; expected %v", current, expect)
	}
}

func TestUint64ForEach(t *testing.T) {
	inputInterface := getInput("uint64")
	input := inputInterface.([]uint64)

	expectInterface := getExpect("uint64")
	expect := expectInterface.([]uint64)

	current := make([]uint64, len(expect))

	if err := Uint64ForEach(input, func(e uint64, i int) {
		current[i] = uint64Function(e)
	}); err != nil {
		t.Errorf("running Uint64ForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("Uint64ForEach returns %v; expected %v", current, expect)
	}
}

func TestUint8ForEach(t *testing.T) {
	inputInterface := getInput("uint8")
	input := inputInterface.([]uint8)

	expectInterface := getExpect("uint8")
	expect := expectInterface.([]uint8)

	current := make([]uint8, len(expect))

	if err := Uint8ForEach(input, func(e uint8, i int) {
		current[i] = uint8Function(e)
	}); err != nil {
		t.Errorf("running Uint8ForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("Uint8ForEach returns %v; expected %v", current, expect)
	}
}

func TestUintptrForEach(t *testing.T) {
	inputInterface := getInput("uintptr")
	input := inputInterface.([]uintptr)

	expectInterface := getExpect("uintptr")
	expect := expectInterface.([]uintptr)

	current := make([]uintptr, len(expect))

	if err := UintptrForEach(input, func(e uintptr, i int) {
		current[i] = uintptrFunction(e)
	}); err != nil {
		t.Errorf("running UintptrForEach got the error: %s",  err.Error())
	}

	if got := reflect.DeepEqual(current, expect); !got {
		t.Errorf("UintptrForEach returns %v; expected %v", current, expect)
	}
}
