package mcproto

import (
"bufio"
"bytes"
"compress/zlib"
"encoding/binary"
"encoding/json"
"fmt"
"io"
"log"
"net"
"time"
)

func StartMCProxy(mcPort string) {
ln, err := net.Listen("tcp", ":"+mcPort)
if err != nil {
log.Printf("MC proxy failed to start: %v", err)
return
}
log.Printf("MC Server: 0.0.0.0:%s", mcPort)

for {
conn, err := ln.Accept()
if err != nil {
continue
}
go handleMCConnection(conn)
}
}

func handleMCConnection(conn net.Conn) {
defer conn.Close()

conn.SetDeadline(time.Now().Add(30 * time.Second))

reader := bufio.NewReader(conn)

_ , _ = readVarInt(reader)
packetID, _ := readVarInt(reader)

if packetID != 0x00 {
return
}

protocolVer, _ := readVarInt(reader)
hostLen, _ := readVarInt(reader)
hostBytes := make([]byte, hostLen)
io.ReadFull(reader, hostBytes)
host := string(hostBytes)

port := make([]byte, 2)
io.ReadFull(reader, port)

nextState, _ := readVarInt(reader)

if nextState == 1 {
handlePing(conn, protocolVer, host)
} else if nextState == 2 {
handleLogin(conn, protocolVer, host)
}
}

func handlePing(conn net.Conn, protocolVer int32, host string) {
conn.SetDeadline(time.Now().Add(5 * time.Second))

reader := bufio.NewReader(conn)
readVarInt(reader)
readVarInt(reader)

status := map[string]interface{}{
"version": map[string]interface{}{
"name":     "Terraria Bridge 1.20.1",
"protocol": 763,
},
"players": map[string]interface{}{
"max":    100,
"online": 0,
"sample": []map[string]string{},
},
"description": map[string]string{
"text": "Terraria x Minecraft Bridge",
},
"favicon": "",
}

jsonData, _ := json.Marshal(status)

var buf bytes.Buffer
writeVarInt(&buf, 0x00)
writeString(&buf, string(jsonData))

var out bytes.Buffer
writeVarInt(&out, int32(buf.Len()))
out.Write(buf.Bytes())

conn.Write(out.Bytes())

readVarInt(reader)
_ , _ = readVarInt(reader)
payload := make([]byte, 8)
io.ReadFull(reader, payload)

var pong bytes.Buffer
writeVarInt(&pong, int32(9))
writeVarInt(&pong, 0x01)
pong.Write(payload)

var pongOut bytes.Buffer
writeVarInt(&pongOut, int32(pong.Len()))
pongOut.Write(pong.Bytes())

conn.Write(pongOut.Bytes())
}

func handleLogin(conn net.Conn, protocolVer int32, host string) {
conn.SetDeadline(time.Now().Add(10 * time.Second))

reader := bufio.NewReader(conn)
readVarInt(reader)
readVarInt(reader)

nameLen, _ := readVarInt(reader)
nameBytes := make([]byte, nameLen)
io.ReadFull(reader, nameBytes)
playerName := string(nameBytes)

log.Printf("MC Login: %s from %s", playerName, conn.RemoteAddr())

uuid := "00000000-0000-0000-0000-000000000000"

var buf bytes.Buffer
writeVarInt(&buf, 0x02)
writeString(&buf, uuid)
writeString(&buf, playerName)
writeVarInt(&buf, 0)

var out bytes.Buffer
writeVarInt(&out, int32(buf.Len()))
out.Write(buf.Bytes())

conn.Write(out.Bytes())

loginAck := make([]byte, 3)
io.ReadFull(conn, loginAck)

dimCodec := `{"type":"minecraft:overworld","generator":{"type":"minecraft:flat","settings":{"layers":[{"block":"minecraft:air","height":1}],"structures":{"structures":{}}}}}`
dimCodecLen := len(dimCodec)
dimBytes := make([]byte, dimCodecLen+1)
dimBytes[0] = byte(dimCodecLen)
copy(dimBytes[1:], dimCodec)

var joinBuf bytes.Buffer
writeVarInt(&joinBuf, 0x25)
writeVarInt(&joinBuf, 0)
writeString(&joinBuf, "minecraft:overworld")
writeString(&joinBuf, dimCodec)
writeString(&joinBuf, "minecraft:overworld")
writeString(&joinBuf, dimCodec)
writeVarInt(&joinBuf, 0)
writeVarInt(&joinBuf, 8)
writeVarInt(&joinBuf, 8)
writeBool(&joinBuf, false)
writeBool(&joinBuf, true)
writeBool(&joinBuf, false)
writeBool(&joinBuf, false)

var joinOut bytes.Buffer
writeVarInt(&joinOut, int32(joinBuf.Len()))
joinOut.Write(joinBuf.Bytes())

conn.Write(joinOut.Bytes())

posBuf := make([]byte, 41)
posBuf[0] = 0x1A
binary.BigEndian.PutUint64(posBuf[1:9], 0)
binary.BigEndian.PutUint64(posBuf[9:17], 64)
binary.BigEndian.PutUint64(posBuf[17:25], 1)
posBuf[25] = 0
posBuf[26] = 0
posBuf[27] = 0x01
posBuf[28] = 0x00
binary.BigEndian.PutUint32(posBuf[29:33], 0)
binary.BigEndian.PutUint32(posBuf[33:37], 0)
posBuf[37] = 0
posBuf[38] = 0x00
posBuf[39] = 0x01
posBuf[40] = 0x00

var posOut bytes.Buffer
writeVarInt(&posOut, int32(len(posBuf)))
posOut.Write(posBuf)

conn.Write(posOut.Bytes())

keepAlive := make(chan bool)
go func() {
for {
time.Sleep(5 * time.Second)
select {
case <-keepAlive:
return
default:
var ka bytes.Buffer
id := time.Now().UnixMilli()
writeVarInt(&ka, 0x1F)
binary.Write(&ka, binary.BigEndian, id)
var kaOut bytes.Buffer
writeVarInt(&kaOut, int32(ka.Len()))
kaOut.Write(ka.Bytes())
conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
if _, err := conn.Write(kaOut.Bytes()); err != nil {
return
}
}
}
}()

for {
_, err := readVarInt(reader)
if err != nil {
break
}
_, err = readVarInt(reader)
if err != nil {
break
}
}

close(keepAlive)
log.Printf("MC Disconnected: %s", playerName)
}

func readVarInt(r io.ByteReader) (int32, error) {
var result int32
var shift uint
for {
b, err := r.ReadByte()
if err != nil {
return 0, err
}
result |= int32(b&0x7F) << shift
if b&0x80 == 0 {
break
}
shift += 7
}
return result, nil
}

func writeVarInt(w io.Writer, value int32) {
for {
b := byte(value & 0x7F)
value >>= 7
if value != 0 {
b |= 0x80
}
w.Write([]byte{b})
if value == 0 {
break
}
}
}

func writeString(w io.Writer, s string) {
writeVarInt(w, int32(len(s)))
w.Write([]byte(s))
}

func writeBool(w io.Writer, b bool) {
if b {
w.Write([]byte{0x01})
} else {
w.Write([]byte{0x00})
}
}

func readString(r io.Reader) (string, error) {
length, err := readVarInt(r.(io.ByteReader))
if err != nil {
return "", err
}
buf := make([]byte, length)
_, err = io.ReadFull(r, buf)
return string(buf), err
}

var _ = zlib.NewReader
var _ = fmt.Sprintf
