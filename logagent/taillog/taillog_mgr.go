package taillog

import "github.com/biligo/logagent/etcd"

var taskMgr *tailLogMgr

type tailLogMgr struct {
	logEntry []*etcd.LogEntry
	//taskMap map[stirng]*TailTask
}

func InitMgr(logEntryConf []*etcd.LogEntry) {
	taskMgr = &tailLogMgr{
		logEntry: logEntryConf, // store the current log info
	}
	for _, logEntry := range logEntryConf {
		// conf: *etcd.LogEntry
		NewTailTask(logEntry.Path, logEntry.Topic)
	}

}
