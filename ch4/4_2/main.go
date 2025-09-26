/* Exercise 4.2: Write a program that prints the SHA256 hash of its standard input by default but
   supports a command-line flag to print the SHA384 or SHA512 hash instead.
*/
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var sha384Flag = flag.Bool("sha384", false, "prints SHA384 hash of its standard input")
var sha512Flag = flag.Bool("sha512", false, "prints SHA512 hash of its standard input")

func main() {
	flag.Parse()

	if *sha384Flag && *sha512Flag {
		log.Fatal("You can't use both flags!")
	}
	if *sha384Flag {
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("sha384: %x\n", sha512.Sum384(bytes))
	} else if *sha512Flag {
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("sha512: %x\n", sha512.Sum512(bytes))
	} else {
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("sha256: %x\n", sha256.Sum256(bytes))
	}
}
