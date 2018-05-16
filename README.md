#### Atto

Atto uses [Mapbox-GL-Native](https://github.com/mapbox/mapbox-gl-native) and [gofpdf](https://github.com/jung-kurt/gofpdf) to render Mapbox vector tiles to PDF files.

1. Download and build [Mapbox-GL-Native](https://github.com/mapbox/mapbox-gl-native) for your platform.

2. Add the absolute or relative path to the `include` and `platform/default` to the CXXFLAGS environment variable as `-I` flags
  
3. Add the absolute or relative path to the build directory eg: `build/linux-x86_64/Release` to the LDFLAGS environment variable as `-L` flags

4. Install [mason-js](https://github.com/mapbox/mason-js) and run `mason-js install` and then `mason-js link` in the atto root directory to download the required dependencies 

5. `go build` to build...

