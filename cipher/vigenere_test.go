package cipher

import "testing"

func TestGenerateTabulaRecta(t *testing.T) {
	cases := []struct {
		alphabet          []rune
		wantedTabulaRecta [][]rune
	}{
		{[]rune{'A', 'B', 'C', 'D', 'E'},
			[][]rune{{'A', 'B', 'C', 'D', 'E'}, {'B', 'C', 'D', 'E', 'A'}, {'C', 'D', 'E', 'A', 'B'}, {'D', 'E', 'A', 'B', 'C'}, {'E', 'A', 'B', 'C', 'D'}}},
		{[]rune{'A', 'Ą', 'B'},
			[][]rune{{'A', 'Ą', 'B'}, {'Ą', 'B', 'A'}, {'B', 'A', 'Ą'}}},
	}

	for _, c := range cases {
		given := generateTabulaRecta(c.alphabet)
		if !areTabulaRectasEqual(c.wantedTabulaRecta, given) {
			t.Errorf("Not equal!\nExpected: %q\nGiven: %q\n", c.wantedTabulaRecta, given)
		}
	}
}

func areRunesEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, val := range a {
		if val != b[i] {
			return false
		}
	}
	return true
}

func areTabulaRectasEqual(a, b [][]rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, row := range a {
		if !areRunesEqual(row, b[i]) {
			return false
		}
	}

	return true
}

func TestEncrypt(t *testing.T) {
	cases := []struct {
		characters, text, key, expected string
	}{
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "COKOLWIEK", "KEY", "MSIYPUSII"},
		{"ABCDEFGHIJKLMNOPRSTUWYZ", "Pudzianowski", "MARIUSZ", "Euuhesmdwkte"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "MICHIGAN TECHNOLOGICAL UNIVERSITY", "HOUGHTON", "TWWNPZOA ASWNUHZBNWWGS NBVCSLYPMM"},
	}

	for _, c := range cases {
		vigenere := New(c.characters, c.key)
		encrypted := vigenere.Encrypt(c.text)
		if encrypted != c.expected {
			t.Errorf("Not equal!\nExpected: %q\nGiven: %q\n", c.expected, encrypted)
		}
	}
}
