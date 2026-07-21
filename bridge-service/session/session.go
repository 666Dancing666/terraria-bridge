package session

import "sync"

type Session struct {
    mu        sync.Mutex
    TShockConn interface{}
    MCConn     interface{}
}

var current = &Session{}

func Get() *Session {
    return current
}

func (s *Session) SetTShock(conn interface{}) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.TShockConn = conn
}

func (s *Session) SetMC(conn interface{}) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.MCConn = conn
}

func (s *Session) BothConnected() bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.TShockConn != nil && s.MCConn != nil
}
