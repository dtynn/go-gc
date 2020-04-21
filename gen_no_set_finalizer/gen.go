// WARNING: This file has automatically been generated
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package gen

/*
#cgo LDFLAGS: -L${SRCDIR}/.. -lgogc -ldl
#include "../gogc.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// GogcVerifyPost function as declared in go-gc/gogc.h:27
func GogcVerifyPost(replicasPtr []GogcPublicReplicaInfo, replicasLen uint) GogcPublicReplicaInfoResponse {
	creplicasPtr, creplicasAllocMap := unpackArgSGogcPublicReplicaInfo(replicasPtr)
	defer creplicasAllocMap.Free()

	creplicasLen, _ := (C.size_t)(replicasLen), cgoAllocsUnknown
	__ret := C.gogc_verify_post(creplicasPtr, creplicasLen)
	packSGogcPublicReplicaInfo(replicasPtr, creplicasPtr)
	__v := *NewGogcPublicReplicaInfoResponseRef(unsafe.Pointer(&__ret))
	return __v
}

// GogcVerifyPostSleep function as declared in go-gc/gogc.h:30
func GogcVerifyPostSleep(replicasPtr []GogcPublicReplicaInfo, replicasLen uint) GogcPublicReplicaInfoResponse {
	creplicasPtr, creplicasAllocMap := unpackArgSGogcPublicReplicaInfo(replicasPtr)
	defer creplicasAllocMap.Free()

	creplicasLen, _ := (C.size_t)(replicasLen), cgoAllocsUnknown
	__ret := C.gogc_verify_post_sleep(creplicasPtr, creplicasLen)
	packSGogcPublicReplicaInfo(replicasPtr, creplicasPtr)
	__v := *NewGogcPublicReplicaInfoResponseRef(unsafe.Pointer(&__ret))
	return __v
}
