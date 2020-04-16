package main

import "C"
import (
	"log"
	"math/rand"
	"os"
	// "runtime"
	"sync"
	"time"

	"github.com/dtynn/go-gc/gen"
)

func main() {
	log.Println(os.Getpid())
	time.Sleep(20 * time.Second)
	ctrl := make(chan struct{}, 1024)
	var wg sync.WaitGroup

	attempts := 10000000
	wg.Add(attempts)
	rand.Seed(time.Now().UnixNano())

	count := 120
	reps := make([]gen.GogcPublicReplicaInfo, count)
	for j := 0; j < count; j++ {
		commR := [32]byte{}
		rand.Read(commR[:])
		reps[j] = gen.GogcPublicReplicaInfo{
			SectorId: uint64(j),
			CommR:    commR,
		}
	}

	for i := 0; i < attempts; i++ {
		go func(i int) {
			defer func() {
				wg.Done()
				<-ctrl
			}()

			ctrl <- struct{}{}

			log.Println(i)
			gen.GogcVerifyPost(reps, uint(len(reps)))
			// runtime.GC()
			// free := make([]byte, len(reps)*int((C.size_t)(gen.SizeOfGogcPublicReplicaInfoValue)))
			// if len(free) == 0 {
			//     panic("0?")
			// }
		}(i)
	}

	wg.Wait()
}
