package xmlrpc

import (
	"HashtableRepo/hashtable"
	"bytes"
	"github.com/divan/gorilla-xmlrpc/xml"
	"github.com/gorilla/rpc"
	"log"
	"net/http"
	//"time"
)

func Create(isType string) {

	if isType == "server" {
		Server()
	} else {
		Client()
	}
}

type MessageService struct{}

func (ms *MessageService) Print(r *http.Request, args *struct{ Who string }, reply *struct{ Message string }) error {
	log.Println("Call RPC Print")
	//hashtable.Print()

	var p hashtable.Pair
	hashtable.Put(p.New(1, "one"))
	/*hashtable.Put(p.New(2, "two"))
	hashtable.Put(p.New(5, "five"))
	hashtable.Put(p.New(7, "seven"))
	hashtable.Put(p.New(12, "twelve"))
	hashtable.Put(p.New(49, "fourty-nine"))
	hashtable.Put(p.New(33, "thirty-three"))
	hashtable.Put(p.New(8, "eigth"))*/

	str := hashtable.ToString()
	log.Println("Hashtable: " + str)
	reply.Message = "Hello, " + args.Who + "!"
	return nil
}

func Server() {
	RPC := rpc.NewServer()
	xmlrpcCodec := xml.NewCodec()
	RPC.RegisterCodec(xmlrpcCodec, "text/xml")
	RPC.RegisterService(new(MessageService), "")
	http.Handle("/hashtable", RPC)

	log.Println("Starting XML-RPC Server on localhost:8081/hashtable")
	log.Fatal(http.ListenAndServe(":8081", nil))

	hashtable.Init(16)
}

func XmlRpcCall(method string, args struct{ Who string }) (reply struct{ Message string }, err error) {

	buf, _ := xml.EncodeClientRequest(method, &args)

	resp, err := http.Post("http://localhost:8081/hashtable", "text/xml", bytes.NewBuffer(buf))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = xml.DecodeClientResponse(resp.Body, &reply)
	return
}

func Client() {

	log.Println("Run client...")
	//time.Sleep(5000 * time.Millisecond)
	//reply, err := XmlRpcCall("MessageService.Reply", struct{ Who string }{"User 1"})
	//reply, err := XmlRpcCall("Service.Print", struct{ Who string }{"User 1"})
	reply, err := XmlRpcCall("MessageService.Print", struct{ Who string }{"User 1"})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response: %s\n", reply.Message)
}
