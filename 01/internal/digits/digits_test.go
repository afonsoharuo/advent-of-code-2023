package digits

import "testing"

func TestExtractNumberSuccess(t *testing.T) {
	testCases := make(map[string]int)
	testCases["1"] = 11
	testCases["9"] = 99
	testCases["123"] = 13
	testCases["44"] = 44
	testCases["1a2"] = 12
	testCases["z1a2"] = 12
	testCases["z1a2z"] = 12
	testCases["aejfaejf2ha"] = 22
	testCases["z231a5421z"] = 21

	for line, expResult := range testCases {
		result, _ := ExtractNumber(line)
		if result != expResult {
			t.Fatalf(`ExtractDigits("") = %d, want %d`, result, expResult)
		}
	}
}

func TestExtractNumberErrorNoDigit(t *testing.T) {
	line := "abcd"
	result, err := ExtractNumber(line)
	if err == nil {
		t.Fatalf(`ExtractDigits("%s") = %d, %v, want 0, error`, line, result, err)
	}
}

func TestExtractNumberEmptyLine(t *testing.T) {
	line := ""
	result, err := ExtractNumber(line)
	if err == nil {
		t.Fatalf(`ExtractDigits("%s") = %d, %v, want 0, error`, line, result, err)
	}
}
