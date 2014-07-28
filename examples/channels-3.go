package main

import (
	"fmt"
)

// START OMIT
func process(tasks []string) <-chan string {
	ch := make(chan string)
	go func() {
		for index, task := range tasks {
			ch <- fmt.Sprintf("processed task %d: %s", index, task)
		}
		close(ch)
	}()
	return ch
}

func main() {
	results := process([]string{"foo", "bar", "baz"})

	for result := range results { // HL
		fmt.Println(result) // HL
	} // HL
}

// END OMIT
