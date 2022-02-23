JDK = /usr/lib/jvm/java-11-openjdk-amd64

INCLUDES = -I"$(JDK)/include/linux" -I"$(JDK)/include"

CFLAGS += $(INCLUDES)

OBJS = Display.o

TARGET = ../HelloJNI/libHelloNative.so

Display.o: Display.c fr_imtne_ilog_Display.h

$(TARGET): $(OBJS)
	$(CXX) -shared -o $(TARGET) $(OBJS)

all: $(TARGET)

clean:
	$(RM) $(OBJS) $(TARGET)
