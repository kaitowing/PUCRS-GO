// por Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
// PROBLEMA:
//   o dorminhoco especificado no arquivo Ex1-ExplanacaoDoDorminhoco.pdf nesta pasta
// ESTE ARQUIVO
//   Um template para criar um anel generico.
//   Adapte para o problema do dorminhoco.
//   Nada está dito sobre como funciona a ordem de processos que batem.
//   O ultimo leva a rolhada ...
//   ESTE  PROGRAMA NAO FUNCIONA.    É UM RASCUNHO COM DICAS.

package main

import (
	"fmt"
	"math/rand"
)

const NJ = 5 // numero de jogadores
const M = 3  // numero de cartas

type carta string // carta é um strirng

var ch [NJ]chan carta // NJ canais de itens tipo carta

func horaDeBater(cartas []carta, id int) bool {
	if cartas[0] == cartas[1] && cartas[1] == cartas[2] && cartas[2] == cartas[3] {
		return true
	} else {
		return false
	}
}

func jogador(id int, in chan carta, out chan carta, cartasIniciais []carta, batida chan bool) {
	mao := cartasIniciais    // estado local - as cartas na mao do jogador
	nroDeCartas := M         // quantas cartas ele tem
	if id == 0 {
		nroDeCartas = 4
	}
	fmt.Println(id, "começou")

	for {
		select {
		case bateu := <-batida:
			fmt.Println(id, "bateu")
			batida <- bateu
			return
		case cartaRecebida := <-in:
		//	fmt.Println(id, "recebe carta")
			mao[nroDeCartas] = cartaRecebida
			nroDeCartas++
		//	fmt.Println(id, "Numero de cartas: ", nroDeCartas)
		default:
			if nroDeCartas == 4 {
			//	fmt.Println(id, "joga")
				randomIndex := rand.Intn(nroDeCartas-1) // Gera um índice aleatório
				out <- mao[randomIndex]
				// Move a carta jogada para a última posição e atualiza nroDeCartas
				mao[randomIndex] = mao[nroDeCartas-1]
				nroDeCartas--
				if (horaDeBater(mao, id)) {
					batida <- true
					return
				}
			}
		}
	}
}

func main() {
	batida := make(chan bool)

	for i := 0; i < NJ; i++ {
		ch[i] = make(chan carta)
	}
	// cria um baralho com NJ*M cartas
	for i := 0; i < NJ; i++ {
		// escolhe aleatoriamente (tira) cartas do baralho, passa cartas para jogador
		cartasEscolhidas := make([]carta, M+1)
		for j := 1; j < M+1; j++ {
			cartasEscolhidas[j] = carta(string(rune(j)))
		}

		if i == 0 {
			cartasEscolhidas[3] = carta(string(rand.Intn(2) + 1))
		}

		go jogador(i, ch[i], ch[(i+1)%NJ], cartasEscolhidas, batida) // cria processos conectados circularmente
	}

	<-make(chan bool)
}