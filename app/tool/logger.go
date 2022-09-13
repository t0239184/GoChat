package tool

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var Logger logrus.Logger
var requestId string

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

type ApLogHook struct {
}

/* Implement Logrus Hook Interface */
func (hook *ApLogHook) Levels() []logrus.Level {
   return logrus.AllLevels
}

/* Implement Logrus Hook Interface */
func (hook *ApLogHook) Fire(entry *logrus.Entry) error {
   entry.Data["request_id"] = GetRequestId()
   return nil
}

func init() {
	fmt.Println("[init] logger")
	generateRequestId()
	nowTime := time.Now()
	var logFilePath string
	if len(logFilePath) <= 0 {
		if dir, err := os.Getwd(); err == nil {
			logFilePath = dir + "/logs/"
		}
	}

	if err := os.MkdirAll(logFilePath, os.ModePerm); err != nil {
		fmt.Println(err.Error())
	}

	logFileName := nowTime.Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)

	if _, err := os.Stat(fileName); err != nil {
		_, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println("write file log error", err)
	}

	Logger = *logrus.New()
	Logger.SetOutput(io.MultiWriter(src, os.Stdout))
	//logger.SetLevel(logrus.InfoLevel)
	Logger.AddHook(&ApLogHook{})

	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
		FullTimestamp:true,
	})
}



func AccessLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		reqUri := c.Request.RequestURI
		reqMethod := c.Request.Method
		clientIP := c.ClientIP()
		reqUa := c.Request.UserAgent()

		beforeBody := make(map[string]interface{})
		beforeBody["clientIp"] = clientIP
		beforeBody["userAgent"] = reqUa
		beforeBody["requestMethod"] = reqMethod
		LogInfoAccess(beforeBody, reqUri)

		c.Next()

		/* After */
		responseStr := blw.body.String()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		statusCode := c.Writer.Status()

		resultBody := make(map[string]interface{})
		resultBody["response"] = responseStr
		resultBody["requestMethod"] = reqMethod
		resultBody["latencyTime"] = latencyTime
		resultBody["statusCode"] = statusCode
		LogInfoAccess(resultBody, reqUri)

		generateRequestId()
	}
}

/* For Access Log */
func LogInfoAccess(fields logrus.Fields, url string) {
	Logger.SetLevel(logrus.InfoLevel)
	Logger.WithFields(fields).Info(url)
}

func generateRequestId() {
	requestId = uuid.New().String()
}

func GetRequestId() string {
	return requestId
}
