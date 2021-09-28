package common

import (
	"fmt"
	"reflect"
)

const ERROR_MSG string = "The two values should be the same."

type Element struct {
	Value interface{}
}

func (e *Element) String() string {
	return fmt.Sprint(e.Value)
}

func EqualType(l, r interface{}) bool {
	return reflect.TypeOf(l) == reflect.TypeOf(r)
}
