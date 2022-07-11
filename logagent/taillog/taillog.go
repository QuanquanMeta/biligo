package taillog

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
)

func Init(fileName string) (err error) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Printf("tail file failed, err:%v\n", err)
		return
	}
	return
}

func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
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
func main() {

}
