package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"log"
	"math/rand"
	"os"
	// "runtime"
	// "sync"
	"sync/atomic"
	"time"

	"github.com/dtynn/go-gc/gen"
)

func main() {
	log.Println(os.Getpid())
	// time.Sleep(20 * time.Second)
	ctrl := make(chan struct{}, 10)
	// var wg sync.WaitGroup

	// attempts := 10000000
	// wg.Add(attempts)
	rand.Seed(time.Now().UnixNano())

	var rounds int64

	for i := 0; ; i++ {
		ctrl <- struct{}{}

		base := rand.Int() + 1
		commR := [32]byte{}
		rand.Read(commR[:])

		count := 800 << 10
		reps := make([]gen.GogcPublicReplicaInfo, count)
		for j := 0; j < count; j++ {
			reps[j] = gen.GogcPublicReplicaInfo{
				SectorId: uint64(j + base),
				CommR:    commR,
			}
		}

		go func(i int) {
			defer func() {
				<-ctrl
				if total := atomic.AddInt64(&rounds, 1); total%10 == 0 {
					log.Printf("total %d rounds", total)
				}
			}()

			ptr, err := C.calloc(800<<10, 40)
			defer C.free(ptr)

			if err != nil {
				panic("boooooooom " + err.Error())
			}

			time.Sleep(50 * time.Millisecond)
		}(i)
	}

	// wg.Wait()
}
