package asm

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/buildsi/codegen/generate"
	"github.com/vsoch/gosmeagle/corpus"
)

// Generate assembly from a smeagle output json
func Generate(jsonFile string, vars map[string]generate.Function) string {

	// This is output generated from CPP smeagle
	c := corpus.Load(jsonFile)
	return generateAssembly(&c, vars)
}

// generateAssembly for a corpus, assumed to be a main function for this simple program!
func generateAssembly(c *corpus.LoadedCorpus, vars map[string]generate.Function) string {

	output := ""
	for _, loc := range c.Functions {

		// Assume the smeagle testing function called Function for now
		if !strings.Contains(loc.Name, "Function") {
			continue
		}

		// Allocate space on the stack all at once for local variables
		spaceBytes := int64(0)
		for _, param := range loc.Parameters {
			spaceBytes += param.GetSize()
		}
		output += fmt.Sprintf(" subq  $%d, %%rsp # Allocate %d bytes of space on the stack for local variables\n", spaceBytes, spaceBytes)

		lookup := map[string]generate.FormalParam{}
		if vars != nil {
			for _, variable := range vars["Function"].FormalParams {
				lookup[variable.Name] = variable
			}
		}

		for _, param := range loc.Parameters {

			// Currently just support for Integer types
			if param.GetClass() == "Integer" {

				var constant string
				if val, ok := lookup[param.GetName()]; ok {
					constant = val.Value
				} else {
					constant = string(rand.Intn(100))
				}
				// Move constant into rbp+ frameoffset
				// mov $0x55 8(%rbp) - 8 changes relative to constant in smeagle fact
				// going to be 8 bytes off
				if strings.HasPrefix(param.GetLocation(), "framebase") {
					output += fmt.Sprintf("        pushq $%s\n", constant)
				} else {
					output += fmt.Sprintf("        mov   $%s,%s\n", constant, param.GetLocation())
				}
			} else {
				fmt.Println(param.GetClass())
			}
		}
		output += fmt.Sprintf("        callq %s@PLT\n", loc.Name)
	}
	return output
}
