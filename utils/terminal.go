package utils

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// Get the environment into a map
func GetEnvironment() map[string]string {
	vars := make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		vars[pair[0]] = pair[1]
	}
	return vars
}

// Run one command!
func RunCommand(cmd []string, env []string, dir string, outfile string) (string, error) {

	// Define the command!
	Cmd := exec.Command(cmd[0], cmd[1:]...)
	Cmd.Env = os.Environ()

	// Change te working directory if needed
	if dir != "" {
		Cmd.Dir = dir
	}

	// If we have environment strings, add them
	if len(env) > 0 {
		Cmd.Env = append(Cmd.Env, env...)
	}

	// Do we want to write to file or buffer?
	if outfile != "" {
		outf, err := os.Create(outfile)
		if err != nil {
			log.Fatalf("%x", err)
		}
		defer outf.Close()
		Cmd.Stdout = outf
		Cmd.Run()
		Cmd.Wait()
		return outfile, nil
	}

	out, err := Cmd.Output()
	return string(out), err
}
