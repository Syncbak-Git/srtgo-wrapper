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
#include "callback.h"

int srtListenCB(void* opaque, SRTSOCKET ns, int hs_version, const struct sockaddr* peeraddr, const char* streamid)
{
	return srtListenCBWrapper(opaque, ns, hs_version, (struct sockaddr*)peeraddr, (char*)streamid);
}

void srtConnectCB(void* opaque, SRTSOCKET ns, int errorcode, const struct sockaddr* peeraddr, int token)
{
	srtConnectCBWrapper(opaque, ns, errorcode, (struct sockaddr*)peeraddr, token);
}
*/
import "C"
