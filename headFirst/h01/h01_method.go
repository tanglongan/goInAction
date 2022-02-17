package main

import "fmt"

type MyString string
type Number int

func (m MyString) sayHi() {
	fmt.Println("Hi from：", m)
}

func (m MyString) method() {
	fmt.Println("Method with value receiver")
}

func (m *MyString) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

func (n *Number) Double() {
	*n *= 2
}

func main() {
	value := MyString("a MyType value")
	value.sayHi()

	antherValue := MyString("another value")
	antherValue.sayHi()

	number := Number(4)
	fmt.Println("Original of number:", number)
	number.Double()
	fmt.Println("Original of number:", number)

	fmt.Println("=================================")

	valueString := MyString("MyString")
	pointer := &valueString

	valueString.method()
	valueString.pointerMethod()
	pointer.method()
	pointer.pointerMethod()

}
