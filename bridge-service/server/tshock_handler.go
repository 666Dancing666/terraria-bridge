package server

import (
    "log"
    "terraria-bridge/protocol"
    "terraria-bridge/session"

    "github.com/gorilla/websocket"
)

func handleTShock(conn *websocket.Conn) {
    log.Println("✅ TShock 已连接")
    session.Get().SetTShock(conn)

    for {
        var msg protocol.Message
        err := conn.ReadJSON(&msg)
        if err != nil {
            log.Printf("❌ TShock 连接断开: %v", err)
            session.Get().SetTShock(nil)
            break
        }
        log.Printf("📤 [TShock → MC] 转发消息: %s", msg.Type)

        s := session.Get()
        if mcConn, ok := s.MCConn.(*websocket.Conn); ok && mcConn != nil {
            err := mcConn.WriteJSON(msg)
            if err != nil {
                log.Printf("❌ 转发给 MC 失败: %v", err)
            } else {
                log.Printf("   ✅ 已转发")
            }
        } else {
            log.Printf("   ⚠️ MC 未连接，消息丢弃")
        }
    }
}
