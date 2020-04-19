// WARNING: This file has automatically been generated
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package gen

/*
#cgo LDFLAGS: -L${SRCDIR}/.. -lgogc
#include "../gogc.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// GogcCount function as declared in go-gc/gogc.h:33
func GogcCount(num uint) uint {
	cnum, cnumAllocMap := (C.uintptr_t)(num), cgoAllocsUnknown
	if cnumAllocMap != cgoAllocsUnknown {
		defer cnumAllocMap.Free()
	}
	__ret := C.gogc_count(cnum)
	runtime.KeepAlive(cnum)
	__v := (uint)(__ret)
	return __v
}

// GogcStrCount function as declared in go-gc/gogc.h:35
func GogcStrCount(ptr string, num uint) uint {
	ptr = safeString(ptr)
	cptr, cptrAllocMap := unpackPUint8TString(ptr)
	if cptrAllocMap != cgoAllocsUnknown {
		defer cptrAllocMap.Free()
	}
	cnum, cnumAllocMap := (C.uintptr_t)(num), cgoAllocsUnknown
	if cnumAllocMap != cgoAllocsUnknown {
		defer cnumAllocMap.Free()
	}
	__ret := C.gogc_str_count(cptr, cnum)
	runtime.KeepAlive(cnum)
	runtime.KeepAlive(ptr)
	runtime.KeepAlive(cptr)
	__v := (uint)(__ret)
	return __v
}

// GogcVerifyPost function as declared in go-gc/gogc.h:37
func GogcVerifyPost(replicasPtr []GogcPublicReplicaInfo, replicasLen uint) {
	creplicasPtr, creplicasPtrAllocMap := unpackArgSGogcPublicReplicaInfo(replicasPtr)
	if creplicasPtrAllocMap != cgoAllocsUnknown {
		defer creplicasPtrAllocMap.Free()
	}
	creplicasLen, creplicasLenAllocMap := (C.size_t)(replicasLen), cgoAllocsUnknown
	if creplicasLenAllocMap != cgoAllocsUnknown {
		defer creplicasLenAllocMap.Free()
	}
	C.gogc_verify_post(creplicasPtr, creplicasLen)
	runtime.KeepAlive(creplicasLen)
	packSGogcPublicReplicaInfo(replicasPtr, creplicasPtr)
	runtime.KeepAlive(creplicasPtr)
}

// GogcVerifyPostCount function as declared in go-gc/gogc.h:39
func GogcVerifyPostCount(replicasPtr []GogcPublicReplicaInfo, replicasLen uint) *GogcPublicReplicaInfoCountResponse {
	creplicasPtr, creplicasPtrAllocMap := unpackArgSGogcPublicReplicaInfo(replicasPtr)
	if creplicasPtrAllocMap != cgoAllocsUnknown {
		defer creplicasPtrAllocMap.Free()
	}
	creplicasLen, creplicasLenAllocMap := (C.size_t)(replicasLen), cgoAllocsUnknown
	if creplicasLenAllocMap != cgoAllocsUnknown {
		defer creplicasLenAllocMap.Free()
	}
	__ret := C.gogc_verify_post_count(creplicasPtr, creplicasLen)
	runtime.KeepAlive(creplicasLen)
	packSGogcPublicReplicaInfo(replicasPtr, creplicasPtr)
	runtime.KeepAlive(creplicasPtr)
	__v := NewGogcPublicReplicaInfoCountResponseRef(unsafe.Pointer(__ret))
	return __v
}
