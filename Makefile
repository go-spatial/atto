TARGET=atto

# PLACE YOUR MAPBOXGL INCLUDES HERE
CGO_CXXFLAGS="\
	-I/Users/julio/lib/c/mapbox-gl-native/include \
	-I/Users/julio/lib/c/mapbox-gl-native/platform/default"

GO_FLAGS=CGO_CXXFLAGS=$(CGO_CXXFLAGS)

$(TARGET):
	$(GO_FLAGS) go build -x .

.PHONY: test
test:
	$(GO_FLAGS) go test ./...

.PHONY: clean
clean:
	rm -f *.o *.so *.a $(TARGET)
