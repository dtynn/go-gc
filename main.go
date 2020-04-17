package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"log"
	"math/rand"
	"os"
	"unsafe"
	// "runtime"
	// "sync"
	"time"

	"github.com/dtynn/go-gc/gen"
)

func main() {
	log.Println(os.Getpid())
	time.Sleep(20 * time.Second)
	// ctrl := make(chan struct{}, 2<<10)
	// var wg sync.WaitGroup

	// attempts := 10000000
	// wg.Add(attempts)
	rand.Seed(time.Now().UnixNano())

	ptrs := make([]unsafe.Pointer, 0, 5)
	for i := 0; ; i++ {
		ptr := gen.AllocGogcPublicReplicaInfoMemory(800 << 10)
		ptrs = append(ptrs, ptr)

		if len(ptrs) == 5 {
			for i := range ptrs {
				C.free(ptrs[i])
			}

			ptrs = ptrs[:0]
		}
	}

	// wg.Wait()
}
