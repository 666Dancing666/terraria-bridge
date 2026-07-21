package server

import (
"encoding/json"
"log"
"net/http"
"terraria-bridge/config"
"terraria-bridge/filter"
"terraria-bridge/session"

"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
CheckOrigin: func(r *http.Request) bool {
return true
},
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
tsOnline := session.Manager.GetTShock() != nil
players := len(session.Manager.GetAllPlayers())
stats := filter.Default.Stats()

status := map[string]interface{}{
"status":       "running",
"tshock_online": tsOnline,
"mc_players":    players,
"blocked_msgs":  stats,
}

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(status)
}

func Start(cfg *config.Config) {
http.HandleFunc("/health", healthHandler)

http.HandleFunc("/tshock", func(w http.ResponseWriter, r *http.Request) {
conn, err := upgrader.Upgrade(w, r, nil)
if err != nil {
log.Printf("WebSocket upgrade failed: %v", err)
return
}
handleTShock(conn)
})

http.HandleFunc("/mc", func(w http.ResponseWriter, r *http.Request) {
conn, err := upgrader.Upgrade(w, r, nil)
if err != nil {
log.Printf("WebSocket upgrade failed: %v", err)
return
}
handleMC(conn)
})

log.Printf("TShock: ws://0.0.0.0%s/tshock", cfg.Port)
log.Printf("MC:     ws://0.0.0.0%s/mc", cfg.Port)
log.Printf("Health: http://0.0.0.0%s/health", cfg.Port)
log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
