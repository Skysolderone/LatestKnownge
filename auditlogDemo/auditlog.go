package main

import (
	"encoding/json"
	"os"
	"time"
)

type AuditLogEntry struct {
	Actor   string      `json:"actor"`
	Action  string      `json:"action"`
	Module  string      `json:"module"`
	When    time.Time   `json:"when"`
	Details interface{} `json:"details"`
}

func logAuditEvent(actor string, action string, module string, details interface{}) {
	entry := AuditLogEntry{Actor: actor, Action: action, Module: module, When: time.Now(), Details: details}
	logEntry, _ := json.Marshal(entry)
	file, _ := os.OpenFile("audit.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	defer file.Close()
	file.WriteString(string(logEntry) + "\n")
}
