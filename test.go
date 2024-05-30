package main

import (
	"fmt"
	"sync"
)

func main() {

	f := Foo{look: sync.Mutex{}}
	f.bar()

}

type Foo struct {
	look sync.Mutex
}

func (f *Foo) bar() {
	f.look.Lock()
	defer f.look.Unlock()
	fmt.Println("bar")
}

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsuppered op：%s", op)
	}
}
