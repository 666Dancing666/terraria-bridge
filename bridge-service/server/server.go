package server

import (
"encoding/json"
"fmt"
"log"
"net/http"
"terraria-bridge/config"
"terraria-bridge/filter"
"terraria-bridge/metrics"
"terraria-bridge/session"
"time"

"github.com/gorilla/websocket"
)

var startTime time.Time

var upgrader = websocket.Upgrader{
CheckOrigin: func(r *http.Request) bool {
return true
},
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
tsOnline := session.Manager.GetTShock() != nil
players := session.Manager.GetAllPlayers()
stats := filter.Default.Stats()
latency := metrics.Stats.Snapshot()

playerList := make([]map[string]interface{}, 0)
for _, ps := range players {
playerList = append(playerList, map[string]interface{}{
"id": ps.ID,
})
}

status := map[string]interface{}{
"status":        "running",
"tshock_online": tsOnline,
"mc_players":    len(players),
"blocked_msgs":  stats,
"players":       playerList,
"uptime":        fmt.Sprintf("%ds", int(time.Since(startTime).Seconds())),
"latency":       latency,
}

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(status)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "text/html; charset=utf-8")
http.ServeFile(w, r, "web/dashboard.html")
}

func Start(cfg *config.Config) {
startTime = time.Now()

http.HandleFunc("/health", healthHandler)
http.HandleFunc("/", dashboardHandler)

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

log.Printf("Dashboard: http://0.0.0.0%s", cfg.Port)
log.Printf("TShock:    ws://0.0.0.0%s/tshock", cfg.Port)
log.Printf("MC:        ws://0.0.0.0%s/mc", cfg.Port)
log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
