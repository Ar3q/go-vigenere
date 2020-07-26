package main

import (
	"fmt"
	"github.com/Ar3q/go-vigenere/cipher"
)

func main() {
	fmt.Println("Nice")

	text := "COKOLWIEK"
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	key := "KEY"
	vigenere := cipher.New(characters, key)
	mess := vigenere.Encrypt(text)

	fmt.Printf("Plain text: %s\nEncrypted: %s\n", text, mess)
}
