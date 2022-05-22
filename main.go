package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// scanner para leer lo que ingrese el usuario
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese la operacion (suma, de la forma numero + numero, ej: 2+2): ")
	scanner.Scan()

	// sacar el texto ingresado
	operacion := scanner.Text()

	fmt.Println("La operacion ingresada es: ", operacion)

	// split Tomar un string y fraccionarlo
	valores := strings.Split(operacion, "+")

	fmt.Println("Estos son los valores ingresados: ", valores)
	fmt.Println("Primer y segundo valor sumados como texto: ", valores[0]+valores[1])

	// convertimos valores string en enteros
	operador1, _ := strconv.Atoi(valores[0])
	operador2, _ := strconv.Atoi(valores[1])

	fmt.Println("Suma de los dos operadores matematicamente: ", operador1+operador2)

}
