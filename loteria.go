package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Hello world para seguir a tradição...")
	fmt.Println("")

	fmt.Println("Em qual jogo deseja apostar? (1: mega sena - 2: lotofácil)")
	var game string

	fmt.Scanln(&game)
	i, err := strconv.Atoi(game)
	if err != nil {
		fmt.Println("número inválido...")
	}

	if i == 1 {
		jogarMegaSena()
	} else {
		jogarLotofacil()
	}
}

func jogarLotofacil() {
	min := 1
	max := 25
	var numerosParaApostar [15]int

	for i := 0; i < 15; i++ {
		numeroAleatorio := rand.IntN(max-min) + min

		if len(numerosParaApostar) > 0 && numeroRepetidoLotofacil(numerosParaApostar, numeroAleatorio) {
			fmt.Println("já existe ", numeroAleatorio, " na lista")

			for {
				fmt.Println("gerando outro número...")
				numeroAleatorio := rand.IntN(max-min) + min
				fmt.Println("novo número gerado: ", numeroAleatorio)
				if !numeroRepetidoLotofacil(numerosParaApostar, numeroAleatorio) {
					fmt.Println("adiciona na lista pois não existia: ", numeroAleatorio)
					fmt.Println("")
					numerosParaApostar[i] = numeroAleatorio
					break
				} else {
					fmt.Println("novo número gerado já existia: ", numeroAleatorio, "! Tentando gerar outro...")
					fmt.Println("")
				}
			}

		} else {
			numerosParaApostar[i] = numeroAleatorio
		}
	}

	fmt.Println("### Exibindo números gerados na sequencia para loto fácil ###")
	fmt.Println(numerosParaApostar)

	fmt.Println("### Exibindo números após ordenação para loto fácil ###")
	sort.Ints(numerosParaApostar[:])
	fmt.Println(numerosParaApostar)
}

func jogarMegaSena() {
	min := 1
	max := 60
	var numerosParaApostar [6]int

	for i := 0; i < 6; i++ {
		numeroAleatorio := rand.IntN(max-min) + min

		if len(numerosParaApostar) > 0 && numeroRepetido(numerosParaApostar, numeroAleatorio) {
			fmt.Println("já existe ", numeroAleatorio, " na lista")

			for {
				fmt.Println("gerando outro número...")
				numeroAleatorio := rand.IntN(max-min) + min
				fmt.Println("novo número gerado: ", numeroAleatorio)
				if !numeroRepetido(numerosParaApostar, numeroAleatorio) {
					fmt.Println("adiciona na lista pois não existia: ", numeroAleatorio)
					fmt.Println("")
					numerosParaApostar[i] = numeroAleatorio
					break
				} else {
					fmt.Println("novo número gerado já existia: ", numeroAleatorio, "! Tentando gerar outro...")
					fmt.Println("")
				}
			}

		} else {
			numerosParaApostar[i] = numeroAleatorio
		}
	}

	fmt.Println("### Exibindo números gerados na sequencia ###")
	fmt.Println(numerosParaApostar)

	fmt.Println("### Exibindo números após ordenação ###")
	sort.Ints(numerosParaApostar[:])
	fmt.Println(numerosParaApostar)
}

func numeroRepetido(lista [6]int, numeroGerado int) bool {
	for i := 0; i < len(lista); i++ {
		if lista[i] == numeroGerado {
			return true
		}
	}
	return false
}

func numeroRepetidoLotofacil(lista [15]int, numeroGerado int) bool {
	for i := 0; i < len(lista); i++ {
		if lista[i] == numeroGerado {
			return true
		}
	}
	return false
}
