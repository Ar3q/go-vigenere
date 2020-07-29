# go-vigenere

Simple terminal program for text encryption, decryption with Vigenere cipher

## Using

```bash
# build
go build
# encrypt text with custom alphabet
./go-vigenere -k KEY -t TEXT_TO_ENCRYPT -a CUSTOM_ALPHABET -e
# decrypt text with default alphabet 
./go-vigenere -k KEY -t TEXT_TO_DECRYPT -d
```
```bash
# examples with output
./go-vigenere -k ASD -t PUDZIAN -e
Decrypted (original) text: PUDZIAN
Encrypted text: PMGZADN

./go-vigenere -t "Qjwrgf fą awexnąw, ąej xif cfphlnimwnbb!"  \
-a AĄBCDEFGHIJKLMNOPQRSTUVWXYZŻ -k SARNIEŻNIWO -d
Encrypted (original) text: Qjwrgf fą awexnąw, ąej xif cfphlnimwnbb!
Decrypted text: Zjedzą go sarenki, tak jak poprzedniego!
```
