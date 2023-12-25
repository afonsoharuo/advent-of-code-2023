package digits

import (
	"errors"
	"log"
	"strconv"
	"unicode"
)

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
