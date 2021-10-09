package common

import (
	"fmt"
	"math"
	"reflect"
)

type Element struct {
	Value interface{}
}

func (e *Element) String() string {
	return fmt.Sprint(e.Value)
}

func EqualType(l, r interface{}) bool {
	return reflect.TypeOf(l) == reflect.TypeOf(r)
}

func IsNil(v interface{}) bool {
	return reflect.ValueOf(v).IsNil()
}

func Max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}
