package main

import (
	"fmt"
	"flag"
	"math/rand"
	"log"
	
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
	for _, loc := range c.Functions {
		for _, param := range loc.Parameters {
			if param.GetClass() == "Integer" {
				constant := rand.Intn(100)		
				fmt.Printf("mov $0x%d,%s\n", constant, param.GetLocation())
			}
		}
		fmt.Printf("callq %s\n", loc.Name)
	}
}
