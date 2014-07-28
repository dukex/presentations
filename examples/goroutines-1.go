package main

import (
	"fmt"
	"time"
)

// START OMIT
func ola(name string) {
	fmt.Println("Buscando", name)
	time.Sleep(2 * time.Second)
	fmt.Println("Olá", name)
}

func main() {
	ola("João")
	ola("Pedro")

	time.Sleep(4 * time.Second)
}

// END OMIT
