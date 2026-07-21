package session

import (
"sync"
"terraria-bridge/buffer"
"github.com/gorilla/websocket"
)

type PlayerSession struct {
ID       string
MCConn   *websocket.Conn
MCBuffer *buffer.MessageBuffer
}

type SessionManager struct {
mu           sync.Mutex
TShockConn   *websocket.Conn
TSBuffer     *buffer.MessageBuffer
players       map[string]*PlayerSession
}

var Manager = &SessionManager{
TSBuffer: buffer.New(1000),
players:  make(map[string]*PlayerSession),
}

func (sm *SessionManager) SetTShock(conn *websocket.Conn) {
sm.mu.Lock()
defer sm.mu.Unlock()
sm.TShockConn = conn
}

func (sm *SessionManager) GetTShock() *websocket.Conn {
sm.mu.Lock()
defer sm.mu.Unlock()
return sm.TShockConn
}

func (sm *SessionManager) AddPlayer(id string, conn *websocket.Conn) *PlayerSession {
sm.mu.Lock()
defer sm.mu.Unlock()
ps := &PlayerSession{
ID:       id,
MCConn:   conn,
MCBuffer: buffer.New(1000),
}
sm.players[id] = ps
return ps
}

func (sm *SessionManager) RemovePlayer(id string) {
sm.mu.Lock()
defer sm.mu.Unlock()
delete(sm.players, id)
}

func (sm *SessionManager) GetPlayer(id string) *PlayerSession {
sm.mu.Lock()
defer sm.mu.Unlock()
return sm.players[id]
}

func (sm *SessionManager) GetAllPlayers() []*PlayerSession {
sm.mu.Lock()
defer sm.mu.Unlock()
players := make([]*PlayerSession, 0, len(sm.players))
for _, ps := range sm.players {
players = append(players, ps)
}
return players
}
