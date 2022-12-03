package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// return the result of the round with the inputs 'a' and 'b'
func play(a, b byte) int {
	var res int = 0
	a = a - 'A' + 1 // 1 = Rock, 2 = Paper, 3 = Scissors
	b = b - 'X' + 1 // 1 = Lose, 2 = Draw, 3 = Win

	if b == 1 { // lose
		b = a - 1
		if b == 0 {
			b = 3
		}
		res += int(b)

	} else if b == 2 { // draw
		b = a
		res += int(b) + 3

	} else if b == 3 { // win
		b = a + 1
		if b == 4 {
			b = 1
		}
		res += int(b) + 6
	}

	return res
}

func main() {
	somme := 0
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println("error when reading the input file")
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line != "" {
			opponent := line[0]
			recommendation := line[2]
			somme += play(opponent, recommendation)
		}
	}
	fmt.Println(somme)
}
