all:
	g++ -fPIC -g -shared -O3 -g -o libfoo.so foo.c
	g++ -O3 -g -o binary main.c -L. -lfoo -Wl,-rpath,`pwd`
