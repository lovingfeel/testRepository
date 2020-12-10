package main

import (
	"fmt"
	"log"
	"os"
)
var logger *log.Logger
func init()  {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile )
	log.SetPrefix("success")
	logFile, err := os.OpenFile("./c.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常")
	}
	log.SetOutput(logFile)

	//自定义log
	//logFile, err := os.OpenFile("./c.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	//if err != nil {
	//	log.Panic("打开日志文件异常")
	//}
	//logger = log.New(logFile, "success", log.Ldate | log.Ltime | log.Lshortfile)
}
func main() {
	defer fmt.Println("panic退出前处理")
	log.Println("println日志")
	//logger.Println("println日志heheh")
	//log.Panic("panic日志")
	//log.Fatal("程序退出日志")
}
