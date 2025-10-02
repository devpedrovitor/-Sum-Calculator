package main

import "fmt"

func main() {
	var number1, numer2 int
	fmt.Print("Insira um número:")
	_, err := fmt.Scan(&number1, &numer2)

	if err != nil {
		fmt.Print("Erro na entrada de valores!")
	}

	fmt.Println("A Soma do dois valores é: ", Sum(number1, numer2))
}

func Sum(a int , b int) int{
	return a + b;
}