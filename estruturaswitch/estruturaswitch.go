package main

import (
	"fmt"
)

func main() {
	var mes int
	fmt.Print("Digite o número do mês: ")
	fmt.Scanf("%d", &mes)

	switch mes {
	case 1:
		fmt.Println("Esse mês é janeiro.");
	case 2:
		fmt.Println("Este mês é fevereiro");
	case 3:
		fmt.Println("Este mês é março");
	case 4:
		fmt.Println("Este mês é abril");
	case 5:
		fmt.Println("Este mês é maio");
	case 6:
		fmt.Println("Este mês é junho");
	case 7:
		fmt.Println("Este mês é julho");
	case 8:
		fmt.Println("Este mês é agosto");
	case 9:
		fmt.Println("Este mês é setembro");
	case 10:
		fmt.Println("Este mês é outubro");
	case 11:
		fmt.Println("Este mês é novembro");
	case 12:
		fmt.Println("Este mês é dezembro");
	default:
		fmt.Println("Mês inválido");

	}

	switch mes {
	case 1, 2, 3:
		fmt.Println("Primeiro trimestre");
	case 4, 5, 6:
		fmt.Println("Segundo trimestre");
	case 7, 8, 9:
		fmt.Println("Terceiro trimestre");
	case 10, 11, 12:
		fmt.Println("Quarto trimestre");
	}

}
