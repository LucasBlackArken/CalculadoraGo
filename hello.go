package main

// Importação pacotes
import (
	"fmt"
)

func main() {

	// Variáveis principais
	var numero1, numero2 float64
	var operador string

	// Demostração Calculadora
	fmt.Println("Bem-vindo à calculadora em Go!")
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&numero1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&numero2)

	fmt.Print("Escolha uma operação (+, -, *, /): ")
	fmt.Scanln(&operador)

	resultado := 0.0

	// Lógica calculadora
	switch operador {
	case "+":
		resultado = numero1 + numero2
	case "-":
		resultado = numero1 - numero2
	case "*":
		resultado = numero1 * numero2
	case "/":
		if numero2 != 0 {
			resultado = numero1 / numero2
		} else {
			fmt.Println("Erro: Divisão por zero!")
			return
		}
	default:
		fmt.Println("Operador inválido")
		return
	}

	// Retorno e visualização resultado
	fmt.Printf("%.2f %s %.2f = %.2f\n", numero1, operador, numero2, resultado)
}
