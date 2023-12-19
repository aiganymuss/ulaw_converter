# Places a static or shared (dynamic) library into ../lib

all: dynamic
	
dynamic:
	gcc -c -fPIC decode.c -o decode.o -lsox
	gcc -shared -o libdecode.so decode.o
	mv -f libdecode.so ../lib
	rm -f decode.o

static:
	gcc -c decode.c -lsox
	ar rc libdecode.a decode.o
	ranlib libdecode.a
	mv -f libdecode.a ../lib
	rm -f decode.o
