#include <stdint.h>
#include <stdio.h>

void bigcall(long a, long b, long c, long d, long e, __int128_t f)
{
   printf("%ld %ld %ld %ld %ld 0x%lx%16.0lx\n", a, b, c, d, e, (unsigned long) (f >> 64), (unsigned long) (f & 0xffffffffffffffff));
}

