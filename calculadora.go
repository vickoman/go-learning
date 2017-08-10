package main

import "fmt"

func main() {
	num1 := getNumbers("Ingrese el primer numero")
	num2 := getNumbers("Ingrese el segundo numero")
	fmt.Printf("Se procedera a sumar estos dos numeros %d, %d \n", num1, num2)
	fmt.Printf("La suma de %d y %d es : %d \n", num1, num2, suma(num1, num2))
	fmt.Printf("La resta de %d y %d es : %d \n", num1, num2, resta(num1, num2))
	fmt.Printf("La Multiplicacion de %d y %d es : %d \n", num1, num2, multiplica(num1, num2))
	fmt.Printf("La Division de %d y %d es : %d \n", num1, num2, divide(num1, num2))
}

func getNumbers(msg string) int {
	var num int
	fmt.Printf("%s: ", msg)
	fmt.Scanf("%d", &num)
	return num
}

func suma(a int, b int) int {
	return a + b
}

func resta(a int, b int) int {
	return a - b
}
func multiplica(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	return a / b
}
