package asm

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/vsoch/gosmeagle/corpus"
)

// Generate assembly from a smeagle output json
func Generate(jsonFile string) string {

	// This is output generated from CPP smeagle
	c := corpus.Load(jsonFile)
	return generateAssembly(&c)
}

// generateAssembly for a corpus, assumed to be a main function for this simple program!
func generateAssembly(c *corpus.LoadedCorpus) string {

	output := ""

	// This seems to always be printed at the start of main
	output += "endbr64\n"
	for _, loc := range c.Functions {

		// Allocate space on the stack all at once for local variables
		spaceBytes := int64(0)
		for _, param := range loc.Parameters {
			spaceBytes += param.GetSize()
		}
		output += fmt.Sprintf("subq  $%d, %%rsp # Allocate %d bytes of space on the stack for local variables\n", spaceBytes, spaceBytes)
		for _, param := range loc.Parameters {
			if param.GetClass() == "Integer" {
				constant := rand.Intn(100)
				// Move constant into rbp+ frameoffset
				// mov $0x55 8(%rbp) - 8 changes relative to constant in smeagle fact
				// going to be 8 bytes off
				if strings.HasPrefix(param.GetLocation(), "framebase") {
					output += fmt.Sprintf("pushq $0x%d\n", constant)
				} else {
					output += fmt.Sprintf("mov   $0x%d,%s\n", constant, param.GetLocation())
				}
			}
		}
		output += fmt.Sprintf("callq %s\n", loc.Name)
	}
	return output
}
