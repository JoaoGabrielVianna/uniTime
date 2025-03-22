package config

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Logger struct com diferentes níveis de log
type Logger struct {
	Info    *log.Logger
	Success *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Context string // Aqui é onde o nome do contexto será armazenado
}

// Cores ANSI para terminal
const (
	colorReset  = "\033[0m"
	colorBlue   = "\033[34;1m"
	colorGreen  = "\033[32;1m"
	colorYellow = "\033[33;1m"
	colorRed    = "\033[31;1m"
)

// Função para formatar a mensagem de log
func formatMessage(level string, color string, context string, message string) string {
	return fmt.Sprintf("%s[%s]%s [%s] - %s%s%s %s", color, context, colorReset, time.Now().Format("2006/01/02"), color, level, colorReset, message)
}

// Função para criar um novo logger
func NewLogger(context string) *Logger {
	// Retorna o Logger configurado com o contexto
	return &Logger{
		Info:    log.New(os.Stdout, "", 0),
		Success: log.New(os.Stdout, "", 0),
		Warning: log.New(os.Stdout, "", 0),
		Error:   log.New(os.Stderr, "", 0),
		Context: context,
	}
}

// Funções para facilitar o uso do logger
func (l *Logger) InfoLog(message string) {
	l.Info.Println(formatMessage("INFO", colorBlue, l.Context, message))
}

func (l *Logger) SuccessLog(message string) {
	l.Success.Println(formatMessage("SUCCESS", colorGreen, l.Context, message))
}

func (l *Logger) WarningLog(message string) {
	l.Warning.Println(formatMessage("WARNING", colorYellow, l.Context, message))
}

func (l *Logger) ErrorLog(message string) {
	l.Error.Println(formatMessage("ERROR", colorRed, l.Context, message))
}
