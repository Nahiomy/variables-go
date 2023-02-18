package main

import "fmt"

func main() {
	// path -> ruta
	fmt.Println("Hola mundo desde Golang")
	fmt.Print("Esta es la segunda linea\n\n")
	fmt.Println("Esta es la tercera linea")

	//Golang es un lenguaje de progamación fuertemente tipado

	//Tipos de variables primitivos
	//String -> cadena de texto -> ""
	//Integer -> enteros
	//Boolean -> Lógicos -> verdadero o falso -> true or false
	//Double o float -> números decimales

	//Declaración de variables

	//1. Declaración básica
	var name string

	//Asignar valor a una variable
	name = "Naho"

	fmt.Println("El nombres es: ", name)


	//2. Declarar la variable y asignarle un valor
	var last_name string = "Janse"

	fmt.Println("El apellido es: ", last_name)


	//3. La manera propia de Golang (:=)
	age := 19 

	fmt.Println ("La edad es:", age)

	fmt.Println("Hola", name, last_name, "Tienes", age, "años de edad.")

	//Format verbs
	//%v -> general
	//%s -> string
	//%d -> entero
	//%g -> float
	//%t -> boolean


	//VALOR CERO
	//String -> ""
	//Integer -> 0
	//Float -> 0
	//Boolean -> false

	var cadena string
	fmt.Printf("Valor cero de string:%s\n", cadena) 


	var numero int 
	fmt.Printf("Valor cero de integer:%d\n", numero)


	var decimal float32
	fmt.Printf("El valor de float es:%g\n", decimal)

	var logico bool 
	fmt.Printf ("El valor cero de boolean:%t\n", logico)


	//Recibir datos del usuario
	fmt.Println("Por favor ingresa tu ciudad")
	var city string 
	fmt.Scanf("%s", &city)

	fmt.Println("La ciudad es:", city)

}