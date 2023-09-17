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
)

const NJ = 5           // numero de jogadores
const M = 4            // numero de cartas

type carta string      // carta é um strirng

var ch [NJ]chan carta  // NJ canais de itens tipo carta  

func horaDeBater(cartas []carta) bool {
	 if (cartas[0] == cartas[1] && cartas[1] == cartas[2] && cartas[2] == cartas[3]) {
		 return true
	 } else {
		 return false
	 }
}

func jogador(id int, in chan carta, out chan carta, cartasIniciais []carta, batida chan bool) {
	mao := cartasIniciais    // estado local - as cartas na mao do jogador
	nroDeCartas := M         // quantas cartas ele tem 
  cartaRecebida := " "     // carta recebida é vazia

	for {
		select{
		case bateu := <-batida:
			fmt.Println(id, " bateu")
			return
		default:
			fmt.Println(id, " joga")
			if (nroDeCartas == 4) {
				out <- mao[3]
			} else {
				if(horaDeBater(mao)){
					batida <- true
				} else {
					out <- cartaRecebida
				}
			}
		}
	}
}

func main() {
	
	for i := 0; i < NJ; i++ {
		ch[i] = make(chan struct{})
	}
	// cria um baralho com NJ*M cartas
	for i := 0; i < NJ; i++ {
		// escolhe aleatoriamente (tira) cartas do baralho, passa cartas para jogador
		go jogador(i, ch[i], ch[(i+1)%N], cartasEscolhidas , batida) // cria processos conectados circularmente
	}
	
	<-make(chan struct{}) // bloqueia
}


