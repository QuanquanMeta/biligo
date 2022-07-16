package taillog

import (
	"context"
	"fmt"
	"time"

	"github.com/biligo/logagent/kafka"
	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
)

// TailTask is a log collecter
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	// for exisiting t.run
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancle := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancle,
	}
	tailObj.init() // init task
	return
}

func (t *TailTask) init() {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Printf("tail file failed, err:%v\n", err)

	}
	go t.run() //send to log to agent
}

// func ReadChan() <-chan *tail.Line {
// 	return tailObj.Lines
// }

func (t *TailTask) ReadChan() <-chan *tail.Line {
	return t.instance.Lines
}

func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("tail task:%s_%s exit...", t.path, t.topic)
			return
		case line := <-t.instance.Lines: // get line of logs from tailObj
			// 32. it to Kafka
			fmt.Printf("got log data form %s successful, log:%v, line:%s\n", t.path, t.topic, line.Text)
			kafka.SendToChan(t.topic, line.Text) // function call funct. good to set it async
			// send log to a chan
			// use single treahd obj
		default:
			time.Sleep(time.Microsecond * 300)
		}
	}
}

func testTail() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Printf("tail file failed, err:%v\n", err)
		return
	}

	var (
		line *tail.Line
		ok   bool
	)

	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename %s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:", line.Text)
	}

}
