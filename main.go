package main

import "C"
import (
	"fmt"
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

	fire()
	// wg.Wait()
}

func fire() {

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
				if e := recover(); e != nil {
					panic(e)
				}
			}()

			defer func() {
				<-ctrl
				if total := atomic.AddInt64(&rounds, 1); total%10 == 0 {
					log.Printf("total %d rounds", total)
				}
			}()

			// log.Println(i)
			gen.GogcVerifyPost(reps, uint(len(reps)))
			for j := range reps {
				rep := reps[j]
				rep.Deref()
				if rep.SectorId != uint64(base+j) {
					panic(fmt.Sprintf("sector id mismatch: base=%d, prev=%d, after=%v", base, base+j, rep.SectorId))
				}

				if rep.CommR != commR {
					panic(fmt.Sprintf("comm_r mismatch, prev=%v, after=%v", commR, rep.CommR))
				}
			}
			// runtime.GC()
			// free := make([]byte, len(reps)*int((C.size_t)(gen.SizeOfGogcPublicReplicaInfoValue)))
			// if len(free) == 0 {
			//     panic("0?")
			// }
		}(i)
	}
}
