package cli

import (
	"fmt"
	"github.com/DataDrake/cli-ng/v2/cmd"
	"github.com/buildsi/smeagleasm/asm"
)

// Args and flags for generate
type GenArgs struct {
	JsonFile []string `desc:"smeagle json to generate assembly for"`
}

type GenFlags struct{}

// Parser looks at symbols and ABI in Go
var Generator = cmd.Sub{
	Name:  "gen",
	Alias: "g",
	Short: "generate assembly from smeagle output",
	Flags: &GenFlags{},
	Args:  &GenArgs{},
	Run:   RunGen,
}

func init() {
	cmd.Register(&Generator)
}

// RunParser reads a file and creates a corpus
func RunGen(r *cmd.Root, c *cmd.Sub) {
	args := c.Args.(*GenArgs)
	assembly := asm.Generate(args.JsonFile[0], nil)
	fmt.Println(assembly)
}
