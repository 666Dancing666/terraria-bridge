package server

import (
    "log"
    "terraria-bridge/protocol"
    "terraria-bridge/session"

    "github.com/gorilla/websocket"
)

func handleMC(conn *websocket.Conn) {
    log.Println("✅ MC 客户端已连接")
    session.Get().SetMC(conn)

    for {
        var msg protocol.Message
        err := conn.ReadJSON(&msg)
        if err != nil {
            log.Printf("❌ MC 连接断开: %v", err)
            session.Get().SetMC(nil)
            break
        }
        log.Printf("📤 [MC → TShock] 转发消息: %s", msg.Type)

        s := session.Get()
        if tsConn, ok := s.TShockConn.(*websocket.Conn); ok && tsConn != nil {
            err := tsConn.WriteJSON(msg)
            if err != nil {
                log.Printf("❌ 转发给 TShock 失败: %v", err)
            } else {
                log.Printf("   ✅ 已转发")
            }
        } else {
            log.Printf("   ⚠️ TShock 未连接，消息丢弃")
        }
    }
}
