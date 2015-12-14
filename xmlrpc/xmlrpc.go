package xmlrpc

import (
	"bytes"
	"github.com/divan/gorilla-xmlrpc/xml"
	"github.com/gorilla/rpc"
	"log"
	"net/http"
)

func Create(isType string) {

	if isType == "server" {
		Server()
	} else {
		Client()
	}
}

type MessageService struct{}

func (ms *MessageService) Reply(r *http.Request, args *struct{ Who string }, reply *struct{ Message string }) error {
	log.Println("Reply", args.Who)
	reply.Message = "Hello, " + args.Who + "!"
	return nil
}

func Server() {
	RPC := rpc.NewServer()
	xmlrpcCodec := xml.NewCodec()
	RPC.RegisterCodec(xmlrpcCodec, "text/xml")
	RPC.RegisterService(new(MessageService), "")
	http.Handle("/RPC2", RPC)

	log.Println("Starting XML-RPC Server on localhost:1234/RPC2")
	log.Fatal(http.ListenAndServe(":1234", nil))
}

func XmlRpcCall(method string, args struct{ Who string }) (reply struct{ Message string }, err error) {

	buf, _ := xml.EncodeClientRequest(method, &args)

	resp, err := http.Post("http://localhost:1234/RPC2", "text/xml", bytes.NewBuffer(buf))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = xml.DecodeClientResponse(resp.Body, &reply)
	return
}

func Client() {
	reply, err := XmlRpcCall("MessageService.Reply", struct{ Who string }{"User 1"})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response: %s\n", reply.Message)
}
