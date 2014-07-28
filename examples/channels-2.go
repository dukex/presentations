package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func from(connection chan int) {
	connection <- rand.Intn(100)
}

func to(connection chan int) {
	i := <-connection
	fmt.Printf("Someone sent me %d\n", i)
}

func main() {
	connection := make(chan int)
	go from(connection)
	go to(connection)

	time.Sleep(2 * time.Second)
}

// END OMIT
