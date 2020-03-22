package main

import (
    "fmt"
    "time"
    "strings"
    "os"
    "bufio"
    "io"
)

type Reader interface{
    Read(rc chan []byte)
}

type Writer interface{
    Write(wc chan string)
}

type ReadFromFile struct {
    path string
}

func (r *ReadFromFile) Read(rc chan []byte){
    f , err  := os.Open(r.path)

    if err != nil {
        panic(fmt.Sprintf("open file error:%s",err.Error()))
    }
    f.Seek(0,2)

    rd := bufio.NewReader(f)
    for{
        line,err := rd.ReadBytes('\n')
        if err  == io.EOF {
            time.Sleep(500 * time.Millisecond)
            continue
        }else if err  != nil {
            panic(fmt.Sprintf("ReadBytes error:%s",err.Error()))
        }
        rc <- line[:len(line)-1]
        
    }

}

type WriteToInfluxDB struct {
    infuxDBsn string
}

func (w *WriteToInfluxDB) Write(wc chan string){
    for y := range wc {
        fmt.Println(y)
    }

}


type LogProcess struct {
    rc chan []byte
    wc chan string
    read Reader
    write Writer
}

// func (l *LogProcess) ReadFromFile(){
//     line := "message"
//     l.rc <- line 
// }

func (l *LogProcess) Process(){

    for v := range l.rc{
        l.wc  <- strings.ToUpper(string(v))
    }


}



func main(){

        r := &ReadFromFile {
            path: "./access.log",
        }

        w := &WriteToInfluxDB{
            infuxDBsn: "username&password",
        }
        lp := &LogProcess{
            rc: make(chan []byte),
            wc: make(chan string),
            read: r,
            write: w,
        }
        go lp.read.Read(lp.rc)
        go lp.Process()
        go lp.write.Write(lp.wc)
        time.Sleep(30 * time.Second)
}


