package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Ar3q/go-vigenere/cipher"
)

var key = flag.String("k", "", "Key used to encrypt/decrypt text")
var alphabet = flag.String("a", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "Custom alphabet for creating tabula recta")
var text = flag.String("t", "", "Text to encrypt or decrypt")
var encrypt = flag.Bool("e", false, "Bool value, set to true, when text need to be encrypted")
var decrypt = flag.Bool("d", false, "Bool value, set to true, when text need to be decrypted")

func checkFlags() {
	if *key == "" {
		fmt.Println("Key not provided")
		os.Exit(2)
	}

	if *text == "" {
		fmt.Println("Text to encrypt/decrypt not provided")
		os.Exit(2)
	}

	if *encrypt == *decrypt {
		fmt.Println("Only one option can be used (one flag set to true)")
		os.Exit(2)
	}
}

func main() {
	flag.Parse()

	checkFlags()

	vigenere := cipher.New(*alphabet, *key)

	if *encrypt {
		encrypted := vigenere.Encrypt(*text)
		fmt.Printf("Decrypted (original) text: %s\nEncrypted text: %s\n", *text, encrypted)
	} else if *decrypt {
		decrypted := vigenere.Decrypt(*text)
		fmt.Printf("Encrypted (original) text: %s\nDecrypted text: %s\n", *text, decrypted)
	}
}
