package buffer

import (
"sync"
"terraria-bridge/protocol"
)

type MessageBuffer struct {
mu       sync.Mutex
messages []protocol.Message
maxSize  int
}

func New(maxSize int) *MessageBuffer {
return &MessageBuffer{
messages: make([]protocol.Message, 0),
maxSize:  maxSize,
}
}

func (b *MessageBuffer) Push(msg protocol.Message) {
b.mu.Lock()
defer b.mu.Unlock()
if len(b.messages) >= b.maxSize {
b.messages = b.messages[1:]
}
b.messages = append(b.messages, msg)
}

func (b *MessageBuffer) PopAll() []protocol.Message {
b.mu.Lock()
defer b.mu.Unlock()
msgs := b.messages
b.messages = make([]protocol.Message, 0)
return msgs
}

func (b *MessageBuffer) Len() int {
b.mu.Lock()
defer b.mu.Unlock()
return len(b.messages)
}
