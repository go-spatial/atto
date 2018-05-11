#ifndef __MBGL_C_TYPES_H
#define __MBGL_C_TYPES_H

#include <stdint.h>

#ifdef __cplusplus
extern "C"{
#endif

typedef struct MbglScheduler MbglScheduler;

struct MbglRendererFrontend;
typedef struct MbglRendererFrontend MbglRendererFrontend;
typedef struct MbglRendererFrontend MbglHeadlessFrontend;

typedef struct MbglMap MbglMap;

typedef struct MbglMapObserver MbglMapObserver;

struct MbglFileSource;
typedef struct MbglFileSource MbglFileSource;
typedef struct MbglFileSource MbglDefaultFileSource;
typedef struct MbglFileSource MbglOnlineFileSource;

typedef struct MbglScheduler MbglThreadPool;

struct MbglLatLng;
typedef struct MbglLatLng MbglLatLng;

struct MbglImage;
typedef struct MbglImage MbglImage;
typedef struct MbglImage MbglPremultipliedImage;

typedef struct MbglRunLoop MbglRunLoop;

typedef struct MbglStyle MbglStyle;

struct MbglSize {
	uint32_t width;
	uint32_t height;
};

typedef struct MbglSize MbglSize;


#ifdef __cplusplus
}
#endif
#endif /* __MBGL_C_TYPES_H */
