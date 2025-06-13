package main

import (
	"fmt"
	"reflect"
)

func main() {
	array := [5]string{"a", "b", "c", "d", "e"}
	slice := array[0:3]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array))

	i, j := 1, 4
	slice2 := array[i:j]
	fmt.Println(slice2)

	slice3 := array[0:3]
	fmt.Println(slice3)

	slice4 := array[2:]
	fmt.Println(slice4)

	slice5 := array[0:3]
	slice6 := array[2:5]
	fmt.Println(slice5, slice6)

	// change array -> changed slice
	array1 := [5]string{"a", "b", "c", "d", "e"}
	slice7 := array1[1:3]
	array1[1] = "X"
	fmt.Println(array1)
	fmt.Println(slice7)

	// slice of slice
	fmt.Println(slice7[1:])

	sl1 := array1[2:4]
	fmt.Println(sl1[0:])
	fmt.Println(sl1[:len(sl1)])
	fmt.Println(sl1[:])

	//--------------
	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	fmt.Println(s1, s2, s3, s4)
	s4[0] = "XX"
	fmt.Println(s1, s2, s3, s4)
}
