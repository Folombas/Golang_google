package main

import "fmt"

func main() {
	var myName = "Alex"
	fmt.Println("Hello! My name is", myName)
	var name string = "Maria"
	fmt.Println("name =", name)

	userName := "admin"
	fmt.Println("username =", userName)

	var sum int
	fmt.Println("The sum is", sum)

	a, b := 5, 7
	fmt.Println("Число №1 =", a, "Число №2 =", b)

	a, h := 9, 4
	fmt.Println("Число №3 =", a, "Число №4 =", h)

	sum2 := a + b + h
	fmt.Println("sum", sum2)
}
