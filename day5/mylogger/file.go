package mylogger

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"
)

// write to file

var (
	// Max size the chans
	MaxSize int = 50000
)

// FileLogger
type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
	logChan     chan *logMsg
}

type logMsg struct {
	Level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

// ctor
func NewFileLogger(lv LogLevel, fp, fn string, maxSize int64) *FileLogger {
	fl := &FileLogger{
		Level:       lv,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *logMsg, MaxSize),
	}
	err := fl.initFile() // open the filePath and fileName
	if err != nil {
		panic("err")
	}
	return fl
}

func (f *FileLogger) initFile() error {
	if _, err := os.Stat(f.filePath); os.IsNotExist(err) {
		err := os.Mkdir(f.filePath, 0644)
		if err != nil {
			fmt.Printf("create directory failed, err:%v\n", err)
			return err
		}
	}
	fullfileName := path.Join(f.filePath, f.fileName)
	absPath, _ := filepath.Abs(fullfileName)
	fileObj, err := os.OpenFile(absPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}

	errFileObj, err := os.OpenFile(absPath+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed, err:%v\n", err)
		return err
	}

	// log files are all opened

	f.fileObj = fileObj
	f.errFileObj = errFileObj

	// start 5 go roroutine. [TODO] having problems
	// for i := 0; i < 5; i++ {
	// 	go f.writeLogBackground()
	// }
	go f.writeLogBackground()
	return nil
}

func (f *FileLogger) log(lv, format string, a ...any) {

	level, err := parseStrtoLogLevel(lv)
	if err != nil {
		fmt.Println("parseStrtoLogLevel() failed")
		return
	}

	if f.enable(level) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInformation(3)

		logTemp := &logMsg{
			Level:     level,
			msg:       msg,
			funcName:  funcName,
			fileName:  fileName,
			timestamp: now.Format("2006-01-02: 15:04:05"),
			line:      lineNo,
		}

		select { // to prevent f.logChan to be full and blocked
		case f.logChan <- logTemp:
		default:
		}
		// check size
		if f.checkSize(f.fileObj) > f.maxFileSize {
			newFile, err := f.splitFile(f.fileObj) // new log
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
	}
}

// enable
func (f *FileLogger) enable(level LogLevel) bool {
	return level&f.Level != 0
}

func (f *FileLogger) Debug(format string, a ...any) {
	f.log("DEBUG", format, a...)
}

func (f *FileLogger) Info(format string, a ...any) {
	f.log("INFO", format, a...)
}

func (f *FileLogger) Warning(format string, a ...any) {
	f.log("WARNING", format, a...)
}

func (f *FileLogger) Error(format string, a ...any) {
	f.log("ERROR", format, a...)
}

func (f *FileLogger) Fatal(format string, a ...any) {
	f.log("FATAL", format, a...)
}

func (f *FileLogger) Close() {
	f.errFileObj.Close()
	f.fileObj.Close()
}

// log cut
// According to the size
// According to the date

func (f *FileLogger) checkSize(file *os.File) int64 {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get the file info %v", err)
		return 0
	}
	return fileInfo.Size()
}

func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// cut file
	// close current file and rename current file
	// create a new log file and set f.fileobj to the new file

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get the file info %v", err)
		return nil, err
	}

	nowStr := time.Now().Format("2006-01-02_15-04-05-000")
	fullPathName := path.Join(f.filePath, fileInfo.Name())
	newLogName := fmt.Sprintf("%s-%s.log", fullPathName, nowStr)

	file.Close()

	os.Rename(fullPathName, newLogName)

	fileObj, err := os.OpenFile(fullPathName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file fialed, err:%v\n", err)
		return nil, err
	}

	return fileObj, err
}

func (f *FileLogger) checkHour(file *os.File) bool {
	return true
}

func (f *FileLogger) writeLogBackground() {
	for {
		select {
		case logTemp := <-f.logChan:
			logLevelStr, _ := parseLogLevelToString(logTemp.Level)
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", logTemp.timestamp, logLevelStr, logTemp.fileName, logTemp.funcName, logTemp.line, logTemp.msg)

			fmt.Fprint(f.fileObj, logInfo)
			if logTemp.Level >= ERROR {
				if f.checkSize(f.errFileObj) > f.maxFileSize {
					newFile, err := f.splitFile(f.errFileObj) // new log
					if err != nil {
						return
					}
					f.errFileObj = newFile
				}
				fmt.Fprint(f.errFileObj, logInfo)
			}
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}
