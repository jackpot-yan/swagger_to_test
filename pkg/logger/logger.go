package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

// LogLevel 定义日志级别
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

// Logger 结构体
type Logger struct {
	level    LogLevel
	logger   *log.Logger
	file     *os.File
	mutex    sync.Mutex
	filename string
}

// 全局日志实例
var defaultLogger *Logger
var once sync.Once

// 日志级别字符串表示
var levelStrings = map[LogLevel]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

// InitLogger 初始化日志系统
func InitLogger(logFile string, level LogLevel) error {
	once.Do(func() {
		// 创建日志目录
		dir := filepath.Dir(logFile)
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}

		// 打开日志文件
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}

		// 创建日志实例
		defaultLogger = &Logger{
			level:    level,
			logger:   log.New(file, "", log.LstdFlags),
			file:     file,
			filename: logFile,
		}
	})
	return nil
}

// GetLogger 获取日志实例
func GetLogger() *Logger {
	if defaultLogger == nil {
		InitLogger("logs/app.log", INFO)
	}
	return defaultLogger
}

// getCallerInfo 获取调用者信息
func getCallerInfo(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown", 0
	}
	// 获取函数名
	pc, _, _, _ := runtime.Caller(skip)
	fn := runtime.FuncForPC(pc)
	funcName := "unknown"
	if fn != nil {
		funcName = fn.Name()
		parts := strings.Split(funcName, ".")
		funcName = parts[len(parts)-1]
	}
	// 简化文件名
	file = filepath.Base(file)
	return fmt.Sprintf("%s:%s", file, funcName), line
}

// logInternal 内部日志记录方法
func (l *Logger) logInternal(level LogLevel, format string, v ...interface{}) {
	if level < l.level {
		return
	}

	l.mutex.Lock()
	defer l.mutex.Unlock()

	// 获取调用者信息
	caller, line := getCallerInfo(3)
	timeStr := time.Now().Format("2006-01-02 15:04:05.000")
	levelStr := levelStrings[level]
	msg := fmt.Sprintf(format, v...)

	// 格式化日志
	logMsg := fmt.Sprintf("[%s] %s %s:%d %s", levelStr, timeStr, caller, line, msg)

	// 输出到文件和控制台
	l.logger.Println(logMsg)
	if level == FATAL {
		os.Exit(1)
	}
}

// 日志方法
func (l *Logger) Debug(format string, v ...interface{}) {
	l.logInternal(DEBUG, format, v...)
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.logInternal(INFO, format, v...)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	l.logInternal(WARN, format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.logInternal(ERROR, format, v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	l.logInternal(FATAL, format, v...)
}

func Debug(format string, v ...interface{}) {
	GetLogger().Debug(format, v...)
}

func Info(format string, v ...interface{}) {
	GetLogger().Info(format, v...)
}

func Warn(format string, v ...interface{}) {
	GetLogger().Warn(format, v...)
}

func Error(format string, v ...interface{}) {
	GetLogger().Error(format, v...)
}

func Fatal(format string, v ...interface{}) {
	GetLogger().Fatal(format, v...)
}

// Close 关闭日志文件
func (l *Logger) Close() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.file != nil {
		l.file.Close()
		l.file = nil
	}
}
