package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/vsoch/gosmeagle/corpus"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatalf("Please provide an input json corpus!")
	}

	// This is output generated from CPP smeagle
	c := corpus.Load(args[0])

	// This seems to always be printed at the start of main
	fmt.Println("endbr64")
	for _, loc := range c.Functions {

		// Allocate space on the stack all at once for local variables
		spaceBytes := int64(0)
		for _, param := range loc.Parameters {
			spaceBytes += param.GetSize()
		}
		fmt.Printf("subq  $%d, %%rsp # Allocate %d bytes of space on the stack for local variables\n", spaceBytes, spaceBytes)
		for _, param := range loc.Parameters {
			if param.GetClass() == "Integer" {
				constant := rand.Intn(100)
				if strings.HasPrefix(param.GetLocation(), "framebase") {
					fmt.Printf("pushq $0x%d\n", constant)
				} else {
					fmt.Printf("mov   $0x%d,%s\n", constant, param.GetLocation())
				}
			}
		}
		fmt.Printf("callq %s\n", loc.Name)
	}
}
