package srtgo

/*
#cgo LDFLAGS: -L${SRCDIR}/lib -lsrt -Wl,-rpath=${SRCDIR}/lib
#include "srt.h"

extern void srtLogCBWrapper (void* opaque, int level, char* file, int line, char* area, char* message);

void srtLogCB(void* opaque, int level, const char* file, int line, const char* area, const char* message)
{
	srtLogCBWrapper(opaque, level, (char*)file, line, (char*)area,(char*) message);
}
*/
import "C"
