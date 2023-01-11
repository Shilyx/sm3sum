package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/tjfoc/gmsm/sm3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("sm3sum <file.ext>")
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)
	hash := sm3.New()

	for {
		n, err := f.Read(buf)
		if err != nil {
			break
		}

		hash.Write(buf[0:n])
	}

	value := hash.Sum(nil)
	fmt.Println(hex.EncodeToString(value), os.Args[1])
}
