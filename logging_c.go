package srtgo

// #cgo linux CFLAGS: -DLINUX
// #cgo darwin CFLAGS: -DTARGET_OS_MAC
// #cgo freebsd CFLAGS: -DBSD
// #cgo windows CFLAGS: -DERROR_NO_WINDOWS_SUPPORT
// #cgo i386 CFLAGS: -DIA32
// #cgo amd64 CFLAGS: -DAMD64
// #cgo linux CXXFLAGS: -DLINUX
// #cgo darwin CXXFLAGS: -DTARGET_OS_MAC
// #cgo freebsd CXXFLAGS: -DBSD
// #cgo windows CXXFLAGS: -DERROR_NO_WINDOWS_SUPPORT
// #cgo i386 CXXFLAGS: -DIA32
// #cgo amd64 CXXFLAGS: -DAMD64

/*

#cgo CXXFLAGS: -DSRT_ENABLE_ENCRYPTION=0
#cgo CFLAGS: -DSRT_ENABLE_ENCRYPTION=0
#cgo LDFLAGS: -lssl -lcrypto
#include "srt.h"

extern void srtLogCBWrapper (void* opaque, int level, char* file, int line, char* area, char* message);

void srtLogCB(void* opaque, int level, const char* file, int line, const char* area, const char* message)
{
	srtLogCBWrapper(opaque, level, (char*)file, line, (char*)area,(char*) message);
}
*/
import "C"
