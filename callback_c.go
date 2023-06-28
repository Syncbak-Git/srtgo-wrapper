package srtgo

/*

#cgo CXXFLAGS: -DENABLE_ENCRYPTION=0 -DUSE_OPENSSL
#cgo CFLAGS: -DENABLE_ENCRYPTION=0 -DUSE_OPENSSL
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
