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
	b = b - 'X' + 1 // 1 = Rock, 2 = Paper, 3 = Scissors

	if a == b {
		res += int(b) + 3
	} else if b == a+1 || b+2 == a { // we win
		fmt.Println(a, b)
		res += int(b) + 6
	} else if a == b+1 || a+2 == b { // we lose
		res += int(b) + 0
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
