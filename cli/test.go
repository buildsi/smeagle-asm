package cli

import (
	"github.com/DataDrake/cli-ng/v2/cmd"
	"github.com/buildsi/smeagleasm/asm"
)

// Args and flags for test
type TestArgs struct {
	CodeYaml []string `desc:"codegen.yaml file in folder to populate tests"`
}

type TestFlags struct {
	N int `long:"number" desc:"Number of generations to do"`
}

// Parser looks at symbols and ABI in Go
var Tester = cmd.Sub{
	Name:  "test",
	Alias: "t",
	Short: "run tests using a codegen.yaml file",
	Flags: &TestFlags{},
	Args:  &TestArgs{},
	Run:   RunTest,
}

func init() {
	cmd.Register(&Tester)
}

// RunParser reads a file and creates a corpus
func RunTest(r *cmd.Root, c *cmd.Sub) {
	args := c.Args.(*TestArgs)
	flags := c.Flags.(*TestFlags)

	// We need to run one test
	if flags.N == 0 {
		flags.N = 1
	}
	asm.RunTests(args.CodeYaml[0], flags.N)
}
