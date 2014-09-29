package main

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func test1() {
	for i := 0; i < 10; i++ {
		list := make([]string, 0)
		for i := 0; i < 1000; i++ {
			list = append(list, `The GOMAXPROCS variable limits the number of operating system threads that can execute user-level Go code simultaneously. There is no limit to the number of threads that can be blocked in system calls on behalf of Go code; those do not count against the GOMAXPROCS limit. This package's GOMAXPROCS function queries and changes the limit. `)
		}
		// time.Sleep(10 * time.Second)
	}
}

func test2() {
	list := make([]string, 100000)
	for i := 0; i < 10; i++ {
		for i := 0; i < 100000; i++ {
			list[i] = `The GOMAXPROCS variable limits the number of operating system threads that can execute user-level Go code simultaneously. There is no limit to the number of threads that can be blocked in system calls on behalf of Go code; those do not count against the GOMAXPROCS limit. This package's GOMAXPROCS function queries and changes the limit. `
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	f, err := os.Create("test.prof")
	if err != nil {
		log.Print(err.Error())
		return
	}

	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	test2()
}
