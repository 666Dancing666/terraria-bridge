package logger

import (
"encoding/json"
"fmt"
"os"
"sync"
"time"
)

type Logger struct {
mu   sync.Mutex
file *os.File
}

var Default *Logger

func Init(logPath string) error {
os.MkdirAll(logPath, 0755)
f, err := os.OpenFile(
logPath+"/bridge-"+time.Now().Format("2006-01-02")+".log",
os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
if err != nil {
return err
}
Default = &Logger{file: f}
return nil
}

func (l *Logger) Log(direction string, msgType string, payload interface{}) {
l.mu.Lock()
defer l.mu.Unlock()

data, _ := json.Marshal(payload)
line := fmt.Sprintf("[%s] %s %s %s\n",
time.Now().Format("15:04:05.000"),
direction,
msgType,
string(data))
l.file.WriteString(line)
l.file.Sync()
}

func (l *Logger) Close() {
if l.file != nil {
l.file.Close()
}
}
