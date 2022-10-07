package data

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Models struct {
	LogEntry LogEntry
}

type LogEntry struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Data      string    `json:"data"`
	CreatedAt time.Time `jsn:"createdAt"`
}

func New() *Models {
	return &Models{
		LogEntry: LogEntry{},
	}
}

func (l *LogEntry) Insert(entry LogEntry) error {
	f, err := os.OpenFile(".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	message, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	f.WriteString(fmt.Sprintf("%s\n", string(message)))
	return nil
}
