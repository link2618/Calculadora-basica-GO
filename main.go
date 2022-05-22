package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func validarError(valorT string) float64 {
	valor, error := strconv.ParseFloat(valorT, 64)
	if error != nil {
		log.Fatal(error)
	}

	return valor
}

func main() {

	// scanner para leer lo que ingrese el usuario
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Elija la operacion que desea (Solo el numero ej: 1)")
	fmt.Println("1. Suma(+)")
	fmt.Println("2. Resta(-)")
	fmt.Println("3. Multiplicacion(*)")
	fmt.Println("4. Divicion(/)")
	scanner.Scan()

	// sacar el texto ingresado
	opcion := scanner.Text()
	// fmt.Println("La opcion ingresada es: ", opcion)

	fmt.Println("\nIngrese la operacion (ej: 2+2)")
	scanner.Scan()
	operacion := scanner.Text()
	// fmt.Println("La operacion ingresada es: ", operacion)

	var respuesta float64

	// split tomar un string y fraccionarlo
	switch opcion {
	case "1":
		// Suma
		valores := strings.Split(operacion, "+")
		for i := 0; i < len(valores); i++ {
			valor := validarError(valores[i])
			respuesta = respuesta + valor
		}
	case "2":
		// Resta
		valores := strings.Split(operacion, "-")
		for i := 0; i < len(valores); i++ {
			valor := validarError(valores[i])
			if i == 0 {
				respuesta = valor
				continue
			}
			respuesta = respuesta - valor
		}
	case "3":
		// Multiplicacion
		valores := strings.Split(operacion, "*")
		respuesta = 1
		for i := 0; i < len(valores); i++ {
			valor := validarError(valores[i])
			respuesta = respuesta * valor
		}
	case "4":
		// Divicion
		valores := strings.Split(operacion, "/")
		for i := 0; i < len(valores); i++ {
			valor := validarError(valores[i])
			if i == 0 {
				respuesta = valor
				continue
			}
			respuesta = respuesta / valor
		}
	default:
		log.Fatal("\nOpcion no valida ", opcion)
	}

	fmt.Println("\nLa respuesta es:", respuesta)

}
