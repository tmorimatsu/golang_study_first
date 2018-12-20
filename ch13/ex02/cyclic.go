package cyclic

import (
	"reflect"
	"unsafe"
)

// func main() {
// 	a, b, c := node{}, node{}, node{}
// 	a.next = &b
// 	b.next = &c

// 	arr := make(map[string]string)
// 	arr["test1"] = "value1"
// 	print(IsCyclic(arr))
// }

// type node struct {
// 	next *node
// }

func isCyclic(x reflect.Value, seen map[ptr]bool) bool {

	if x.CanAddr() {
		ptr := ptr{unsafe.Pointer(x.UnsafeAddr()), x.Type()}
		if seen[ptr] {
			return true // すでに見た
		}
		seen[ptr] = true
	}

	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return isCyclic(x.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if isCyclic(x.Index(i), seen) {
				return true
			}
		}
		return false

	case reflect.Struct:
		for i := 0; i < x.NumField(); i++ {
			if isCyclic(x.Field(i), seen) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range x.MapKeys() {
			if isCyclic(x.MapIndex(k), seen) {
				return true
			}
		}
		return false
	default:
		return false

	}
	panic("unreachable")
}

func IsCyclic(x interface{}) bool {
	seen := make(map[ptr]bool)
	return isCyclic(reflect.ValueOf(x), seen)
}

type ptr struct {
	x unsafe.Pointer
	t reflect.Type
}
