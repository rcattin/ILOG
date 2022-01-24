#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main(int argc, char *argv[]) {
	char *value;
	char str[256];
	char *namesVars[] = {
		"DISPLAY", "HOME", "NVM_DIR", "PATH",
		"PWD", "SHELL", "USER"
	};
	int i, cntVars = sizeof(namesVars)/sizeof(char *);

	printf("exe name is %s\n", argv[0]);

	value = getcwd(str, 256);
	printf("current directory is %s\n", value);

	printf("some important environment variables:\n");
	for (i=0; i<cntVars; i++) {
		char *name = namesVars[i];
		value = getenv(name);
		printf("%s=%s\n", name, value);
	}
	return EXIT_SUCCESS;
}