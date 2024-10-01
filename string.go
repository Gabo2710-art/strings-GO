package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese una cadena de texto:")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if strings.Contains(input, "Go") {
		fmt.Println("La cadena contiene 'Go'")
	} else if strings.HasPrefix(input, "Hola") {
		fmt.Println("La cadena comienza con 'Hola'")
	} else if strings.HasSuffix(input, "mundo") {
		fmt.Println("La cadena termina con 'mundo'")
	} else {
		fmt.Println("La cadena no cumple con ninguna condición específica")
	}

	fmt.Println("Cadena en mayúsculas:", strings.ToUpper(input))
	fmt.Println("Cadena en minúsculas:", strings.ToLower(input))
	fmt.Println("Longitud de la cadena:", len(input))
}
