package server

import (
    "log"
    "net/http"
    "terraria-bridge/config"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func Start(cfg *config.Config) {
    http.HandleFunc("/tshock", func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            log.Printf("WebSocket 升级失败: %v", err)
            return
        }
        handleTShock(conn)
    })

    http.HandleFunc("/mc", func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            log.Printf("WebSocket 升级失败: %v", err)
            return
        }
        handleMC(conn)
    })

    log.Printf("🚀 中间层服务启动")
    log.Printf("   TShock 连接: ws://0.0.0.0%s/tshock", cfg.Port)
    log.Printf("   MC 连接:     ws://0.0.0.0%s/mc", cfg.Port)
    log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
