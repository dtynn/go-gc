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
import (
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// allocGogcPublicReplicaInfoResponseMemory allocates memory for type C.gogc_PublicReplicaInfoResponse in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGogcPublicReplicaInfoResponseMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGogcPublicReplicaInfoResponseValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGogcPublicReplicaInfoResponseValue = unsafe.Sizeof([1]C.gogc_PublicReplicaInfoResponse{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GogcPublicReplicaInfoResponse) Ref() *C.gogc_PublicReplicaInfoResponse {
	if x == nil {
		return nil
	}
	return x.ref8f165427
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GogcPublicReplicaInfoResponse) Free() {
	if x != nil && x.allocs8f165427 != nil {
		x.allocs8f165427.(*cgoAllocMap).Free()
		x.ref8f165427 = nil
	}
}

// NewGogcPublicReplicaInfoResponseRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGogcPublicReplicaInfoResponseRef(ref unsafe.Pointer) *GogcPublicReplicaInfoResponse {
	if ref == nil {
		return nil
	}
	obj := new(GogcPublicReplicaInfoResponse)
	obj.ref8f165427 = (*C.gogc_PublicReplicaInfoResponse)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GogcPublicReplicaInfoResponse) PassRef() (*C.gogc_PublicReplicaInfoResponse, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref8f165427 != nil {
		return x.ref8f165427, nil
	}
	mem8f165427 := allocGogcPublicReplicaInfoResponseMemory(1)
	ref8f165427 := (*C.gogc_PublicReplicaInfoResponse)(mem8f165427)
	allocs8f165427 := new(cgoAllocMap)
	allocs8f165427.Add(mem8f165427)

	var ccount_allocs *cgoAllocMap
	ref8f165427.count, ccount_allocs = (C.uint64_t)(x.Count), cgoAllocsUnknown
	allocs8f165427.Borrow(ccount_allocs)

	x.ref8f165427 = ref8f165427
	x.allocs8f165427 = allocs8f165427
	return ref8f165427, allocs8f165427

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GogcPublicReplicaInfoResponse) PassValue() (C.gogc_PublicReplicaInfoResponse, *cgoAllocMap) {
	if x.ref8f165427 != nil {
		return *x.ref8f165427, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GogcPublicReplicaInfoResponse) Deref() {
	if x.ref8f165427 == nil {
		return
	}
	x.Count = (uint64)(x.ref8f165427.count)
}

// allocGogcPublicReplicaInfoMemory allocates memory for type C.gogc_PublicReplicaInfo in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGogcPublicReplicaInfoMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGogcPublicReplicaInfoValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGogcPublicReplicaInfoValue = unsafe.Sizeof([1]C.gogc_PublicReplicaInfo{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GogcPublicReplicaInfo) Ref() *C.gogc_PublicReplicaInfo {
	if x == nil {
		return nil
	}
	return x.refafdcbbac
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GogcPublicReplicaInfo) Free() {
	if x != nil && x.allocsafdcbbac != nil {
		x.allocsafdcbbac.(*cgoAllocMap).Free()
		x.refafdcbbac = nil
	}
}

// NewGogcPublicReplicaInfoRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGogcPublicReplicaInfoRef(ref unsafe.Pointer) *GogcPublicReplicaInfo {
	if ref == nil {
		return nil
	}
	obj := new(GogcPublicReplicaInfo)
	obj.refafdcbbac = (*C.gogc_PublicReplicaInfo)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GogcPublicReplicaInfo) PassRef() (*C.gogc_PublicReplicaInfo, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refafdcbbac != nil {
		return x.refafdcbbac, nil
	}
	memafdcbbac := allocGogcPublicReplicaInfoMemory(1)
	refafdcbbac := (*C.gogc_PublicReplicaInfo)(memafdcbbac)
	allocsafdcbbac := new(cgoAllocMap)
	allocsafdcbbac.Add(memafdcbbac)

	var ccomm_r_allocs *cgoAllocMap
	refafdcbbac.comm_r, ccomm_r_allocs = *(*[32]C.uint8_t)(unsafe.Pointer(&x.CommR)), cgoAllocsUnknown
	allocsafdcbbac.Borrow(ccomm_r_allocs)

	var csector_id_allocs *cgoAllocMap
	refafdcbbac.sector_id, csector_id_allocs = (C.uint64_t)(x.SectorId), cgoAllocsUnknown
	allocsafdcbbac.Borrow(csector_id_allocs)

	x.refafdcbbac = refafdcbbac
	x.allocsafdcbbac = allocsafdcbbac
	return refafdcbbac, allocsafdcbbac

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GogcPublicReplicaInfo) PassValue() (C.gogc_PublicReplicaInfo, *cgoAllocMap) {
	if x.refafdcbbac != nil {
		return *x.refafdcbbac, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GogcPublicReplicaInfo) Deref() {
	if x.refafdcbbac == nil {
		return
	}
	x.CommR = *(*[32]byte)(unsafe.Pointer(&x.refafdcbbac.comm_r))
	x.SectorId = (uint64)(x.refafdcbbac.sector_id)
}

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

const sizeOfPtr = unsafe.Sizeof(&struct{}{})

// unpackArgSGogcPublicReplicaInfo transforms a sliced Go data structure into plain C format.
func unpackArgSGogcPublicReplicaInfo(x []GogcPublicReplicaInfo) (unpacked *C.gogc_PublicReplicaInfo, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)

	len0 := len(x)
	mem0 := allocGogcPublicReplicaInfoMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: mem0,
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.gogc_PublicReplicaInfo)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.gogc_PublicReplicaInfo)(h.Data)
	return
}

// packSGogcPublicReplicaInfo reads sliced Go data structure out from plain C format.
func packSGogcPublicReplicaInfo(v []GogcPublicReplicaInfo, ptr0 *C.gogc_PublicReplicaInfo) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGogcPublicReplicaInfoValue]C.gogc_PublicReplicaInfo)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGogcPublicReplicaInfoRef(unsafe.Pointer(&ptr1))
	}
}
