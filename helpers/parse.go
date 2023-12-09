package helpers

import (
	"bufio"
	"log"
	"os"
)

func ParseInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func IsDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func Between(lower, upper, needle int) bool {
	if needle > upper || needle < lower {
		return false
	} else {
		return true
	}

}
