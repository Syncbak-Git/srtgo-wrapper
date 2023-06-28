package srtgo

/*

#cgo CXXFLAGS: -DENABLE_ENCRYPTION=0 -DUSE_OPENSSL
#cgo CFLAGS: -DENABLE_ENCRYPTION=0 -DUSE_OPENSSL
#include "srt.h"

extern void srtLogCBWrapper (void* opaque, int level, char* file, int line, char* area, char* message);

void srtLogCB(void* opaque, int level, const char* file, int line, const char* area, const char* message)
{
	srtLogCBWrapper(opaque, level, (char*)file, line, (char*)area,(char*) message);
}
*/
import "C"
