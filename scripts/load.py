#!/usr/bin/env python3

import random
import sys
import json


def main():

    if len(sys.argv) < 2:
        sys.exit("Please provide Smeagle json output to parse.")

    with open(sys.argv[1], "r") as fd:
        corpus = json.loads(fd.read())

    for loc in corpus["locations"]:
        if "function" not in loc:
            continue

        for param in loc["function"]["parameters"]:

            # if it's some kind of int / long / float just give a constant
            if param["class"] == "Integer":
                const = random.choice([1, 2, 3, 4, 5])
                print("mov $0x%s,%s" % (const, param["location"]))

        print("callq %s" % loc["function"]["name"])


"""
extern void bigcall(long a, long b, long c, long d, long e, __int128_t f);

int main(int argc, char *argv[])
{
   __int128_t c;
   c = 0x0000000000000006;
   c = c << 64;
   c += 0x0000000000000007;
   bigcall(1, 2, 3, 4, 5, c);
   return 0;
}


    1064:	48 83 ec 08          	sub    $0x8,%rsp
    1068:	41 b8 05 00 00 00    	mov    $0x5,%r8d
    106e:	b9 04 00 00 00       	mov    $0x4,%ecx
    1073:	ba 03 00 00 00       	mov    $0x3,%edx
    1078:	6a 06                	pushq  $0x6
    107a:	be 02 00 00 00       	mov    $0x2,%esi
    107f:	bf 01 00 00 00       	mov    $0x1,%edi
    1084:	6a 07                	pushq  $0x7

            print(param) 

    1060:	f3 0f 1e fa          	endbr64 
    1064:	48 83 ec 08          	sub    $0x8,%rsp
    1068:	41 b8 05 00 00 00    	mov    $0x5,%r8d
    106e:	b9 04 00 00 00       	mov    $0x4,%ecx
    1073:	ba 03 00 00 00       	mov    $0x3,%edx
    1078:	6a 06                	pushq  $0x6
    107a:	be 02 00 00 00       	mov    $0x2,%esi
    107f:	bf 01 00 00 00       	mov    $0x1,%edi
    1084:	6a 07                	pushq  $0x7
    1086:	e8 c5 ff ff ff       	callq  1050 <bigcall@plt>

   
{?   4384 1}
{?   4385 1}
{?   4386 1}
{cli   4387 1}
{sub $0x10,%rsp   4388 4}
{mov %rcx,%r9   4392 3}
{xor %eax,%eax   4395 2}
{mov %rsi,%rcx   4397 3}
{pushq 0x18(%rsp)   4400 4}
{lea 0xec5(%rip),%rsi   4404 7}
{pushq 0x28(%rsp)   4411 4}
{push %r8   4415 2}
{mov %rdx,%r8   4417 3}
{mov %rdi,%rdx   4420 3}
{mov $0x1,%edi   4423 5}

{callq 0x1050   4428 5}
{add $0x28,%rsp   4433 4}
{retq   4437 1}
register_tm_clones
__do_global_dtors_aux
"""


if __name__ == "__main__":
    main()
