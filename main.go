package main

import "fmt"

func main() {
	var nAltura float64
	var nPeso float64

	fmt.Print("Digite sua altura em metros: ")
	fmt.Scanln(&nAltura)

	fmt.Print("Digite seu peso em kg: ")
	fmt.Scanln(&nPeso)
	fmt.Print(CalcularIMC(nAltura, nPeso))
}

func CalcularIMC(nAltura float64, nPeso float64) string {
	nIMC := nPeso / (nAltura * nAltura)
	return fmt.Sprintf("Seu IMC é de %.2f, e esta Classificado como %s", nIMC, ClassificarIMC(nIMC))
}

func ClassificarIMC(nIMC float64) string {
	var cClassificacao string

	if nIMC < 18.5 {
		cClassificacao = "MAGREZA"
	} else if nIMC < 24.9 {
		cClassificacao = "NORMAL"
	} else if nIMC < 29.9 {
		cClassificacao = "SOBREPESO"
	} else if nIMC < 39.9 {
		cClassificacao = "OBESIDADE"
	} else if nIMC >= 40.0 {
		cClassificacao = "OBESIDADE GRAVE"
	} else {
		cClassificacao = "Não foi possível obter a classificação do seu IMC."
	}

	return cClassificacao
}
