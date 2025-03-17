package pkg

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func init() {
	// 確保日誌目錄存在
	os.MkdirAll("logs", 0755)

	// 創建日誌文件
	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}

	infoLogger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime)
	errorLogger = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// LogInfo 記錄信息日誌
func LogInfo(format string, v ...interface{}) {
	if infoLogger != nil {
		infoLogger.Printf(format, v...)
	}
}

// LogError 記錄錯誤日誌
func LogError(format string, v ...interface{}) {
	if errorLogger != nil {
		errorLogger.Printf(format, v...)
	}
}

// LogRequest 記錄請求信息
func LogRequest(command string, args []string) {
	LogInfo("Command: %s, Args: %v, Time: %s", command, args, time.Now().Format(time.RFC3339))
}
