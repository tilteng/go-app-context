package app_context

import (
	"time"

	"github.com/shirou/gopsutil/net"
	"github.com/tilteng/go-metrics/metrics"
)

func (self *baseAppContext) sendNetworkStats(previous *metrics.ProcStats, current *metrics.ProcStats) {
	delta := current.Timestamp.Sub(previous.Timestamp).Seconds()

	for i, counters := range current.IOCounters {
		var prev_counters *net.IOCountersStat

		if (i < len(previous.IOCounters)) &&
			(previous.IOCounters[i].Name == counters.Name) {
			prev_counters = &previous.IOCounters[i]
		} else {
			for _, prev := range previous.IOCounters {
				if prev.Name == counters.Name {
					prev_counters = &prev
					break
				}
			}
		}

		if prev_counters == nil {
			continue
		}

		prefix := "proc_stats.net." + counters.Name

		self.metricsClient.Histogram(
			prefix+".bytes_sent",
			float64(counters.BytesSent-prev_counters.BytesSent),
			delta,
			nil,
		)

		self.metricsClient.Histogram(
			prefix+".bytes_recv",
			float64(counters.BytesRecv-prev_counters.BytesRecv),
			delta,
			nil,
		)

		self.metricsClient.Histogram(
			prefix+".packets_sent",
			float64(counters.PacketsSent-prev_counters.PacketsSent),
			delta,
			nil,
		)

		self.metricsClient.Histogram(
			prefix+".packets_recv",
			float64(counters.PacketsRecv-prev_counters.PacketsRecv),
			delta,
			nil,
		)

		self.metricsClient.Count(
			prefix+".num_errors_out",
			int64(counters.Errout-prev_counters.Errout),
			delta,
			nil,
		)
		self.metricsClient.Count(
			prefix+".num_errors_in",
			int64(counters.Errin-prev_counters.Errin),
			delta,
			nil,
		)

		self.metricsClient.Count(
			prefix+".num_dropped_out",
			int64(counters.Dropout-prev_counters.Dropout),
			delta,
			nil,
		)
		self.metricsClient.Count(
			prefix+".num_dropped_in",
			int64(counters.Dropin-prev_counters.Dropin),
			delta,
			nil,
		)
	}
}

func (self *baseAppContext) sendMemStats(previous *metrics.ProcStats, current *metrics.ProcStats) {
	delta := current.Timestamp.Sub(previous.Timestamp).Seconds()

	self.metricsClient.Histogram(
		"proc_stats.mem.alloc.non_freed_bytes",
		float64(current.MemStats.Alloc),
		delta,
		nil,
	)

	self.metricsClient.Histogram(
		"proc_stats.mem.alloc.total_bytes",
		float64(current.MemStats.Alloc),
		delta,
		nil,
	)

	self.metricsClient.Count(
		"proc_stats.mem.alloc.count",
		int64(current.MemStats.Mallocs-previous.MemStats.Mallocs),
		delta,
		nil,
	)

	self.metricsClient.Histogram(
		"proc_stats.mem.heap.bytes_alloc",
		float64(current.MemStats.HeapAlloc),
		delta,
		nil,
	)

	self.metricsClient.Histogram(
		"proc_stats.mem.heap.bytes_in_use",
		float64(current.MemStats.HeapInuse),
		delta,
		nil,
	)

	self.metricsClient.Histogram(
		"proc_stats.mem.heap.bytes_released",
		float64(current.MemStats.HeapReleased),
		delta,
		nil,
	)

	self.metricsClient.Histogram(
		"proc_stats.mem.heap.num_objects",
		float64(current.MemStats.HeapObjects),
		delta,
		nil,
	)

	self.metricsClient.Histogram(
		"proc_stats.mem.gc.pause_ms",
		float64((current.MemStats.PauseTotalNs-previous.MemStats.PauseTotalNs))/float64(time.Millisecond),
		delta,
		nil,
	)

	self.metricsClient.Count(
		"proc_stats.mem.gc.count",
		int64(current.MemStats.NumGC)-int64(previous.MemStats.NumGC),
		delta,
		nil,
	)
}

func (self *baseAppContext) sendCPUStats(previous *metrics.ProcStats, current *metrics.ProcStats) {
	delta := current.Timestamp.Sub(previous.Timestamp).Seconds()

	self.metricsClient.Gauge(
		"proc_stats.num_cpus",
		float64(current.NumCPUs),
		delta,
		nil,
	)

	self.metricsClient.Histogram(
		"proc_stats.cpu.user_percent",
		100.0*(current.CPUTimes.User-previous.CPUTimes.User),
		delta,
		nil,
	)

	self.metricsClient.Histogram(
		"proc_stats.cpu.sys_percent",
		100.0*(current.CPUTimes.System-previous.CPUTimes.System),
		delta,
		nil,
	)
}

func (self *baseAppContext) SendStats(previous *metrics.ProcStats, current *metrics.ProcStats) {
	if !self.metricsEnabled || previous == nil || current == nil {
		return
	}

	self.sendCPUStats(previous, current)
	self.sendMemStats(previous, current)
	self.sendNetworkStats(previous, current)

	delta := current.Timestamp.Sub(previous.Timestamp).Seconds()

	if db := self.DB(); db != nil {
		db_stats := db.Stats()
		self.metricsClient.Gauge(
			"proc_stats.db.num_connections",
			float64(db_stats.OpenConnections),
			delta,
			nil,
		)
	}

	self.metricsClient.Histogram(
		"proc_stats.num_goroutines",
		float64(current.NumGoRoutines),
		delta,
		nil,
	)

	self.metricsClient.Gauge(
		"proc_stats.files.num_open",
		float64(current.NumFDs),
		delta,
		nil,
	)
}
