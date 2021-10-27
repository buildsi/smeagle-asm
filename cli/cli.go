package cli

import (
	"github.com/DataDrake/cli-ng/v2/cmd"
	"github.com/buildsi/smeagleasm/version"
)

// GlobalFlags contains the flags for commands.
type GlobalFlags struct{}

// Root is the main command.
var Root *cmd.Root

// init create the root command and creates subcommands
func init() {
	Root = &cmd.Root{
		Name:      "smeagleasm",
		Short:     "Testing utility for Smeagle to generate assembly via Go",
		Version:   version.Version,
		Copyright: "© 2021 Vanessa Sochat <@vsoch>",
		License:   "Licensed under a combined LICENSE",
	}
	cmd.Register(&cmd.Help)
	cmd.Register(&cmd.Version)
	cmd.Register(&cmd.GenManPages)
}
