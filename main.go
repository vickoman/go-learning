package main

import "fmt"

func main() {
	/**
	/*	 Comentarios en G0
	/*
	*/
	number := sum(50, 30)
	name := getName()
	a, b, c := getVariales()
	const helloworld string = "Hola  %s %s, Bienvenido al fascinante mundo de Go. \n "
	lastname := "Diaz De La Gasca"
	fmt.Print("Hola mundo")
	fmt.Printf(helloworld, name, lastname)
	fmt.Println(number, a, b, c)
}

func getName() string {
	var name string
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	return name
}

func getVariales() (int, int, int) {
	return 1, 2, 3
}

func sum(a int, b int) int {
	return a + b
}
