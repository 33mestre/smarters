// Package models contains the data models for the Smarters project.
package models

import "time"

// LogMessage representa uma mensagem de log
// @description Estrutura que contém os detalhes de uma mensagem de log
type LogMessage struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
