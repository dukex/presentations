package main

import (
	"fmt"
	"time"
)

// START OMIT
func ola(name string, start chan bool) {
	<-start // HL
	fmt.Println("Buscando", name)
	time.Sleep(2 * time.Second)
	fmt.Println("Olá", name)
}

func main() {
	start := make(chan bool)
	go ola("João", start)
	go ola("Pedro", start)

	fmt.Println("Iniciando Saudações")

	start <- true
	start <- true

	time.Sleep(4 * time.Second)
}

// END OMIT
