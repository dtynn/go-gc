package main

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
	ctrl := make(chan struct{}, 1024)
	// var wg sync.WaitGroup

	// attempts := 10000000
	// wg.Add(attempts)
	rand.Seed(time.Now().UnixNano())

	var rounds int64

	for i := 0; ; i++ {
		ctrl <- struct{}{}

		count := 800 << 10
		reps := make([]gen.GogcPublicReplicaInfo, count)
		for j := 0; j < count; j++ {
			commR := [32]byte{}
			rand.Read(commR[:])
			reps[j] = gen.GogcPublicReplicaInfo{
				SectorId: uint64(j + 117),
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

			// log.Println(i)
			gen.GogcVerifyPost(reps, uint(len(reps)))
			// runtime.GC()
			// free := make([]byte, len(reps)*int((C.size_t)(gen.SizeOfGogcPublicReplicaInfoValue)))
			// if len(free) == 0 {
			//     panic("0?")
			// }
		}(i)
	}

	// wg.Wait()
}
