package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Reader interface {
	Read(readChan chan []byte)
}
type Writer interface {
	Write(writeChan chan string)
}

type LogProcess struct {
	readChan  chan []byte
	writeChan chan string
	read      Reader
	write     Writer
}

type ReadFromFile struct {
	path string
}

type WriteToInfluxDB struct {
	influxDBsn string
}

func (r *ReadFromFile) Read(readChan chan []byte) {
	//line:="hello world!"
	//readChan<-line
	//open file
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file error:%s", err.Error()))
	}
	//	read data from file  by line

	//end side
	f.Seek(0, 2)
	bf := bufio.NewReader(f)
	for {
		line, err := bf.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(1 * time.Second)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadBytes() error:%s", err.Error()))
		}
		//delete \n
		readChan <- line[:len(line)-1]
	}

}
func (l *LogProcess) Process() {
	for v := range l.readChan {
		l.writeChan <- strings.ToUpper(string(v))
	}
	//data:=<-l.readChan

}

//func (l *LogProcess)WriteToInfluxDB()  {
//	fmt.Println(<-l.writeChan)
//}

func (w *WriteToInfluxDB) Write(writeChan chan string) {
	for v := range writeChan {
		fmt.Println(v)
	}
}
func main() {
	r := &ReadFromFile{
		path: "./log_test/access.log",
	}
	w := &WriteToInfluxDB{
		influxDBsn: "username&password",
	}
	logProcess := &LogProcess{
		readChan:  make(chan []byte),
		writeChan: make(chan string),
		read:      r,
		write:     w,
	}

	go logProcess.read.Read(logProcess.readChan)
	go logProcess.Process()
	go logProcess.write.Write(logProcess.writeChan)
	time.Sleep(100 * time.Second)
}
