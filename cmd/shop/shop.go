package main

import (
	"math/rand"
	"mxshop/app/mxshop/api"
	"os"
	"runtime"
	"time"
)

func main() {
	randSrc := rand.NewSource(time.Now().UnixNano())
	rand.New(randSrc)
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	admin.NewApp("api-server").Run()
}
