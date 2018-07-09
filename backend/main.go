package main

import (
	"fmt"
)

//Возвращение произведение двух чисел
func MultiplyTwoInt(first, second int) int {
	return first * second
}
func main() {
	fmt.Print("Hello world \n")
	fmt.Printf("Sum of two int: %d\n", MultiplyTwoInt(3, 16))
}
