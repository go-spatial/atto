#ifndef __MBGL_C_TYPES_H
#define __MBGL_C_TYPES_H

#include <stdint.h>

#ifdef __cplusplus
extern "C"{
#endif

typedef void *MbglScheduler;
typedef void *MbglThreadPool;

typedef void *MbglRendererFrontend;
typedef void *MbglHeadlessFrontend;

typedef void *MbglMap;

typedef void *MbglMapObserver;

typedef void *MbglFileSource;
typedef void *MbglDefaultFileSource;
typedef void *MbglOnlineFileSource;

typedef void *MbglLatLng;

typedef void *MbglImage;
typedef void *MbglPremultipliedImage;

typedef void *MbglRunLoop;

typedef void *MbglStyle;

struct _MbglSize {
	uint32_t width;
	uint32_t height;
};

typedef struct _MbglSize MbglSize;


#ifdef __cplusplus
}
#endif
#endif /* __MBGL_C_TYPES_H */
