hello: hello.o
	gcc -o hello hello.o

hello.o: hello.c
	gcc -o hello.o -c hello.c -W -Wall

clean:
	rm -f *.o *.i hello

all: clean hello
