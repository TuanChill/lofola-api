package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LoggerInstance() *zap.Logger {
	// set up logger every day
	currentDate := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("logs/access-%s.log", currentDate)

	// create logs folder if not exist
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	// check if file exists
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	// write logs to file
	writer := zapcore.AddSync(file)

	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encodeConfig),
		writer,
		zap.InfoLevel,
	)
	logger := zap.New(core)

	return logger
}

func LogError(err string) {
	logger := LoggerInstance()
	start := time.Now()

	logger.Error("Error::",
		zap.String("error", err),
		zap.Duration("latency", time.Since(start)),
	)
}

func LogInfo(info string) {
	logger := LoggerInstance()
	start := time.Now()

	logger.Info("Info::",
		zap.String("info", info),
		zap.Duration("latency", time.Since(start)),
	)
}
