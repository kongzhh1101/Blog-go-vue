package core

import (
	"Blog/global"
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type Logformatter struct{}

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *Logformatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据不同的level去展示颜色
	var levelcolor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelcolor = gray
	case logrus.WarnLevel:
		levelcolor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelcolor = red
	default:
		levelcolor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	var log = global.Config.Logger.Perfix

	// 自定义日期格式
	timestamp := entry.Time.Format("2016-01-02 15:01:05")
	if entry.HasCaller() {
		// 自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		// 自定义输出格式
		fmt.Fprintf(b, "[%s][%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", log, timestamp, levelcolor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s][%s] \x1b[%dm[%s]\x1b[0m %s\n", log, timestamp, levelcolor, entry.Level, entry.Message)
	}

	return b.Bytes(), nil
}

func InitLogger() *logrus.Logger {
	level, _ := logrus.ParseLevel(global.Config.Logger.Level)

	mLog := logrus.New()                                // 新建一个实例
	mLog.SetOutput(os.Stdout)                           // 设置输出类型
	mLog.SetReportCaller(global.Config.Logger.ShowLine) // 开启返回函数名和行号
	mLog.SetFormatter(&Logformatter{})                  //设置自定义的Formatter
	mLog.SetLevel(level)                                // 设置最低的Level
	// InitDefaultLogger()
	return mLog
}

func InitDefaultLogger() {
	level, _ := logrus.ParseLevel(global.Config.Logger.Level)
	// 全局Log
	logrus.SetOutput(os.Stdout)                           // 设置输出类型
	logrus.SetReportCaller(global.Config.Logger.ShowLine) // 开启返回函数名和行号
	logrus.SetFormatter(&Logformatter{})                  //设置自定义的Formatter
	logrus.SetLevel(level)
}
