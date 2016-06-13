package main

import (
	"os"
	"runtime/pprof"
)

func main() {
	// CPU
	cpu, _ := os.Create("cpu.out")
	defer cpu.Close()
	pprof.StartCPUProfile(cpu)
	defer pprof.StopCPUProfile()


	//Memory
	mem, _ := os.Create("mem.out")
	defer mem.Close()
defer pprof.WriteHeapProfile(mem)
}
