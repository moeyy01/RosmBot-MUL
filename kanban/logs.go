package kanban

import (
	"bytes"
	"fmt"
	"io"
	"os"
	p "path"

	"github.com/FloatTech/floatbox/file"
	log "github.com/sirupsen/logrus"
)

type ColorFormatter struct{}

func (m ColorFormatter) Format(entry *log.Entry) ([]byte, error) {
	var color int
	switch entry.Level {
	case log.ErrorLevel:
		color = 1 //red
	case log.WarnLevel:
		color = 3 //yellow
	case log.InfoLevel:
		color = 2 //green
	case log.DebugLevel:
		color = 5
	default:
		color = 7 //白
	}
	var buff *bytes.Buffer
	if entry.Buffer == nil {
		buff = new(bytes.Buffer)
	} else {
		buff = entry.Buffer
	}
	//时间
	formatTime := entry.Time.Format("15:04:06")
	//设置格式
	fmt.Fprintf(buff, "\033[3%dm[%s]\033[0m %s %s\n", color, entry.Level, formatTime, entry.Message)
	return buff.Bytes(), nil
}

type ColorNotFormatter struct{}

func (m ColorNotFormatter) Format(entry *log.Entry) ([]byte, error) {
	var buff *bytes.Buffer
	if entry.Buffer == nil {
		buff = new(bytes.Buffer)
	} else {
		buff = entry.Buffer
	}
	//时间
	formatTime := entry.Time.Format("15:04:06")
	//设置格式
	fmt.Fprintf(buff, "[%s] %s %s\n", entry.Level, formatTime, entry.Message)
	return buff.Bytes(), nil
}

func SetLogWithNewFile(path string) {
	if dir := p.Dir(path); file.IsNotExist(dir) {
		_ = os.MkdirAll(dir, 0755)
	}
	file, err := os.Create(path)
	if err != nil {
		log.Errorln("[log]创建log文件失败", err)
		return
	}
	log.SetOutput(io.MultiWriter(file, os.Stdout))
}
