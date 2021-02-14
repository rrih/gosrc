package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// チャネルを close しているので、空になると終わる
	for elem := range queue {
		fmt.Println(elem)
	}
}
