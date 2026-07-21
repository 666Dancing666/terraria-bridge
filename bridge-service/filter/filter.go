package filter

import "sync"

type Filter struct {
mu         sync.Mutex
allowedIn  map[string]bool
allowedOut map[string]bool
blockedIn  map[string]bool
blockedOut map[string]bool
stats      map[string]int
}

var Default = New()

func New() *Filter {
f := &Filter{
allowedIn:  make(map[string]bool),
allowedOut: make(map[string]bool),
blockedIn:  make(map[string]bool),
blockedOut: make(map[string]bool),
stats:      make(map[string]int),
}

f.allowedIn["world_snapshot"] = true
f.allowedIn["tile_update"] = true
f.allowedIn["wall_update"] = true
f.allowedIn["liquid_update"] = true
f.allowedIn["entity_update"] = true
f.allowedIn["entity_remove"] = true
f.allowedIn["player_join"] = true
f.allowedIn["player_leave"] = true
f.allowedIn["chat_message"] = true
	f.allowedIn["time_sync"] = true
	f.allowedIn["weather_sync"] = true
	f.allowedIn["game_event"] = true
f.allowedIn["player_move"] = true

f.allowedOut["player_move"] = true
f.allowedOut["player_action"] = true
f.allowedOut["tile_break"] = true
f.allowedOut["tile_place"] = true
f.allowedOut["interact"] = true
f.allowedOut["chat_message"] = true

return f
}

func (f *Filter) AllowIn(msgType string) {
f.mu.Lock()
defer f.mu.Unlock()
f.allowedIn[msgType] = true
}

func (f *Filter) BlockIn(msgType string) {
f.mu.Lock()
defer f.mu.Unlock()
f.blockedIn[msgType] = true
}

func (f *Filter) AllowOut(msgType string) {
f.mu.Lock()
defer f.mu.Unlock()
f.allowedOut[msgType] = true
}

func (f *Filter) BlockOut(msgType string) {
f.mu.Lock()
defer f.mu.Unlock()
f.blockedOut[msgType] = true
}

func (f *Filter) PassIn(msgType string) bool {
f.mu.Lock()
defer f.mu.Unlock()
if f.blockedIn[msgType] {
return false
}
return f.allowedIn[msgType]
}

func (f *Filter) PassOut(msgType string) bool {
f.mu.Lock()
defer f.mu.Unlock()
if f.blockedOut[msgType] {
return false
}
return f.allowedOut[msgType]
}

func (f *Filter) RecordBlocked(msgType string) {
f.mu.Lock()
defer f.mu.Unlock()
f.stats[msgType]++
}

func (f *Filter) Stats() map[string]int {
f.mu.Lock()
defer f.mu.Unlock()
stats := make(map[string]int)
for k, v := range f.stats {
stats[k] = v
}
return stats
}
