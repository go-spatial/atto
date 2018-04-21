.PHONY: clean

TARGET=atto

LDFLAGS="-lstdc++ \
	/home/bizarro/Documents/Projects/src/github.com/go-spatial/mbgl-c/build/linux-x86_64/Debug/libmbgl-c.a \
	/home/bizarro/Documents/Projects/ThirdParty/mapbox-gl-native/build/linux-x86_64/Debug/libmbgl-core.a \
	/home/bizarro/Documents/Projects/ThirdParty/mapbox-gl-native/build/linux-x86_64/Debug/libmbgl-filesource.a \
	-lGL"

CXXFLAGS="-std=c++14 -fPIC"

CFLAGS=-I/home/bizarro/Documents/Projects/src/github.com/go-spatial/mbgl-c/include

$(TARGET):
	GO_CXXFLAGS=$(CXXFLAGS) GO_LDFLAGS=$(LDFLAGS) CGO_CFLAGS=$(CFLAGS) go build -x .

clean:
	rm -f *.o *.so *.a $(TARGET)
