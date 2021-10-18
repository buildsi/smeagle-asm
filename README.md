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

### 2. Generate Assembly

Now we can run the program to generate assembly:

```bash
$ python load.py smeagle-output.json
```
```
mov $0x4,%rdi
mov $0x3,%rsi
mov $0x5,%rdx
mov $0x3,%rcx
mov $0x1,%r8
mov $0x4,framebase+8
callq bigcall
```

or run with go:

```bash
$ go run main.go smeagle-output.json
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

And then in the [test](test) folder we can generate assembly for the test binary (which calls this function)

```bash
g++ -S -o test.a test.c
```

And then arbitrariy delete the function body and add our output above (test.temp)

And try compiling again

```bash
g++ -c test.temp -o test.o
```

This obviously doesn't work because I don't know what I'm doing :)
