CC = gcc

libcutest.a: CuTest.o
	ar -r libcutest.a CuTest.o

-include rules.mk

tests: AllTests.o CuTestTest.o CuTest.o
	gcc -o tests AllTests.o CuTestTest.o CuTest.o

clean:
	rm -f *.o tests *.a

all: clean libcutest.a tests

test: tests
	./tests

depend: 
	gcc -MM *.c > rules.mk
