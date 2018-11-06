package main

import (
	"fmt"

	L "../ROT13"
)

func main() {
	s := "hello"
	s = "ABCxyz12"

	fmt.Printf("Encrypting String:%s\n", s)
	encryptedString := L.EncryptDecrypt(s)
	fmt.Printf("Encrypted String:%s\n", encryptedString)
	decryptedString := L.EncryptDecrypt(encryptedString)
	fmt.Printf("Decrypted String:%s\n", decryptedString)

}
