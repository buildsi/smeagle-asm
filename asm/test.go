package asm

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	//	"os"
	"path/filepath"

	"github.com/buildsi/codegen/generate"
	"github.com/buildsi/smeagleasm/utils"
)

// Run tests with a codegen.yaml
func RunTests(jsonYaml string, n int) {

	// Crete temporary outdir in /tmp
	tmpdir, err := ioutil.TempDir("", "smeagle-asm")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Writing tests to %s\n", tmpdir)
	//defer os.RemoveAll(tmpdir)

	// generate code in output directory
	generate.Generate(jsonYaml, tmpdir, "random")

	// If we don't have a Makefile, cannot build
	testdir := filepath.Join(tmpdir, "1")
	makefile := filepath.Join(testdir, "Makefile")
	if !utils.Exists(makefile) {
		log.Fatalf("Cannot build tests without Makefile, %s does not exist!\n", makefile)
	}

	// System commands to run makefile, changing to testdir
	_, err = utils.RunCommand([]string{"make"}, []string{}, testdir, "")
	if err != nil {
		log.Fatalf("Error compiling Makefile for test: %s\n", err)
	}

	// A binary must exist	named binary
	binary := filepath.Join(testdir, "binary")
	libfoo := filepath.Join(testdir, "libfoo.so")
	if !utils.Exists(binary) || !utils.Exists(libfoo) {
		log.Fatalf("Expected output binary 'binary' or 'libfoo.so' does not exist!\n")
	}

	// generate smeagle output for the binary (Smeagle must be on path!)
	outfile := filepath.Join(testdir, "smeagle-output.json")
	_, err = utils.RunCommand([]string{"Smeagle", "-l", libfoo}, []string{}, testdir, outfile)
	if err != nil {
		log.Fatalf("Error running Smeagle!: %s!\n", err)
	}

	// Read in smeagle output and generate assembly
	vars := generate.Load(filepath.Join(testdir, "codegen.json"))
	assembly := Generate(outfile, vars)
	content := TemplateContent{assembly}

	// Template a second binary, compile again
	t := template.Must(template.New("binary").Parse(getTemplate()))

	assemblyFile := filepath.Join(testdir, "binary.s")
	WriteTemplate(assemblyFile, t, &content)
	if !utils.Exists(assemblyFile) {
		log.Fatalf("Cannot write template to assembly file!\n")
	}

	// Compile again
	_, err = utils.RunCommand([]string{"g++", "binary.s", "-L.", "-lfoo", "-o", "edits"}, []string{}, testdir, "")
	if err != nil {
		log.Fatalf("Error compiling second assembly file: %s\n", err)
	}

	// Ensure we have added PWD to LD_LIBRARY_PATH
	envar := []string{"export LD_LIBRARY_PATH=."}
	// Run both and compare result - should run in container at this point
	resOriginal, err := utils.RunCommand([]string{"./binary"}, envar, testdir, "")
	if err != nil {
		log.Fatalf("Issue running original binary %s\n", err)
	}

	resEdited, err := utils.RunCommand([]string{"./edits"}, envar, testdir, "")
	if err != nil {
		log.Fatalf("Issue running assembly-generated binary %s: %s\n", resEdited, err)
	}

	// print result here / calculate metrics!
	fmt.Println("Original:\n", resOriginal, "\nGenerated:\n", resEdited)
	if resOriginal == resEdited {
		fmt.Printf("Generated assembly output is the same as original! 😍️\n")
	} else {
		fmt.Printf("Generated assembly output is different! 😭️\n")
	}

	// Can we generate smeagle assembly again?
}
