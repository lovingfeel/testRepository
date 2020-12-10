package main

import (
	"context"
	"fmt"
	"log"
	"mygoproject/mythriftgo/test"
	"mygoproject/mythriftgo/thrift"
	//"net"
)
func main() {
	//transportFactory := thrift.NewTBufferedTransportFactory(8192)
	//mysocket, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "7912"))
	mysocket,_ :=thrift.NewTSocket("127.0.0.1:7912")
	transport :=thrift.NewTFramedTransport(mysocket)
	//protocolFactory := thrift.NewTCompactProtocolFactory()
	protocolFactory :=thrift.NewTBinaryProtocolFactoryDefault();
	//protocol :=thrift.NewTBinaryProtocolTransport(transport)

	//transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "9898"))
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, "error resolving address:", err)
	//	os.Exit(1)
	//}

	//useTransport, err := transportFactory.GetTransport(transport)
	client := test.NewHelloWordServiceClientFactory(transport,protocolFactory)

	transport.Open()
	//if err := transport.Open(); err != nil {
	//	fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:9898", " ", err)
	//	os.Exit(1)
	//}
	defer transport.Close()
	var defaultCtx = context.Background()
	//req := &echo.EchoReq{Msg:"You are welcome."}
	//res, err := client.Echo(req)
	 var tmd int32   = 24
	// var tangage *int32 = &tmd
	req := &test.Request{test.RequestType_SAY_HELLO,"唐",&tmd}
	res, err := client.DoAction(defaultCtx,req)
	if err != nil {
		log.Println("Echo failed:", err)
		return
	}

	log.Println("response:", res)
	fmt.Println("well done")
}
