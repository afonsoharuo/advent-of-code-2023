package digits

import (
	"errors"
	"log"
	"strconv"
	"unicode"
)

// ExtractNumber returns a number from a string by combining the first and last
// digits in it. If a single digit is found, the number is that digit repeated.
// Returns an error if no digits are found in the string.
func ExtractNumber(line string) (int, error) {
	var digits []string
	for _, ch := range line {
		if unicode.IsDigit(ch) {
			digits = append(digits, string(ch))
		}
	}

	if len(digits) == 0 {
		return 0, errors.New("no digits in line")
	}

	nStr := digits[0] + digits[len(digits)-1]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		log.Fatal(err)
	}

	return n, nil
}
