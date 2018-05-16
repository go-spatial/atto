.PHONY: clean

TARGET=atto

LDFLAGS="-lstdc++ \
	-lGL"

CXXFLAGS="-std=c++14 -fPIC"

$(TARGET):
	GO_CXXFLAGS=$(CXXFLAGS) GO_LDFLAGS=$(LDFLAGS) CGO_CFLAGS=$(CFLAGS) go build -x .

clean:
	rm -f *.o *.so *.a $(TARGET)
