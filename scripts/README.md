# Generate Assembly with Python

**this is no longer in use**

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

