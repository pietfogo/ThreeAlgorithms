package main

import "fmt"

func main() {
	var recebeFatorial int
	fmt.Print("Digite um número para calcular o fatorial: ")
	fmt.Scan(&recebeFatorial)

	fatorial := 1
	for i := 2; i <= recebeFatorial; i++ {
		fatorial = fatorial *  i
	}

	fmt.Printf("Valor de fatorial: %v", fatorial)
}
