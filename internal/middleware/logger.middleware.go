package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LoggerMiddleware tạo ra một middleware để ghi log với Zap.
func LoggerMiddleware() gin.HandlerFunc {
	// Cấu hình Zap để ghi log vào file theo ngày
	currentDate := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("logs/access-%s.log", currentDate)

	// Tạo thư mục logs nếu chưa tồn tại
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	// Mở file log
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	// Cấu hình logger
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

	return func(c *gin.Context) {
		start := time.Now()

		// Chạy tiếp các middleware và handler tiếp theo
		c.Next()

		// Ghi log sau khi xử lý xong request
		logger.Info("Request::",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", time.Since(start)),
			zap.String("clientIP", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
		)
	}
}
