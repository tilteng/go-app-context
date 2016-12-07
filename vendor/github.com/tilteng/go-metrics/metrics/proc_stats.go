package metrics

import (
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"
)

type ProcStats struct {
	Timestamp     time.Time
	NumCPUs       int
	NumGoRoutines int
	MemStats      runtime.MemStats
	CPUTimes      cpu.TimesStat
}

func GetProcStats() *ProcStats {
	sample := &ProcStats{
		Timestamp:     time.Now(),
		NumCPUs:       runtime.NumCPU(),
		NumGoRoutines: runtime.NumGoroutine(),
	}

	runtime.ReadMemStats(&sample.MemStats)

	if p, err := process.NewProcess(int32(os.Getpid())); err == nil {
		if t, err := p.Times(); err == nil {
			sample.CPUTimes = *t
		}
	}

	return sample
}
