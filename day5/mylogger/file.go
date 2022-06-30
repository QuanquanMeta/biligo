package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// write to file

// FileLogger
type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

// ctor
func NewFileLogger(lv LogLevel, fp, fn string, maxSize int64) *FileLogger {
	fl := &FileLogger{
		Level:       lv,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err := fl.initFile() // open the filePath and fileName
	if err != nil {
		panic("err")
	}
	return fl
}

func (f *FileLogger) initFile() error {
	fullfileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullfileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}

	errFileObj, err := os.OpenFile(fullfileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed, err:%v\n", err)
		return err
	}

	// log files are all opened

	f.fileObj = fileObj
	f.errFileObj = errFileObj
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
		// check size
		if f.checkSize(f.fileObj) > f.maxFileSize {
			newFile, err := f.splitFile(f.fileObj) // new log
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02: 15:04:05"), lv, fileName, funcName, lineNo, msg)
		if level >= ERROR {
			if f.checkSize(f.errFileObj) > f.maxFileSize {
				newFile, err := f.splitFile(f.errFileObj) // new log
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02: 15:04:05"), lv, fileName, funcName, lineNo, msg)
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
