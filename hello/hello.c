#include <stdio.h>

#define MSG_HELLO "Hello World !\n"

/*
* Ceci est du commentaire
* sur plusieurs lignes
*/

void sayHello() {
	printf(MSG_HELLO);
}

int main() {
	sayHello();
	return 0;
}
