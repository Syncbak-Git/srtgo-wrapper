package srtgo

/*

#cgo CXXFLAGS: -DSRT_ENABLE_ENCRYPTION=0
#cgo CFLAGS: -DSRT_ENABLE_ENCRYPTION=0
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
