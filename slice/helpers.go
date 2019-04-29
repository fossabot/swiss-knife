package slice

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

//go:generate go run gen.go

func getFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func getFuncFirm(fn interface{}) string {
	typ := reflect.TypeOf(fn)

	in := typ.NumIn()
	out := typ.NumOut()

	var inParam, outParam string

	for i := 0; i < in; i++ {
		input := typ.In(i).String()

		if strings.Contains(input, "func") {
			return input
		}

		if i == 0 {
			inParam = input
			continue
		}

		inParam += ", " + input
	}

	for i := 0; i < out; i++ {
		input := typ.Out(i).String()

		if strings.Contains(input, "func") {
			return input
		}

		if i == 0 {
			outParam = input
			continue
		}

		outParam += ", " + input

		if i == out-1 {
			outParam += ")"
		}
	}

	if strings.Contains(outParam, ",") {
		outParam = fmt.Sprintf("(%s)", outParam)
	}

	if out > 0 {
		outParam = " " + outParam
	}

	return fmt.Sprintf("func(%s)%s", inParam, outParam)
}

func isFunc(fn interface{}) bool {
	return reflect.TypeOf(fn).Kind() == reflect.Func
}

func newError(e string) error {
	return errors.New(e)
}