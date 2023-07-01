package main

import (
	"fmt"
)

func soma(a, b float64) float64 {
	return a + b
}

func subtracao(a, b float64) float64 {
	return a - b
}

func multiplicacao(a, b float64) float64 {
	return a * b
}

func divisao(a, b float64) float64 {
	return a / b
}

func main() {
	var operacao string
	var num1, num2 float64

	fmt.Println("Bem-vindo(a) à calculadora!")

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	fmt.Println("Selecione a operação:")
	fmt.Println("1. Soma")
	fmt.Println("2. Subtração")
	fmt.Println("3. Multiplicação")
	fmt.Println("4. Divisão")

	fmt.Print("Digite o número da operação desejada: ")
	fmt.Scanln(&operacao)

	var resultado float64

	switch operacao {
	case "1":
		resultado = soma(num1, num2)
	case "2":
		resultado = subtracao(num1, num2)
	case "3":
		resultado = multiplicacao(num1, num2)
	case "4":
		resultado = divisao(num1, num2)
	default:
		fmt.Println("Operação inválida.")
		return
	}

	fmt.Println("Resultado:", resultado)
}
