package main

import "C"
import (
	"log"
	"math/rand"
	"os"
	"runtime"
	// "sync"
	"time"

	"github.com/dtynn/go-gc/gen"
)

func main() {
	log.Println(os.Getpid())

	rand.Seed(time.Now().UnixNano())

	base := rand.Intn(5000) + 1
	count := 4000000
	reps := make([]gen.GogcPublicReplicaInfo, count)
	for j := 0; j < count; j++ {
		commR := [32]byte{}
		rand.Read(commR[:])
		reps[j] = gen.GogcPublicReplicaInfo{
			SectorId: uint64(base + j),
			CommR:    commR,
		}
	}

	for i := 0; ; i++ {
		free := make([]byte, len(reps)*40)
		if len(free) == 0 {
			panic("0?")
		}

		resp := gen.GogcVerifyPost(reps, uint(len(reps)))
		resp.Deref()
		if resp.Count != uint64(count) {
			panic(resp.Count)
		}

		runtime.GC()
	}
}
