// Package runtimestats collects GO runtime statistics.
//
// To use, simple import runtimestats package in your program.
// Stats collection starts automatically and can not be stopped or changed.
package runtimestats

import (
	"runtime"
	"time"

	"github.com/zviadm/stats-go/metrics"
)

const (
	// After ~Go1.9, ReadMemStats pauses world for only ~25microseconds. Thus it
	// is should be safe to collect memstats every 10 seconds even for very processes
	// with large heaps.
	collectFrequency = 10 * time.Second
)

var (
	mallocsC    = metrics.DefineCounter("runtime/go/mallocs").V()
	freesC      = metrics.DefineCounter("runtime/go/frees").V()
	allocBytesC = metrics.DefineCounter("runtime/go/alloc_bytes").V()

	cgoCallsC   = metrics.DefineCounter("runtime/go/cgo_calls").V()
	goroutinesG = metrics.DefineGauge("runtime/go/goroutines").V()
)

func collect() {
	ticker := time.NewTicker(collectFrequency)
	memstatsPrev := new(runtime.MemStats)
	memstats := new(runtime.MemStats)
	var cgoCallsPrev int64
	for range ticker.C {
		runtime.ReadMemStats(memstats)

		mallocsC.Count(float64(memstats.Mallocs - memstatsPrev.Mallocs))
		freesC.Count(float64(memstats.Frees - memstatsPrev.Frees))
		allocBytesC.Count(float64(memstats.TotalAlloc - memstatsPrev.TotalAlloc))

		cgoCalls := runtime.NumCgoCall()
		cgoCallsC.Count(float64(cgoCalls - cgoCallsPrev))
		cgoCallsPrev = cgoCalls
		goroutinesG.Set(float64(runtime.NumGoroutine()))

		memstats, memstatsPrev = memstatsPrev, memstats
	}
}

func init() {

}
