#include "fr_imtne_ilog_Display.h"

JNIEXPORT void JNICALL Java_fr_imtne_ilog_Display_displayLine(JNIEnv * env, jobject pThis, jint jintLine, jstring jstrMessage){
	const char *msg = (*env)->GetStringUTFChars(env, jstrMessage, 0);
	printf("%d : %s\n", jintLine, msg);
	(*env)->ReleaseStringUTFChars(env, jstrMessage, msg);
}
