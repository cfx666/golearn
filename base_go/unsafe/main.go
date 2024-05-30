package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 10
	ip := &i
	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	println(i)

	p := new(Person)
	pName := (*string)(unsafe.Pointer(p))
	*pName = "乐破晓"
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.Age)))
	*pAge = 20
	fmt.Println(p)

	fmt.Println(unsafe.Sizeof(int(10000000000000000)))
}

type Person struct {
	Name string
	Age  int
}
