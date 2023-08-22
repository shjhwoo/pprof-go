package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"perf/sol"
	"runtime"
	"runtime/pprof"
)

// https://graphviz.org/download/ 시각화를 위해 graphbiz 설치 필요함.
//수집한 pprof를 템플릿에 담아 /debug/pprof 로 노출

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
var blockprofile = flag.String("blockprofile", "", "write block profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// Set memory (heap) profile
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // Get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}

	if *blockprofile != "" {
		f, err := os.Create(*blockprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // Get up-to-date statistics
		runtime.SetBlockProfileRate(10)
		f.Close()
	}

	go func() {
		http.ListenAndServe("localhost:3001", nil)
	}()

	var result = sol.Solution([][]int{{1, 4}, {3, 2}, {4, 1}}, [][]int{{3, 3}, {3, 3}})
	fmt.Println("최종 결과값", result)

	var result2 = sol.Solution([][]int{{2, 3, 2}, {4, 2, 4}, {3, 1, 4}}, [][]int{{5, 4, 3}, {2, 4, 1}, {3, 1, 1}})
	fmt.Println("최종 결과값", result2)

	// Add some code here to keep your program running
	// For example, you can use a blocking channel like this:
	select {}
}

//go run main.go 실행 후 http://localhost:3001/debug/pprof/ 접속.
