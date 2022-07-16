package taillog

import (
	"fmt"
	"time"

	"github.com/biligo/logagent/etcd"
)

var taskMgr *tailLogMgr

type tailLogMgr struct {
	logEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func InitMgr(logEntryConf []*etcd.LogEntry) {
	taskMgr = &tailLogMgr{
		logEntry:    logEntryConf, // store the current log info
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry), // chan without buff
	}

	for _, logEntry := range logEntryConf {
		// conf: *etcd.LogEntry
		// record tailtask
		tailObj := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		taskMgr.tskMap[mk] = tailObj
	}

	go taskMgr.run()
}

// watch newConfChan, process requests
func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			// config create
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.tskMap[conf.Path]
				if ok {
					continue
				} else {
					// new
					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.tskMap[mk] = tailObj
				}
			}

			// find it is in t.logEntry, but not in newConf. delete it
			for _, c1 := range t.logEntry {
				isDelete := true
				for _, c2 := range newConf {
					if c2.Path == c1.Path && c2.Topic == c1.Topic {
						isDelete = false
						continue
					}
				}
				if isDelete {
					// delete
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					t.tskMap[mk].cancelFunc()
				}
			}

			// config update
			fmt.Println("new config comes!", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// export an interface for newConfChan of tskMgr
func NewConf() chan<- []*etcd.LogEntry {
	return taskMgr.newConfChan
}
