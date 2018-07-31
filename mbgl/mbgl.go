package mbgl

/*
#cgo CFLAGS: -fPIC
#cgo CFLAGS: -D_GLIBCXX_USE_CXX11_ABI=1
#cgo CXXFLAGS: -std=c++14 -std=gnu++14
#cgo CXXFLAGS: -g
#cgo CXXFLAGS: -I../mason_packages/.link/include
#cgo LDFLAGS: -L../mason_packages/.link/lib
#cgo LDFLAGS: -lsqlite3 -lz
 */
import "C"

/*
// not necessary in mac
-luv -lpthread -ldl -lcurl -lnu -lpng16 -lm -ljpeg -lwebp -licuuc -ldl

//breaks mac
-lrt -lnsl -lGL -lX11
 */