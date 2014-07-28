package main

import (
//"fmt"
)

func main() {
	meuchannel := make(chan int)
	outrochannel := make(chan bool)
	z := make(chan bool)
	// START OMIT
	for {
		select {
		case x := <-meuchannel:
			// faz algo com x
			fmt.Println(x) // OMIT
		case y, ok := <-outrochannel:
			// faz algo com y
			// ok verifica se o channel está fechado
			fmt.Println(ok, y) // OMIT
		case <-z:
			// faz algo quando z for enviado
		default:
			// nenhum dos anteriores for selecionados
		}
	}
	// END OMIT
}
