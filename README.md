# Smeagle Asm

This is an example of generating assemly from Smeagle output. This is a terrible idea.

## Usage

### 1. Generate Smeagle Output

To generate the Smeagle output (I used a container):

```bash
$ docker pull ghcr.io/buildsi/smeagle

# Bind to directory with library to generate for
$ docker run --rm -it -v $PWD:/data ghcr.io/buildsi/smeagle bash
```

Then inside the container at `/code`, you will want to compile Smeagle, and run
it on your bound library in `/data`, also saving to `/data` so that you are writing
to your local filesystem (outside the container).

```
$ make
$ ./build/standalone/Smeagle -l /data/libtest.so > /data/smeagle-output.json
```

This is how I generated the testing [smeagle-output.json](smeagle-output.json)

### 2. Generate Terminal Assembly

For an old Python script, see [scripts](scripts).
Currently we can generate with go:

```bash
$ go run main.go gen smeagle-output.json
```
```
endbr64
subq  $56, %rsp # Allocate 56 bytes of space on the stack for local variables
mov   $0x81,%rdi
mov   $0x87,%rsi
mov   $0x47,%rdx
mov   $0x59,%rcx
mov   $0x81,%r8
pushq $0x18
callq bigcall # This is not right for the symbol, obv.
```

or compile first and run the binary:

```bash
$ make
$ ./smeagleasm gen smeagle-output.json
```

### 3. Codegen + Assembly

What we really want to do for a more robust testing of smeagle is:

1. Use codegen to generate some number of c/c++ scripts that do something in main, matched with a library? The functions should confirm values passed.
2. Compile and generate smeagle output for step 1
3. Run assembly generator scripts here and plug result into a main function template
4. Compile template and run and confirm same answer

#### Codegen

You can use (or add new examples) to [examples](examples). For each example, you should include a codegen.yaml that uses random generation,
and **includes a Makefile to compile some main function to a binary named `binary`, and a library to libfoo.so**

#### Running

Steps 2-4 are done by the library here. Since we also need Smeagle, we do the whole shanbang in a container.

##### 1. Build the container

```bash
$ docker build -t ghcr.io/buildsi/smeagleasm .
```

##### 2. Shell into it

```bash
$ docker run -it --rm ghcr.io/buildsi/smeagleasm bash
```

If you want to bind code locally (e.g., to develop/make changes and then try running):

```bash
$ docker run -it --rm -v $PWD:/src ghcr.io/buildsi/smeagleasm
```

The working directory, /src has the executable "smeagleasm" and "Smeagle" is in /code/build/standalone.

```bash
 ls
Dockerfile  README.md  cli	 go.mod  libtest.so  scripts		  smeagleasm  utils
Makefile    asm        examples  go.sum  main.go     smeagle-output.json  test	      version
root@59c7c62db9a2:/src# which Smeagle
/code/build/standalone/Smeagle
```

Now let's run the test generator!

```bash
$ go run main.go test examples/cpp/simple/codegen.yaml 
```

**TODO**

- Corpus Loader needs to skip "empty" functions (e.g., init and fini currently don't add anything)
- Compare output between generated and original
- Try generating smeagle output again?
