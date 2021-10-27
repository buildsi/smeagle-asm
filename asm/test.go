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
	assembly := Generate(outfile)
	content := TemplateContent{assembly}

	// Template a second binary, compile again
	t := template.Must(template.New("binary").Parse(getTemplate()))

	assemblyFile := filepath.Join(testdir, "binary.s")
	WriteTemplate(assemblyFile, t, &content)
	if !utils.Exists(assemblyFile) {
		log.Fatalf("Cannot write template to assembly file!\n")
	}

	// Compile again
	_, err = utils.RunCommand([]string{"g++", "binary.s", "-L.", "-ltest", "-o", "edits"}, []string{}, testdir, "")
	if err != nil {
		log.Fatalf("Error compiling second assembly file: %s\n", err)
	}

	// Run both and compare result - should run in container at this point
	resOriginal, err := utils.RunCommand([]string{"./binary"}, []string{}, testdir, "")
	if err != nil {
		log.Fatalf("Issue running original binary %x\n", err)
	}

	resEdited, err := utils.RunCommand([]string{"./edits"}, []string{}, testdir, "")
	if err != nil {
		log.Fatalf("Issue running assembly-generated binary %x\n", err)
	}

	// print result here / calculate metrics!
	fmt.Println(resOriginal, resEdited)

	// Can we generate smeagle assembly again?
}
