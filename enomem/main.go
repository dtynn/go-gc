package main

import "C"
import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"sync/atomic"
	"time"

	gen "github.com/dtynn/go-gc/gen_no_set_finalizer"
)

func main() {
	log.Println(os.Getpid())

	fire()
}

func fire() {
	ctrl := make(chan struct{}, 10)

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

			resp := gen.GogcVerifyPost(reps, uint(len(reps)))
			resp.Deref()

			if resp.Count != uint64(count) {
				panic(fmt.Sprintf("count: %d != %d", resp.Count, count))
			}

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

			runtime.GC()
		}(i)
	}
}
