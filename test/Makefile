all:
	g++ -fPIC -shared -O3 -g -o libtest.so testlib.c
	g++ -O3 -g -o test test.c -L. -ltest -Wl,-rpath,`pwd`
	
	# Assembly
	# g++ -S -o test.a test.c
