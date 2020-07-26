package cipher

type Vigenere struct {
	TabulaRecta [][]rune
	key         string
}

func New(characters, key string) *Vigenere {
	alphabet := []rune(characters)
	tabulaRecta := generateTabulaRecta(alphabet)

	vigenere := Vigenere{tabulaRecta, key}
	return &vigenere
}

func generateTabulaRecta(alphabet []rune) [][]rune {
	alphabetLen := len(alphabet)
	tabulaRecta := make([][]rune, alphabetLen)

	for j := 0; j < alphabetLen; j++ {
		tabulaRecta[j] = make([]rune, alphabetLen)
		for i := 0; i < alphabetLen-j; i++ {
			tabulaRecta[j][i] = alphabet[j+i]
		}
		for i := 0; i < j; i++ {
			tabulaRecta[j][alphabetLen-j+i] = alphabet[i]
		}
	}

	return tabulaRecta
}

func (v Vigenere) encrypt(plainText string) string {
	return "a"
}
