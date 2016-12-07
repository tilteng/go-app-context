package app_context

import "github.com/tilteng/go-metrics/metrics"

func (self *baseAppContext) SendStats(previous *metrics.ProcStats, current *metrics.ProcStats) {
	if !self.metricsEnabled || previous == nil || current == nil {
		return
	}

	delta := current.Timestamp.Sub(previous.Timestamp).Seconds()

	self.metricsClient.Gauge(
		"proc_stats.num_cpus",
		float64(current.NumCPUs),
		delta,
		nil,
	)

	self.metricsClient.Histogram(
		"proc_stats.num_goroutines",
		float64(current.NumGoRoutines),
		delta,
		nil,
	)

	self.metricsClient.Histogram(
		"proc_stats.cpu.user_time",
		current.CPUTimes.User-previous.CPUTimes.User,
		delta,
		nil,
	)

	self.metricsClient.Histogram(
		"proc_stats.cpu.sys_time",
		current.CPUTimes.System-previous.CPUTimes.System,
		delta,
		nil,
	)

	self.metricsClient.Histogram(
		"proc_stats.cpu.idle_time",
		current.CPUTimes.Idle-previous.CPUTimes.Idle,
		delta,
		nil,
	)

	self.metricsClient.Histogram(
		"proc_stats.cpu.iowait_time",
		current.CPUTimes.Iowait-previous.CPUTimes.Iowait,
		delta,
		nil,
	)

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
		"proc_stats.mem.gc.pause_ns",
		float64(current.MemStats.PauseTotalNs-previous.MemStats.PauseTotalNs),
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
