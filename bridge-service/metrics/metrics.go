package metrics

import (
"fmt"
"sync"
"time"
)

type LatencyStats struct {
mu          sync.Mutex
totalIn     int64
totalOut    int64
lastLatency int64
lastTime    int64
}

var Stats = &LatencyStats{}

func (s *LatencyStats) RecordIn() {
s.mu.Lock()
defer s.mu.Unlock()
s.totalIn++
s.lastTime = time.Now().UnixMilli()
}

func (s *LatencyStats) RecordOut() {
s.mu.Lock()
defer s.mu.Unlock()
s.totalOut++
if s.lastTime > 0 {
s.lastLatency = time.Now().UnixMilli() - s.lastTime
s.lastTime = 0
}
}

func (s *LatencyStats) Snapshot() map[string]interface{} {
s.mu.Lock()
defer s.mu.Unlock()

latency := "N/A"
if s.lastLatency > 0 {
latency = fmt.Sprintf("%dms", s.lastLatency)
}

return map[string]interface{}{
"total_in":     s.totalIn,
"total_out":    s.totalOut,
"last_latency": latency,
}
}
