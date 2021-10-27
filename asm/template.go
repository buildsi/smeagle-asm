package asm

import (
	"bytes"
	"github.com/buildsi/smeagleasm/utils"
	"html/template"
	"log"
	"os"
)

// WriteTemplate to a filepath
func WriteTemplate(path string, t *template.Template, content *TemplateContent) {

	var buf bytes.Buffer
	if err := t.Execute(&buf, content); err != nil {
		log.Fatalf("Cannot write template to buffer: %x", err)
	}
	utils.WriteFile(path, buf.String())
}

// WriteFile and some content to the filesystem
func WriteFile(path string, content string) error {

	filey, err := os.Create(path)
	if err != nil {
		return err
	}
	defer filey.Close()

	_, err = filey.WriteString(content)
	if err != nil {
		return err
	}
	err = filey.Sync()
	return nil
}

type TemplateContent struct {
	Call string
}

// Template with empty "main" function to generate
func getTemplate() string {
	return `.file	"test.c"
	.text
	.globl	main
	.type	main, @function
main:
.LFB0:
	endbr64
	pushq	%rbp
	movq	%rsp, %rbp

       {{ .Call }}

	addq	$16, %rsp
	movl	$0, %eax
	leave
	ret
.LFE0:
	.size	main, .-main
	.ident	"Dinosaur Produced"`
}
