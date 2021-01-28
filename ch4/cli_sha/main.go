package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	hashPtr := flag.String("hash", "sha256", "hash type")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	var hash []byte
	switch *hashPtr {
	case "sha256":
		tmp := sha256.Sum256([]byte(text))
		hash = tmp[:]
	case "sha354":
		tmp := sha512.Sum384([]byte(text))
		hash = tmp[:]
	case "sha512":
		tmp := sha512.Sum512([]byte(text))
		hash = tmp[:]
	default:
		panic(fmt.Sprintf("Unsupported hash %s", *hashPtr))
	}

	fmt.Printf("%s : %x\n", *hashPtr, hash)
}
