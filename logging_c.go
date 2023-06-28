package srtgo

/*

#cgo CXXFLAGS: -DSRT_ENABLE_ENCRYPTION=0
#cgo CFLAGS: -DSRT_ENABLE_ENCRYPTION=0
#include "srt.h"

extern void srtLogCBWrapper (void* opaque, int level, char* file, int line, char* area, char* message);

void srtLogCB(void* opaque, int level, const char* file, int line, const char* area, const char* message)
{
	srtLogCBWrapper(opaque, level, (char*)file, line, (char*)area,(char*) message);
}
*/
import "C"
