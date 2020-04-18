package main

import "C"
import (
	"math/rand"
	// "runtime"
	// "sync"

	"github.com/dtynn/go-gc/gen"
)

func main() {
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

	gen.GogcVerifyPost(reps, uint(len(reps)))

}
