package main

import (
	"fmt"
	"os"
)

func main() {
	num1, num2 := obtenerNumeros()

	fmt.Println("Seleccione una operación:")
	fmt.Println("1. Suma")
	fmt.Println("2. Resta")
	fmt.Println("3. Multiplicación")
	fmt.Println("4. División")

	operacion := obtenerOperacion()

	var resultado float64

	switch operacion {
	case 1:
		resultado = suma(num1, num2)
	case 2:
		resultado = resta(num1, num2)
	case 3:
		resultado = multiplicacion(num1, num2)
	case 4:
		resultado = division(num1, num2)
	default:
		fmt.Println("Operación no válida")
		os.Exit(1)
	}

	fmt.Printf("El resultado es: %f\n", resultado)
}

func obtenerNumeros() (float64, float64) {
	var num1 float64
	var num2 float64

	fmt.Print("Ingrese el primer número: ")
	fmt.Scanln(&num1)

	fmt.Print("Ingrese el segundo número: ")
	fmt.Scanln(&num2)

	return num1, num2
}

func obtenerOperacion() int {
	var operacion int
	fmt.Scanln(&operacion)
	return operacion
}

func suma(a, b float64) float64 {
	//TODO calcular la suma y retornar el resultado
}

func resta(a, b float64) float64 {
	//TODO calcular la resta y retornar el resultado
}

func multiplicacion(a, b float64) float64 {
	//TODO calcular la multiplicacion y retornar el resultado
}

func division(a, b float64) float64 {
	//TODO calcular la division y retornar el resultado
	
	//TODO si el divisor es cero imprimir un mensaje
	//que diga "No se puede dividir por cero"
}
