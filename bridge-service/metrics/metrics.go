package metrics

import (
"sync"
"time"
)

type LatencyStats struct {
mu          sync.Mutex
totalIn     int64
totalOut    int64
lastInTime  time.Time
lastOutTime time.Time
lastLatency time.Duration
}

var Stats = &LatencyStats{}

func (s *LatencyStats) RecordIn() {
s.mu.Lock()
defer s.mu.Unlock()
s.totalIn++
s.lastInTime = time.Now()
if !s.lastOutTime.IsZero() {
s.lastLatency = s.lastInTime.Sub(s.lastOutTime)
}
}

func (s *LatencyStats) RecordOut() {
s.mu.Lock()
defer s.mu.Unlock()
s.totalOut++
s.lastOutTime = time.Now()
}

func (s *LatencyStats) Snapshot() map[string]interface{} {
s.mu.Lock()
defer s.mu.Unlock()

latency := "N/A"
if s.lastLatency > 0 {
latency = s.lastLatency.String()
}

return map[string]interface{}{
"total_in":     s.totalIn,
"total_out":    s.totalOut,
"last_latency": latency,
}
}
